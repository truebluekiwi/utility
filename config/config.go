package config

import (
    "fmt"
    "net/http"
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

    viper.MustBindEnv("vault-token", "VAULT_TOKEN")

    vaultToken := viper.GetString("vault-token")
    vaultPath := viper.GetString("vault-path")
    vaultURL := viper.GetString("vault-url")
    vaultKeys := viper.GetStringMapStringSlice("vault-keys")

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
        value := fmt.Sprintf("%s/%s", vaultPath, i)
        data, err := vaultClient.Logical().Read(value)
        if err != nil {
            return errors.Errorf("failed to read vault value %s, err: %v", value, err)
        }
        d := data.Data["data"].(map[string]interface{})

        for _, k := range ks {
            name := i + "-" + k
            viper.Set(name, d[k])
        }
    }

    if err := viper.Unmarshal(&v); err != nil {
        return err
    }

    return nil
}
