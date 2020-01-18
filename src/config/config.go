package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Configuration struct {
	Env  string `env:"ENV" envDefault:"DEVELOPMENT"`
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port int    `env:"PORT" envDefault:"5000"`
}

var cfg Configuration

func Init() {
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func GetConfig() Configuration {
	return cfg
}
