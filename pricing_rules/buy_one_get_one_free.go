package pricing_rules

type BuyOneGetOneFree struct {
	unitPrice uint
}

func NewBuyOneGetOneFree(unitPrice uint) BuyOneGetOneFree {
	return BuyOneGetOneFree{unitPrice}
}

func (r BuyOneGetOneFree) Price(numberOfItems uint) uint {
	effectiveNumberOfItems := numberOfItems/2 + numberOfItems%2
	return effectiveNumberOfItems * r.unitPrice
}

var _ PricingRule = new(BuyOneGetOneFree)
