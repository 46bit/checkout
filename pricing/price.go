package pricing

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type Price struct {
	Price       uint
	PricingRule *PricingRule
}

type yamlPrice struct {
	Price       uint                   `yaml:"price"`
	PricingRule map[string]interface{} `yaml:"pricing_rule,omitempty"`
}

func (p Price) MarshalYAML() (interface{}, error) {
	var yamlPricingRule map[string]interface{}
	if p.PricingRule != nil {
		if err := redecodeYaml(p.PricingRule, &yamlPricingRule); err != nil {
			return nil, err
		}
		switch (*p.PricingRule).(type) {
		case *BulkDiscount:
			yamlPricingRule["type"] = "bulk_discount"
		case *BuyOneGetOneFree:
			yamlPricingRule["type"] = "buy_one_get_one_free"
		}
	}

	return yamlPrice{
		Price:       p.Price,
		PricingRule: yamlPricingRule,
	}, nil
}

var _ yaml.Marshaler = new(Price)

func (p *Price) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var yamlPrice yamlPrice
	if err := unmarshal(&yamlPrice); err != nil {
		return err
	}
	p.Price = yamlPrice.Price
	if _, ok := yamlPrice.PricingRule["type"]; !ok {
		return nil
	}

	ruleType, ok := yamlPrice.PricingRule["type"].(string)
	if !ok {
		return fmt.Errorf("rule type field not a string")
	}
	var pricingRule PricingRule
	switch ruleType {
	case "bulk_discount":
		pricingRule = &BulkDiscount{}
	case "buy_one_get_one_free":
		pricingRule = &BuyOneGetOneFree{}
	default:
		return fmt.Errorf("rule type '%s' unrecognised", ruleType)
	}
	p.PricingRule = &pricingRule

	delete(yamlPrice.PricingRule, "rule")
	return redecodeYaml(yamlPrice.PricingRule, &p.PricingRule)
}

var _ yaml.Unmarshaler = new(Config)

func redecodeYaml(from, to interface{}) error {
	data, err := yaml.Marshal(from)
	if err == nil {
		err = yaml.UnmarshalStrict([]byte(data), to)
	}
	return err
}
