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

	//for _, path := range strings.Split(vaultKeys, ",") {
	//	secrets, err := vaultClient.Logical().Read(fmt.Sprintf("secret/data/%s", path))
	//	if err != nil {
	//		return errors.Errorf("Failed to read Vault secrets: %v", err)
	//	}
	//	data := secrets.Data["data"].(map[string]interface{})

	//switch path {
	//case "aws-s3":
	//	viper.Set("aws-s3-region", data["region"].(string))
	//	viper.Set("aws-s3-key", data.(map[string]interface{})["key"].(string))
	//	viper.Set("aws-s3-secret", data.(map[string]interface{})["secret"].(string))
	//	viper.Set("aws-s3-email-bucket", data.(map[string]interface{})["email-templates-bucket"].(string))
	//case "aws-ses":
	//	viper.Set("aws-ses-region", data.(map[string]interface{})["region"].(string))
	//	viper.Set("aws-ses-key", data.(map[string]interface{})["key"].(string))
	//	viper.Set("aws-ses-secret", data.(map[string]interface{})["secret"].(string))
	//}
	//}
}
