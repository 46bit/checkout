package pricing_rules

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"math/rand"
	"reflect"
	"testing/quick"
)

type Config struct {
	PricingRules map[string]PricingRule
}

func NewConfig() Config {
	return Config{
		PricingRules: map[string]PricingRule{},
	}
}

func (c Config) MarshalYAML() (interface{}, error) {
	var configMap_ struct {
		PricingRules map[string]map[string]interface{} `yaml:"pricing_rules"`
	}
	data, err := yaml.Marshal(c.PricingRules)
	if err != nil {
		return nil, err
	}
	err = yaml.UnmarshalStrict([]byte(data), &configMap_.PricingRules)
	if err != nil {
		return nil, err
	}

	for itemID, pricingRule := range c.PricingRules {
		switch pricingRule.(type) {
		case *BulkDiscount:
			configMap_.PricingRules[itemID]["rule"] = "bulk_discount"
		case *BuyOneGetOneFree:
			configMap_.PricingRules[itemID]["rule"] = "buy_one_get_one_free"
		case *Fixed:
			configMap_.PricingRules[itemID]["rule"] = "fixed"
		}
	}

	return configMap_, nil
}

var _ yaml.Marshaler = new(Config)

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if c.PricingRules == nil {
		c.PricingRules = map[string]PricingRule{}
	}

	var configMap_ struct {
		PricingRules map[string]map[string]interface{} `yaml:"pricing_rules"`
	}
	if err := unmarshal(&configMap_); err != nil {
		return err
	}

	for itemID, pricingRuleMap := range configMap_.PricingRules {
		typeInterface, ok := pricingRuleMap["rule"]
		if !ok {
			return fmt.Errorf("no rule type field when deserialising pricing rule")
		}
		type_, ok := typeInterface.(string)
		if !ok {
			return fmt.Errorf("rule type field not a string when deserialising pricing rule")
		}

		delete(pricingRuleMap, "rule")
		data, err := yaml.Marshal(pricingRuleMap)
		if err != nil {
			return err
		}

		var pricingRule PricingRule
		switch type_ {
		case "bulk_discount":
			pricingRule = new(BulkDiscount)
		case "buy_one_get_one_free":
			pricingRule = new(BuyOneGetOneFree)
		case "fixed":
			pricingRule = new(Fixed)
		default:
			return fmt.Errorf("rule type '%s' unrecognised when deserialising pricing rule", type_)
		}

		err = yaml.UnmarshalStrict([]byte(data), pricingRule)
		if err != nil {
			return err
		}
		c.PricingRules[itemID] = pricingRule
	}

	return nil
}

var _ yaml.Unmarshaler = new(Config)

func (c Config) Generate(rand *rand.Rand, size int) reflect.Value {
	c2 := NewConfig()
	rulesCount := rand.Intn(size)
	for i := 0; i < rulesCount; i++ {
		chars := "abcdefghijklmnopqrstuvwxyz0123456789-_"
		length := rand.Intn(size) + 1
		bytes := make([]byte, length)
		for i := 0; i < length; i++ {
			bytes[i] = chars[rand.Intn(len(chars))]
		}
		itemID := string(bytes)

		var pricingRule PricingRule
		switch rand.Int31n(3) {
		case 0:
			pricingRule = BulkDiscount{}.Generate(rand, size).Interface().(PricingRule)
		case 1:
			pricingRule = BuyOneGetOneFree{}.Generate(rand, size).Interface().(PricingRule)
		case 2:
			pricingRule = Fixed{}.Generate(rand, size).Interface().(PricingRule)
		}

		c2.PricingRules[itemID] = pricingRule
	}
	return reflect.ValueOf(c2)
}

var _ quick.Generator = new(Config)
