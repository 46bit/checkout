package payment_test

import (
	"github.com/46bit/checkout/payment"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Payment", func() {
	var apiKey string

	BeforeEach(func() {
		apiKey = os.Getenv("PAY_API_KEY")
	})

	Context("PaymentLink", func() {
		It("returns a payment link", func() {
			nextUrl, err := payment.PaymentLink(1199, "test-payment", apiKey)
			Expect(err).To(BeNil())
			Expect(nextUrl).To(HavePrefix("https://www.payments.service.gov.uk/secure/"))
		})
	})
})
