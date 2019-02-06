package pricing_rules

type Fixed struct {
	unitPrice uint
}

func NewFixed(unitPrice uint) Fixed {
	return Fixed{unitPrice}
}

func (r Fixed) Price(numberOfItems uint) uint {
	return numberOfItems * r.unitPrice
}

var _ PricingRule = new(Fixed)
