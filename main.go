package main

import (
	"bufio"
	"flag"
	"github.com/46bit/checkout/checkout"
	"github.com/46bit/checkout/payment"
	"github.com/46bit/checkout/pricing"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

func main() {
	var pricingRules string
	flag.StringVar(&pricingRules, "pricing-rules", "pricing-rules.yml", "Filepath of the pricing rules YAML to use")
	var enableTakingPayments bool
	flag.BoolVar(&enableTakingPayments, "enable-taking-payments", false, "Set to true to print payment links using GOV.UK Pay")
	flag.Parse()
	log.Printf("Using pricing rules from '%s'\n", pricingRules)
	var payApiKey string
	if enableTakingPayments {
		payApiKey = os.Getenv("PAY_API_KEY")
		if len(payApiKey) == 0 {
			log.Fatal("Taking payments was enabled but the `PAY_API_KEY` environment variable was not provided")
		}
	}

	pricingRulesYaml, err := ioutil.ReadFile(pricingRules)
	if err != nil {
		log.Fatal(err)
	}
	var config pricing.Config
	err = yaml.UnmarshalStrict([]byte(pricingRulesYaml), &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Available item IDs:")
	for itemID := range config.PricingRules {
		log.Printf("• %s", itemID)
	}

	c := checkout.New(config.PricingRules)
	log.Println("Type one Item ID per line to add it to your basket")
	log.Println("Press Ctrl+C to finish and print the total")

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			err = c.Scan(scanner.Text())
			if err != nil {
				log.Println(err)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Println(c.Total())
			os.Exit(1)
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	<-s
	pricePounds := c.Total() / 100
	pricePence := c.Total() % 100

	log.Printf("Total price: £%d.%02d", pricePounds, pricePence)
	if enableTakingPayments {
		nextUrl, err := payment.PaymentLink(c.Total(), "test-payment", payApiKey)
		if err != nil {
			log.Fatalf("Error generating a payment link: '%s'", err.Error())
		}
		log.Printf("Click here to pay: %s", nextUrl)
	}
	os.Exit(1)
}
