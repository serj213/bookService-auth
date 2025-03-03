package main

import (
	"log"

	"github.com/serj213/bookAuth/internal/config"
	"github.com/serj213/bookAuth/pkg/pg"
	"go.uber.org/zap"
)

const (
	local = "local"
	dev = "develop"
)


func run() error {
	cfg, err := config.Read()
	if err != nil {
		return err
	}

	log := setupLogger(cfg.Env)
	log = log.With(zap.String("env", cfg.Env))
	log.Info("logger enabled")

	pg, err := pg.Deal(cfg.Dsn)
	if err != nil {
		return err
	}

	_ = pg


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