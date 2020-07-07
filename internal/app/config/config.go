package config

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/danya1off/http-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

var (
	configPath = "config.toml"
)

// Config ...
type Config struct {
	Logger   *logrus.Logger
	DBConfig store.Config
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// FromFile ...
func FromFile() *Config {
	conf := New()
	_, err := toml.DecodeFile(configPath, conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

// New ...
func New() *Config {
	return &Config{
		Logger:   logrus.New(),
		DBConfig: store.NewConfig(),
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
