package pricing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"testing/quick"
)

var quickConfig = &quick.Config{MaxCount: 1000}

func TestPricingRules(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PricingRules Suite")
}
