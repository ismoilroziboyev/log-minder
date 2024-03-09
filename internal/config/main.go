package config

import (
	"os"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Load() *Config {
	var cfg Config

	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		logrus.Fatal("cannot find .env file to load environment variables")
	}

	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		logrus.Fatalf("error occured while parsing additional configs from environment: %s", err.Error())
	}

	return &cfg
}
