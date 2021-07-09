package loggerfx

import (
	"go.uber.org/config"
	"go.uber.org/zap"
)

const configKey = "log"

func New(config config.Provider) (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()

	err := config.Get(configKey).Populate(&cfg)
	if err != nil {
		return nil, err
	}
	logger, err := cfg.Build()
	if err != nil{
		return nil, err
	}
	return logger, nil
}