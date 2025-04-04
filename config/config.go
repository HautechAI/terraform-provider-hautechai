package config

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"log"
)

//go:embed config.yml
var configYAML []byte

func GetApiUrl() string {
	apiUrl := viper.GetString("apiUrl")
	if apiUrl == "" {
		log.Fatal("apiUrl is not set in the config file")
	}

	return apiUrl
}

func GetProviderAddress() string {
	providerAddress := viper.GetString("providerAddress")
	if providerAddress == "" {
		log.Fatal("providerAddress is not set in the config file")
	}

	return providerAddress
}

func init() {
	viper.SetConfigFile("./config.yml")

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewReader(configYAML))
	if err != nil {
		log.Fatalf("Failed to load embedded config: %v", err)
	}
}
