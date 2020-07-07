package main

import (
	"log"

	"github.com/danya1off/http-rest-api/internal/app/config"
	"github.com/danya1off/http-rest-api/internal/app/restapi"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.FromFile()
	logger, err := configureLogger(config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	config.Logger = logger
	config.DBConfig.Logger = logger
	api := restapi.New(config)
	if err := api.Start(); err != nil {
		logger.WithError(err).Fatal("Couldn't start the server!")
	}
}

func configureLogger(level string) (*logrus.Logger, error) {
	l := logrus.New()
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	l.SetLevel(lvl)
	return l, nil
}
