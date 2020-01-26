package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type configuration struct {
	Env  string `env:"ENV" envDefault:"DEVELOPMENT"`
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port int    `env:"PORT" envDefault:"5000"`

	RedisAddr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	ReidsPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDatabase int    `env:"REDIS_DATABASEA" envDefault:"0"`
}

var cfg *configuration

func Init() {
	cfg = &configuration{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func GetConfig() *configuration {
	return cfg
}
