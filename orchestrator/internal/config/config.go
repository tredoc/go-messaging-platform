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
	Port         int
	Env          Environment
	MessageAddr  string
	TemplateAddr string
}

func GetConfig() (Config, error) {
	port := os.Getenv("ORCHESTRATOR_PORT")
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

	messageAddr := os.Getenv("MESSAGE_ADDR")
	if messageAddr == "" {
		return Config{}, errors.New("message service address is not specified")
	}

	templateAddr := os.Getenv("TEMPLATE_ADDR")
	if templateAddr == "" {
		return Config{}, errors.New("template service address is not specified")
	}

	return Config{
		Port:         portInt,
		Env:          Environment(env),
		MessageAddr:  messageAddr,
		TemplateAddr: templateAddr,
	}, nil
}
