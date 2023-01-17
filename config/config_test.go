package config

import (
	"github.com/sanity-io/litter"
	"testing"
)

type Config struct {
	ENV        string `mapstructure:"env"`
	Port       string `mapstructure:"port"`
	BaseURL    string `mapstructure:"base-url"`
	TestKey    string `mapstructure:"test-key"`
	TestSecret string `mapstructure:"test-secret"`
}

func TestLoadConfigs(t *testing.T) {
	config := &Config{}
	if err := LoadConfigs("config", config); err != nil {
		t.Fatal(err)
	}

	litter.Dump(config)
}
