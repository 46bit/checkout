package pricing_rules_test

import (
	. "github.com/46bit/checkout/pricing_rules"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing/quick"
)

var _ = Describe("Fixed", func() {
	Context("New", func() {
		It("returns a new Fixed", func() {
			Expect(NewFixed(42)).ToNot(BeNil())
		})
	})

	Context("Price", func() {
		It("returns 0 for 0 items", func() {
			price := NewFixed(11).Price(0)
			Expect(price).To(Equal(uint(0)))
		})

		It("returns correctly for 1 item", func() {
			price := NewFixed(77).Price(1)
			Expect(price).To(Equal(uint(77)))
		})

		It("returns correctly for 2 items", func() {
			price := NewFixed(99).Price(2)
			Expect(price).To(Equal(uint(198)))
		})

		It("quickchecks", func() {
			f := func(unitPrice, numberOfItems uint) bool {
				actual := NewFixed(unitPrice).Price(numberOfItems)
				expectation := numberOfItems * unitPrice
				return actual == expectation
			}
			Expect(quick.Check(f, quickConfig)).To(BeNil())
		})
	})
})
