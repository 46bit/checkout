package pricing_rules

//go:generate counterfeiter . PricingRule
type PricingRule interface {
	Price(numberOfItems uint) uint
}
