package main

import (
	"bufio"
	"flag"
	"github.com/46bit/checkout/checkout"
	"github.com/46bit/checkout/pricing_rules"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

func main() {
	var pricingRules string
	flag.StringVar(&pricingRules, "pricing-rules", "pricing-rules.yml", "Filepath of the pricing rules YAML to use")
	flag.Parse()
	log.Printf("Using pricing rules from '%s'\n", pricingRules)

	pricingRulesYaml, err := ioutil.ReadFile(pricingRules)
	if err != nil {
		log.Fatal(err)
	}
	var config pricing_rules.Config
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
	log.Printf("\nTotal price: %d", c.Total())
	os.Exit(1)
}
