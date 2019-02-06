package pricing_rules_test

import (
	. "github.com/46bit/checkout/pricing_rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BulkDiscount", func() {
	Context("New", func() {
		It("returns a new BulkDiscount", func() {
			Expect(NewBulkDiscount(3, 2, 1)).ToNot(BeNil())
		})
	})

	Context("Price", func() {
		var bulk_discount_a BulkDiscount
		var bulk_discount_b BulkDiscount

		BeforeEach(func() {
			bulk_discount_a = NewBulkDiscount(2, 2, 1)
			bulk_discount_b = NewBulkDiscount(3, 7, 4)
		})

		It("returns 0 for 0 items", func() {
			price := bulk_discount_a.Price(0)
			Expect(price).To(Equal(uint(0)))

			price = bulk_discount_b.Price(0)
			Expect(price).To(Equal(uint(0)))
		})

		It("returns correctly for 1 item", func() {
			price := bulk_discount_a.Price(1)
			Expect(price).To(Equal(uint(2)))

			price = bulk_discount_b.Price(1)
			Expect(price).To(Equal(uint(7)))
		})

		It("returns correctly for 2 items", func() {
			price := bulk_discount_a.Price(2)
			Expect(price).To(Equal(uint(2)))

			price = bulk_discount_b.Price(2)
			Expect(price).To(Equal(uint(14)))
		})

		It("returns correctly for 3 items", func() {
			price := bulk_discount_a.Price(3)
			Expect(price).To(Equal(uint(3)))

			price = bulk_discount_b.Price(3)
			Expect(price).To(Equal(uint(12)))
		})
	})
})
