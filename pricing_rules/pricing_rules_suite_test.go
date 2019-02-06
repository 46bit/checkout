package pricing_rules_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPricingRules(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PricingRules Suite")
}
