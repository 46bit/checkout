package checkout_test

import (
	"github.com/46bit/checkout/checkout"
	"github.com/46bit/checkout/pricing_rules"
	"github.com/46bit/checkout/pricing_rules/pricing_rulesfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checkout", func() {
	var c *checkout.Checkout

	BeforeEach(func() {
		cake := &pricing_rulesfakes.FakePricingRule{}
		cake.PriceCalls(func(numberOfItems uint) uint { return numberOfItems })
		blanket := &pricing_rulesfakes.FakePricingRule{}
		blanket.PriceCalls(func(numberOfItems uint) uint { return numberOfItems * 3 })
		pricingRules := map[string]pricing_rules.PricingRule{
			"cake":    cake,
			"blanket": blanket,
		}
		c = checkout.New(pricingRules)
	})

	Context("Scan", func() {
		It("returns an error when scanning an unknown item", func() {
			err := c.Scan("a-nonexistent-product-code")
			Expect(err).ToNot(BeNil())
		})

		It("can scan a known item", func() {
			err := c.Scan("cake")
			Expect(err).To(BeNil())
		})
	})

	Context("Total", func() {
		It("returns the cost correctly", func() {
			Expect(c.Scan("blanket")).To(BeNil())
			Expect(c.Scan("cake")).To(BeNil())
			Expect(c.Scan("blanket")).To(BeNil())
			total := c.Total()
			Expect(total).To(Equal(uint(7)))
		})
	})
})
