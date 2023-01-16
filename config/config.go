package config

import (
	"net/http"
	"time"

	vault "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func LoadConfigs(fileName string) error {
	viper.SetConfigName(fileName)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return errors.Errorf("ReadInConfig err: %v", err.Error())
	}

	viper.MustBindEnv("vault-token", "VAULT_TOKEN")

	vaultToken := viper.GetString("vault-token")
	vaultURL := viper.GetString("vault-url")
	vaultKeys := viper.GetString("vault-keys")

	if vaultToken == "" || vaultURL == "" || vaultKeys == "" {
		return errors.Errorf("Invalid Vault attributes, %s %s %s", vaultURL, vaultKeys, vaultToken)
	}

	vaultClient, err := vault.NewClient(&vault.Config{
		Address: vaultURL,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		return errors.Errorf("Failed to init Vault: %v", err)
	}

	vaultClient.SetToken(vaultToken)

	return nil
}
