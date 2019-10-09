package pricing_test

import (
	. "github.com/46bit/checkout/pricing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing/quick"
)

var _ = Describe("Fixed", func() {
	Context("Price", func() {
		var fixed *Fixed

		BeforeEach(func() {
			fixed = &Fixed{UnitPrice: 11}
		})

		It("returns 0 for 0 items", func() {
			price := fixed.Price(0)
			Expect(price).To(Equal(uint(0)))
		})

		It("returns correctly for 1 item", func() {
			fixed.UnitPrice = 77
			price := fixed.Price(1)
			Expect(price).To(Equal(uint(77)))
		})

		It("returns correctly for 2 items", func() {
			fixed.UnitPrice = 99
			price := fixed.Price(2)
			Expect(price).To(Equal(uint(198)))
		})

		It("quickchecks", func() {
			f := func(unitPrice, numberOfItems uint) bool {
				fixed := &Fixed{UnitPrice: unitPrice}
				actual := fixed.Price(numberOfItems)
				expectation := numberOfItems * unitPrice
				return actual == expectation
			}
			Expect(quick.Check(f, quickConfig)).To(BeNil())
		})
	})
})
