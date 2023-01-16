package config

import (
	"testing"
)

func TestLoadConfigs(t *testing.T) {
	if err := LoadConfigs("config"); err != nil {
		t.Fatal(err)
	}
}
