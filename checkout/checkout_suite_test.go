package checkout_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCheckout(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Checkout Suite")
}
