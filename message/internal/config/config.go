package config

import (
	"errors"
	"os"
	"strconv"
)

type Environment string

const (
	Development Environment = "dev"
	Production  Environment = "prod"
)

type Config struct {
	Port     int
	Env      Environment
	MongoURI string
}

func GetConfig() (Config, error) {
	port := os.Getenv("MESSAGE_PORT")
	if port == "" {
		return Config{}, errors.New("server port is not specified")
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return Config{}, errors.New("server port is not a number")
	}

	env := os.Getenv("ENV")
	if env == "" {
		return Config{}, errors.New("environment is not specified")
	}

	if Environment(env) != Development && Environment(env) != Production {
		return Config{}, errors.New("environment is not valid")
	}

	mongoURI := os.Getenv("MESSAGE_MONGO_URI")
	if mongoURI == "" {
		return Config{}, errors.New("mongo uri is not specified")
	}

	return Config{
		Port:     portInt,
		Env:      Environment(env),
		MongoURI: mongoURI,
	}, nil
}
