package config

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	vault "github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// LoadConfigs - read the provided config file and unmarshal configs into struct v
func LoadConfigs(fileName string, v interface{}) error {
	viper.SetConfigName(fileName)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return errors.Errorf("failed to read config err: %v", err)
	}

	viper.MustBindEnv("vaultToken", "VAULT_TOKEN")

	vaultToken := viper.GetString("vaultToken")
	vaultPath := viper.GetString("vaultPath")
	vaultURL := viper.GetString("vaultURL")
	vaultKeys := viper.GetStringMapStringSlice("vaultKeys")

	if vaultToken == "" || vaultURL == "" || vaultPath == "" {
		return errors.Errorf("invalid Vault attributes, %s %s %s", vaultToken, vaultURL, vaultPath)
	}

	vaultClient, err := vault.NewClient(&vault.Config{
		Address: vaultURL,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		return errors.Errorf("failed to init Vault: %v", err)
	}

	vaultClient.SetToken(vaultToken)

	for i, ks := range vaultKeys {
		data, err := vaultClient.Logical().Read(fmt.Sprintf("%s/%s", vaultPath, i))
		if err != nil {
			return errors.Errorf("failed to read vault: %v", err)
		}
		d := data.Data["data"].(map[string]interface{})

		for _, k := range ks {
			name := i + strings.ToUpper(k)
			viper.Set(name, d[k])
		}
	}

	if err := viper.Unmarshal(&v); err != nil {
		return err
	}

	return nil
}
