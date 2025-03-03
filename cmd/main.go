package main

import (
	"log"

	"github.com/serj213/bookAuth/internal/config"
	"go.uber.org/zap"
)

const (
	local = "local"
	dev = "develop"
)


func run() error {
	cfg, err := config.Deal()
	if err != nil {
		return err
	}

	log := setupLogger(cfg.Env)

	log = log.With(zap.String("env", cfg.Env))
	log.Info("logger enabled")

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}


func setupLogger(env string) *zap.SugaredLogger  {

	var log *zap.Logger

	switch(env){
	case local:
		log = zap.Must(zap.NewDevelopment())
	case dev:
		log = zap.Must(zap.NewProduction())
	default:
		log = zap.Must(zap.NewProduction())
	}

	return log.Sugar()
}