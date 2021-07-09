package configfx

import (
	"go.uber.org/config"
	"os"
)

const defaultFile = "config/config.yaml"

func New() (config.Provider, error) {
	file := defaultFile
	if envfile := os.Getenv("CONFIG_FILE"); envfile != "" {
		file = envfile
	}
	configfile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return config.NewYAML(config.Source(configfile))
}

