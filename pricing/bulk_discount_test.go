package pricing_test

import (
	. "github.com/46bit/checkout/pricing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing/quick"
)

var _ = Describe("BulkDiscount", func() {
	Context("Price", func() {
		var bulk_discount *BulkDiscount

		BeforeEach(func() {
			bulk_discount = &BulkDiscount{
				StandardPrice:        7,
				MinimumNumberOfItems: 2,
				DiscountedPrice:      4,
			}
		})

		It("returns 0 for 0 items", func() {
			price := bulk_discount.Price(0)
			Expect(price).To(Equal(uint(0)))
		})

		It("returns correctly for 1 item", func() {
			price := bulk_discount.Price(1)
			Expect(price).To(Equal(uint(7)))
		})

		It("returns correctly for 2 items", func() {
			price := bulk_discount.Price(2)
			Expect(price).To(Equal(uint(8)))
		})

		It("quickchecks", func() {
			f := func(standardPrice, minimumNumberOfItems, discountedPrice, numberOfItems uint) bool {
				effectivePrice := standardPrice
				if numberOfItems >= minimumNumberOfItems {
					effectivePrice = discountedPrice
				}
				expectation := numberOfItems * effectivePrice
				bd := &BulkDiscount{
					StandardPrice:        standardPrice,
					MinimumNumberOfItems: minimumNumberOfItems,
					DiscountedPrice:      discountedPrice,
				}
				actual := bd.Price(numberOfItems)
				return actual == expectation
			}
			Expect(quick.Check(f, quickConfig)).To(BeNil())
		})
	})
})
