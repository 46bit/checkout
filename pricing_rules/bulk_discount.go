package pricing_rules

type BulkDiscount struct {
	minimumNumberOfItems uint
	standardPrice        uint
	discountedPrice      uint
}

func NewBulkDiscount(minimumNumberOfItems, standardPrice, discountedPrice uint) BulkDiscount {
	return BulkDiscount{
		minimumNumberOfItems,
		standardPrice,
		discountedPrice,
	}
}

func (r BulkDiscount) Price(numberOfItems uint) uint {
	unitPrice := r.standardPrice
	if numberOfItems >= r.minimumNumberOfItems {
		unitPrice = r.discountedPrice
	}
	return numberOfItems * unitPrice
}

var _ PricingRule = new(BulkDiscount)
