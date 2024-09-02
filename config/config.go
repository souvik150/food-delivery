package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port               string        `mapstructure:"PORT"`
	DBHost             string        `mapstructure:"POSTGRES_HOST"`
	DBUserName         string        `mapstructure:"POSTGRES_USER"`
	DBUserPassword     string        `mapstructure:"POSTGRES_PASSWORD"`
	DBName             string        `mapstructure:"POSTGRES_DB"`
	DBPort             string        `mapstructure:"POSTGRES_PORT"`
	ClientOrigin       string        `mapstructure:"CLIENT_ORIGIN"`
	AWSBucketName      string        `mapstructure:"AWS_BUCKET_NAME"`
	AWSRegion          string        `mapstructure:"AWS_REGION"`
	AWSAccessKey       string        `mapstructure:"AWS_ACCESS_KEY"`
	AWSSecretKey       string        `mapstructure:"AWS_SECRET_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}