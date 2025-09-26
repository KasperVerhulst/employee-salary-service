package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port            int    `envconfig:"PORT" default:"8080"`
	VerifySignature bool   `envconfig:"VERIFY_SIGNATURE" default:"false"`
	JWKURI          string `envconfig:"JWK_URI"`
}

func NewConfig() *Config {

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
