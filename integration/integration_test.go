package integration_test

import (
	"github.com/46bit/checkout/checkout"
	"github.com/46bit/checkout/pricing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checkout", func() {
	var c *checkout.Checkout

	BeforeEach(func() {
		pricingRules := map[string]pricing.PricingRule{
			"FR1": &pricing.BuyOneGetOneFree{UnitPrice: 311},
			"SR1": &pricing.BulkDiscount{
				StandardPrice:        500,
				MinimumNumberOfItems: 3,
				DiscountedPrice:      450,
			},
			"CF1": &pricing.Fixed{UnitPrice: 1123},
		}
		c = checkout.New(pricingRules)
	})

	Context("Scan", func() {
		It("returns an error when scanning an unknown item", func() {
			err := c.Scan("a-fake-product-code")
			Expect(err).ToNot(BeNil())
		})

		It("can scan a known item", func() {
			err := c.Scan("FR1")
			Expect(err).To(BeNil())
		})
	})

	Context("Total", func() {
		It("returns the cost of test data 1 correctly", func() {
			Expect(c.Scan("FR1")).To(BeNil())
			Expect(c.Scan("SR1")).To(BeNil())
			Expect(c.Scan("FR1")).To(BeNil())
			Expect(c.Scan("FR1")).To(BeNil())
			Expect(c.Scan("CF1")).To(BeNil())
			total := c.Total()
			Expect(total).To(Equal(uint(2245)))
		})

		It("returns the cost of test data 2 correctly", func() {
			Expect(c.Scan("FR1")).To(BeNil())
			Expect(c.Scan("FR1")).To(BeNil())
			total := c.Total()
			Expect(total).To(Equal(uint(311)))
		})

		It("returns the cost of test data 3 correctly", func() {
			Expect(c.Scan("SR1")).To(BeNil())
			Expect(c.Scan("SR1")).To(BeNil())
			Expect(c.Scan("FR1")).To(BeNil())
			Expect(c.Scan("SR1")).To(BeNil())
			total := c.Total()
			Expect(total).To(Equal(uint(1661)))
		})
	})
})
