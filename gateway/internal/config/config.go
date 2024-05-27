package config

import (
	"errors"
	"os"
)

type Config struct {
	Port             string
	Env              string
	OrchestratorPort string
	MessagePort      string
	TemplatePort     string
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

	orchestratorPort := os.Getenv("ORCHESTRATOR_PORT")
	if orchestratorPort == "" {
		return nil, errors.New("orchestrator service port is not specified")
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
		Port:             os.Getenv("PORT"),
		Env:              os.Getenv("ENV"),
		OrchestratorPort: os.Getenv("ORCHESTRATOR_PORT"),
		MessagePort:      os.Getenv("MESSAGE_PORT"),
		TemplatePort:     os.Getenv("TEMPLATE_PORT"),
	}, nil
}
