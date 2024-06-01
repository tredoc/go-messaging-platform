package config

import (
	"errors"
	"os"
)

type Config struct {
	Port             string
	Env              string
	OrchestratorAddr string
	MessageAddr      string
	TemplateAddr     string
}

func GetConfig() (*Config, error) {
	port := os.Getenv("GATEWAY_PORT")
	if port == "" {
		return nil, errors.New("server port is not specified")
	}

	env := os.Getenv("ENV")
	if env == "" {
		return nil, errors.New("environment is not specified")
	}

	orchestratorAddr := os.Getenv("ORCHESTRATOR_ADDR")
	if orchestratorAddr == "" {
		return nil, errors.New("orchestrator service address is not specified")
	}

	messageAddr := os.Getenv("MESSAGE_ADDR")
	if messageAddr == "" {
		return nil, errors.New("message service address is not specified")
	}

	templateAddr := os.Getenv("TEMPLATE_ADDR")
	if templateAddr == "" {
		return nil, errors.New("template service address is not specified")
	}

	return &Config{
		Port:             port,
		Env:              env,
		OrchestratorAddr: orchestratorAddr,
		MessageAddr:      messageAddr,
		TemplateAddr:     templateAddr,
	}, nil
}
