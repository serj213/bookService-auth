package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct {
	Env string `yaml:"env" env-default:"develop"`
	Dsn string `yaml:"dsn" env-required:"true"`
}


func Read() (*Config, error) {
	configPath := flag.String("configPath", "config/local.yaml", "paths configs")

	if *configPath == "" {
		return nil, fmt.Errorf("config empty")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}