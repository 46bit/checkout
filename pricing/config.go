package pricing

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"math/rand"
	"reflect"
	"testing/quick"
)

type Config struct {
	Pricing map[string]Price `yaml:"pricing"`
}

func NewConfig() Config {
	return Config{
		Pricing: map[string]Price{},
	}
}
