package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"runtime/debug"
)

type config struct {
	Port    string `envconfig:"PORT" default:"8080"`
	Timeout int    `envconfig:"TIMEOUT" default:"30"`
	AppName string `envconfig:"APP_NAME" default:""`
}

var Config *config

func NewConfig() error {
	if Config != nil {
		return nil
	}

	cfg := &config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return fmt.Errorf("failed to process enviroment variables: %w", err)
	}

	if cfg.AppName == "" {
		info, _ := debug.ReadBuildInfo()
		cfg.AppName = info.Main.Path
	}

	Config = cfg
	return nil
}
