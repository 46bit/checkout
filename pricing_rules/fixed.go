package pricing_rules

import (
	"math/rand"
	"reflect"
	"testing/quick"
)

type Fixed struct {
	UnitPrice uint `yaml:"unit_price"`
}

func (r *Fixed) Price(numberOfItems uint) uint {
	return numberOfItems * r.UnitPrice
}

var _ PricingRule = new(Fixed)

func (r Fixed) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(&Fixed{
		UnitPrice: uint(rand.Uint32()),
	})
}

var _ quick.Generator = new(Fixed)
