package config

import (
	"errors"
	"os"
)

type Config struct {
	Port string
	Env  string
}

func GetConfig() (Config, error) {
	port := os.Getenv("MESSAGE_PORT")
	if port == "" {
		return Config{}, errors.New("server port is not specified")
	}

	env := os.Getenv("ENV")
	if env == "" {
		return Config{}, errors.New("environment is not specified")
	}

	return Config{
		Port: port,
		Env:  env,
	}, nil
}
