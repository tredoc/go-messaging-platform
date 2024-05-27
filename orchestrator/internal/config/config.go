package config

import (
	"errors"
	"os"
)

type Config struct {
	Port         string
	Env          string
	MessagePort  string
	TemplatePort string
}

func GetConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("server port is not specified")
	}

	env := os.Getenv("ENV")
	if env == "" {
		return nil, errors.New("environment is not specified")
	}

	messagePort := os.Getenv("MESSAGE_PORT")
	if messagePort == "" {
		return nil, errors.New("message service port is not specified")
	}

	templatePort := os.Getenv("TEMPLATE_PORT")
	if templatePort == "" {
		return nil, errors.New("template service port is not specified")
	}

	return &Config{
		Port:         port,
		Env:          env,
		MessagePort:  messagePort,
		TemplatePort: templatePort,
	}, nil
}
