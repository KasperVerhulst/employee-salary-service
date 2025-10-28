package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port            int    `envconfig:"PORT" default:"8080"`
	VerifySignature bool   `envconfig:"VERIFY_SIGNATURE" default:"false"`
	JWKURI          string `envconfig:"JWK_URI"`

	DefaultPageLimit int `envconfig:"DEFAULT_PAGE_LIMIT" default:"10"`
	MaxPageLimit     int `envconfig:"MAX_PAGE_LIMIT" default:"200"`
}

func NewConfig() *Config {

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
