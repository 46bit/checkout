package pricing_rules

import (
	"math/rand"
	"reflect"
	"testing/quick"
)

type BuyOneGetOneFree struct {
	UnitPrice uint `yaml:"unit_price"`
}

func NewBuyOneGetOneFree(unitPrice uint) *BuyOneGetOneFree {
	return &BuyOneGetOneFree{UnitPrice: unitPrice}
}

func (r *BuyOneGetOneFree) Price(numberOfItems uint) uint {
	effectiveNumberOfItems := numberOfItems/2 + numberOfItems%2
	return effectiveNumberOfItems * r.UnitPrice
}

var _ PricingRule = new(BuyOneGetOneFree)

func (r BuyOneGetOneFree) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(&BuyOneGetOneFree{
		UnitPrice: uint(rand.Uint32()),
	})
}

var _ quick.Generator = new(BuyOneGetOneFree)
