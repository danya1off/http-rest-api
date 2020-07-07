package store

import "github.com/sirupsen/logrus"

// Config ...
type Config struct {
	Logger      *logrus.Logger
	DatabaseURL string `toml:"db_url"`
}

// NewConfig ...
func NewConfig() Config {
	return Config{
		Logger: logrus.New(),
	}
}
