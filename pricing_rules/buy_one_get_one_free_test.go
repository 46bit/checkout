package pricing_rules_test

import (
	. "github.com/46bit/checkout/pricing_rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing/quick"
)

var _ = Describe("BuyOneGetOneFree", func() {
	Context("Price", func() {
		var buyOneGetOneFree *BuyOneGetOneFree

		BeforeEach(func() {
			buyOneGetOneFree = &BuyOneGetOneFree{UnitPrice: 7}
		})

		It("returns 0 for 0 items", func() {
			price := buyOneGetOneFree.Price(0)
			Expect(price).To(Equal(uint(0)))
		})

		It("returns correctly for 1 items", func() {
			price := buyOneGetOneFree.Price(1)
			Expect(price).To(Equal(uint(7)))
		})

		It("returns correctly for 2 items", func() {
			price := buyOneGetOneFree.Price(2)
			Expect(price).To(Equal(uint(7)))
		})

		It("returns correctly for 3 items", func() {
			price := buyOneGetOneFree.Price(3)
			Expect(price).To(Equal(uint(14)))
		})

		It("quickchecks", func() {
			f := func(unitPrice, numberOfItems uint) bool {
				buyOneGetOneFree := &BuyOneGetOneFree{UnitPrice: unitPrice}
				actual := buyOneGetOneFree.Price(numberOfItems)
				effectiveNumberOfItems := numberOfItems/2 + numberOfItems%2
				expectation := effectiveNumberOfItems * unitPrice
				return actual == expectation
			}
			Expect(quick.Check(f, quickConfig)).To(BeNil())
		})
	})
})
