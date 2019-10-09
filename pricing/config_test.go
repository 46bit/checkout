package pricing_test

import (
	. "github.com/46bit/checkout/pricing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
	"reflect"
	"testing/quick"
)

var _ = Describe("Config", func() {
	Context("YAML", func() {
		Context("Marshalling", func() {
			var config Config

			BeforeEach(func() {
				config = NewConfig()
			})

			It("serialises no pricing rules into an empty list of pricing rules", func() {
				data, err := yaml.Marshal(&config)
				Expect(err).To(BeNil())
				Expect(string(data)).To(MatchYAML(`pricing: {}`))
			})

			It("serialises fixed rules", func() {
				data, err := yaml.Marshal(Config{
					PricingRules: map[string]PricingRule{
						"FR1": &Fixed{UnitPrice: 772},
						"SR1": &Fixed{UnitPrice: 1100},
						"JP1": &Fixed{UnitPrice: 450},
					},
				})
				Expect(err).To(BeNil())
				Expect(string(data)).To(MatchYAML(`
          pricing:
            FR1:
              rule: fixed
              unit_price: 772
            SR1:
              rule: fixed
              unit_price: 1100
            JP1:
              rule: fixed
              unit_price: 450
        `))
			})

			It("serialises buy-one-get-one-free rules", func() {
				data, err := yaml.Marshal(Config{
					PricingRules: map[string]PricingRule{
						"blanket": &BuyOneGetOneFree{UnitPrice: 1},
						"pillow":  &BuyOneGetOneFree{UnitPrice: 45},
						"candle":  &BuyOneGetOneFree{UnitPrice: 612},
					},
				})
				Expect(err).To(BeNil())
				Expect(string(data)).To(MatchYAML(`
          pricing:
            blanket:
              rule: buy_one_get_one_free
              unit_price: 1
            pillow:
              rule: buy_one_get_one_free
              unit_price: 45
            candle:
              rule: buy_one_get_one_free
              unit_price: 612
        `))
			})

			It("serialises bulk discount rules", func() {
				data, err := yaml.Marshal(Config{
					PricingRules: map[string]PricingRule{
						"alpha": &BulkDiscount{
							StandardPrice:        1100,
							MinimumNumberOfItems: 4,
							DiscountedPrice:      700,
						},
						"beta": &BulkDiscount{
							StandardPrice:        15,
							MinimumNumberOfItems: 10,
							DiscountedPrice:      9,
						},
						"gamma": &BulkDiscount{
							StandardPrice:        3,
							MinimumNumberOfItems: 7,
							DiscountedPrice:      2,
						},
					},
				})
				Expect(err).To(BeNil())
				Expect(string(data)).To(MatchYAML(`
          pricing:
            alpha:
              rule: bulk_discount
              standard_price: 1100
              minimum_number_of_items: 4
              discounted_price: 700
            beta:
              rule: bulk_discount
              standard_price: 15
              minimum_number_of_items: 10
              discounted_price: 9
            gamma:
              rule: bulk_discount
              standard_price: 3
              minimum_number_of_items: 7
              discounted_price: 2
        `))
			})

			It("serialises assorted rules", func() {
				data, err := yaml.Marshal(Config{
					PricingRules: map[string]PricingRule{
						"bogof": &BuyOneGetOneFree{UnitPrice: 9},
						"bulk": &BulkDiscount{
							StandardPrice:        450,
							MinimumNumberOfItems: 3,
							DiscountedPrice:      399,
						},
						"fixed": &Fixed{UnitPrice: 11},
					},
				})
				Expect(err).To(BeNil())
				Expect(string(data)).To(MatchYAML(`
          pricing:
            bogof:
              rule: buy_one_get_one_free
              unit_price: 9
            bulk:
              rule: bulk_discount
              standard_price: 450
              minimum_number_of_items: 3
              discounted_price: 399
            fixed:
              rule: fixed
              unit_price: 11
        `))
			})
		})

		Context("Unmarshalling", func() {
			var config Config

			BeforeEach(func() {
				config = NewConfig()
			})

			It("deserialises an empty string into no pricing rules", func() {
				err := yaml.UnmarshalStrict([]byte(""), &config)
				Expect(err).To(BeNil())
				Expect(config.PricingRules).To(BeEmpty())
			})

			It("deserialises an empty list of pricing rules", func() {
				data := "pricing:\n"
				err := yaml.UnmarshalStrict([]byte(data), &config)
				Expect(err).To(BeNil())
				Expect(config.PricingRules).To(BeEmpty())
			})

			It("deserialises some fixed pricing rules", func() {
				data := `
        pricing:
          FR1:
            rule: fixed
            unit_price: 772
          SR1:
            rule: fixed
            unit_price: 1100
          JP1:
            rule: fixed
            unit_price: 450
        `
				err := yaml.Unmarshal([]byte(data), &config)
				Expect(err).To(BeNil())
				Expect(config).To(Equal(Config{
					PricingRules: map[string]PricingRule{
						"FR1": &Fixed{UnitPrice: 772},
						"SR1": &Fixed{UnitPrice: 1100},
						"JP1": &Fixed{UnitPrice: 450},
					},
				}))
			})

			It("deserialises some buy one get one free rules", func() {
				data := `
        pricing:
          blanket:
            rule: buy_one_get_one_free
            unit_price: 1
          pillow:
            rule: buy_one_get_one_free
            unit_price: 45
          candle:
            rule: buy_one_get_one_free
            unit_price: 612
        `
				err := yaml.Unmarshal([]byte(data), &config)
				Expect(err).To(BeNil())
				Expect(config).To(Equal(Config{
					PricingRules: map[string]PricingRule{
						"blanket": &BuyOneGetOneFree{UnitPrice: 1},
						"pillow":  &BuyOneGetOneFree{UnitPrice: 45},
						"candle":  &BuyOneGetOneFree{UnitPrice: 612},
					},
				}))
			})

			It("deserialises some bulk discount rules", func() {
				data := `
        pricing:
          alpha:
            rule: bulk_discount
            standard_price: 1100
            minimum_number_of_items: 4
            discounted_price: 700
          beta:
            rule: bulk_discount
            standard_price: 15
            minimum_number_of_items: 10
            discounted_price: 9
          gamma:
            rule: bulk_discount
            standard_price: 3
            minimum_number_of_items: 7
            discounted_price: 2
        `
				err := yaml.Unmarshal([]byte(data), &config)
				Expect(err).To(BeNil())
				Expect(config).To(Equal(Config{
					PricingRules: map[string]PricingRule{
						"alpha": &BulkDiscount{
							StandardPrice:        1100,
							MinimumNumberOfItems: 4,
							DiscountedPrice:      700,
						},
						"beta": &BulkDiscount{
							StandardPrice:        15,
							MinimumNumberOfItems: 10,
							DiscountedPrice:      9,
						},
						"gamma": &BulkDiscount{
							StandardPrice:        3,
							MinimumNumberOfItems: 7,
							DiscountedPrice:      2,
						},
					},
				}))
			})

			It("deserialises assorted rules", func() {
				data := `
        pricing:
          bogof:
            rule: buy_one_get_one_free
            unit_price: 9
          bulk:
            rule: bulk_discount
            standard_price: 450
            minimum_number_of_items: 3
            discounted_price: 399
          fixed:
            rule: fixed
            unit_price: 11
        `
				err := yaml.Unmarshal([]byte(data), &config)
				Expect(err).To(BeNil())
				Expect(config).To(Equal(Config{
					PricingRules: map[string]PricingRule{
						"bogof": &BuyOneGetOneFree{UnitPrice: 9},
						"bulk": &BulkDiscount{
							StandardPrice:        450,
							MinimumNumberOfItems: 3,
							DiscountedPrice:      399,
						},
						"fixed": &Fixed{UnitPrice: 11},
					},
				}))
			})
		})

		Context("quickcheck", func() {
			It("serialises and deserialises random rules", func() {
				f := func(expectedConfig Config) bool {
					data, err := yaml.Marshal(&expectedConfig)
					Expect(err).To(BeNil())

					var actualConfig Config
					err = yaml.UnmarshalStrict([]byte(data), &actualConfig)
					Expect(err).To(BeNil())

					return reflect.DeepEqual(actualConfig, expectedConfig)
				}
				Expect(quick.Check(f, quickConfig)).To(BeNil())
			})
		})
	})
})
