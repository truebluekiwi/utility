package config

import (
	"testing"

	"github.com/sanity-io/litter"
)

type Config struct {
	Env     string `mapstructure:"env"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`

	ApiUsername string `mapstructure:"retriever-api-username"`
	ApiPassword string `mapstructure:"retriever-api-password"`

	CacheDb string `mapstructure:"retriever-nzx-cache-db"`

	NzxServerHost      string `mapstructure:"retriever-nzx-server-host"`
	NzxServerPort      int    `mapstructure:"retriever-nzx-server-port"`
	NzxUser            string `mapstructure:"retriever-nzx-user"`
	NzxPassword        string `mapstructure:"retriever-nzx-password"`
	InsecureSkipVerify bool   `mapstructure:"insecure-skip-verify"`

	// aws-cloudwatch
	AwsCloudwatchRegion string `mapstructure:"aws-cloudwatch-region"`
	AwsCloudwatchKey    string `mapstructure:"aws-cloudwatch-key"`
	AwsCloudwatchSecret string `mapstructure:"aws-cloudwatch-secret"`

	// aws-dynamo
	AwsDynamoRegion string `mapstructure:"aws-dynamo-region"`
	AwsDynamoKey    string `mapstructure:"aws-dynamo-key"`
	AwsDynamoSecret string `mapstructure:"aws-dynamo-secret"`

	// aws-s3
	AwsS3Region string `mapstructure:"aws-s3-region"`
	AwsS3Key    string `mapstructure:"aws-s3-key"`
	AwsS3Secret string `mapstructure:"aws-s3-secret"`

	// aws-sns
	AwsSNSRegion       string `mapstructure:"aws-sns-region"`
	AwsSNSKey          string `mapstructure:"aws-sns-key"`
	AwsSNSSecret       string `mapstructure:"aws-sns-secret"`
	AwsSNSAlertSendARN string `mapstructure:"aws-sns-alert-send-arn"`

	SlackToken string `mapstructure:"slack-token"`
}

func TestLoadConfigs(t *testing.T) {
	config := &Config{}
	if err := LoadConfigs("config", config); err != nil {
		t.Fatal(err)
	}

	litter.Dump(config)
}
