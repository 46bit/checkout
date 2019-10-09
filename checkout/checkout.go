package checkout

import (
	"fmt"
	"github.com/46bit/checkout/pricing"
)

type Checkout struct {
	pricingRules map[string]pricing.PricingRule
	scannedItems map[string]uint
}

func New(pricingRules map[string]pricing.PricingRule) *Checkout {
	return &Checkout{
		pricingRules: pricingRules,
		scannedItems: map[string]uint{},
	}
}

func (c *Checkout) Scan(item string) error {
	_, ok := c.pricingRules[item]
	if !ok {
		return fmt.Errorf("item '%s' not found in the pricing rules", item)
	}

	c.scannedItems[item] += 1
	return nil
}

func (c *Checkout) Total() uint {
	total := uint(0)
	for item, count := range c.scannedItems {
		total += c.pricingRules[item].Price(count)
	}
	return total
}
