package pricing_rules

import (
	"math/rand"
	"reflect"
	"testing/quick"
)

type BulkDiscount struct {
	MinimumNumberOfItems uint `yaml:"minimum_number_of_items"`
	StandardPrice        uint `yaml:"standard_price"`
	DiscountedPrice      uint `yaml:"discounted_price"`
}

func (r *BulkDiscount) Price(numberOfItems uint) uint {
	unitPrice := r.StandardPrice
	if numberOfItems >= r.MinimumNumberOfItems {
		unitPrice = r.DiscountedPrice
	}
	return numberOfItems * unitPrice
}

var _ PricingRule = new(BulkDiscount)

func (r BulkDiscount) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(&BulkDiscount{
		MinimumNumberOfItems: uint(rand.Uint32()),
		StandardPrice:        uint(rand.Uint32()),
		DiscountedPrice:      uint(rand.Uint32()),
	})
}

var _ quick.Generator = new(BulkDiscount)
