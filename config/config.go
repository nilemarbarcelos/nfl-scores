package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	Port string `env:"PORT" envDefault:"3001"`
}

func Get() (cfg Config) {
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return
}
