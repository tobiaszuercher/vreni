package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIKeyID  string `envconfig:"API_KEY_ID"`
	APISecret string `envconfig:"API_KEY_SECRET"`
}

func FromEnvironment() *Config {
	var cfg Config

	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &cfg
}
