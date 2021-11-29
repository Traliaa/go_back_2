package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ENV string

const (
	LocalEnv ENV = "local"
	ProdEnv  ENV = "prod"
	path         = "./config"
)

type Config struct {
	Port string
}

func Load(launchMode ENV) (*Config, error) {
	cfgPath := filepath.Join(path, fmt.Sprintf("%s.yaml", launchMode))
	switch launchMode {
	case LocalEnv:
		cfgPath = filepath.Join(path, fmt.Sprintf("%s.yaml", launchMode))

	case ProdEnv:
		cfgPath = filepath.Join(path, fmt.Sprintf("%s.yaml", launchMode))

	default:
		return nil, fmt.Errorf("unexpected LAUNCH_MODE: [%s]", launchMode)
	}

	file, err := os.ReadFile(cfgPath)
	cfg := Config{}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, fmt.Errorf("load .yaml config file: %w", err)
	}

	return &cfg, nil
}
