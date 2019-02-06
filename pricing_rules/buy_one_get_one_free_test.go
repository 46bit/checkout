package pricing_rules_test

import (
	. "github.com/46bit/checkout/pricing_rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BuyOneGetOneFree", func() {
	Context("New", func() {
		It("returns a new BuyOneGetOneFree", func() {
			Expect(NewBuyOneGetOneFree(7)).ToNot(BeNil())
		})
	})

	Context("Price", func() {
		It("returns 0 for 0 items", func() {
			price := NewBuyOneGetOneFree(7).Price(0)
			Expect(price).To(Equal(uint(0)))
		})

		It("returns correctly for 1 items", func() {
			price := NewBuyOneGetOneFree(7).Price(1)
			Expect(price).To(Equal(uint(7)))
		})

		It("returns correctly for 2 items", func() {
			price := NewBuyOneGetOneFree(7).Price(2)
			Expect(price).To(Equal(uint(7)))
		})

		It("returns correctly for 3 items", func() {
			price := NewBuyOneGetOneFree(7).Price(3)
			Expect(price).To(Equal(uint(14)))
		})

		It("returns correctly for 4 items", func() {
			price := NewBuyOneGetOneFree(7).Price(4)
			Expect(price).To(Equal(uint(14)))
		})

		It("returns correctly for 5 items", func() {
			price := NewBuyOneGetOneFree(7).Price(5)
			Expect(price).To(Equal(uint(21)))
		})
	})
})
