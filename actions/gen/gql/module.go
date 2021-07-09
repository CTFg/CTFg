package gql

import (
	"go.uber.org/config"
	"go.uber.org/fx"
	"net/http"
)

const configKey = "gqlclient"

type ClientConfig struct {
	BaseURL string
}

func fxNewClient(client *http.Client, cfg config.Provider) (*Client,error){
	var clientConfig ClientConfig
	err := cfg.Get(configKey).Populate(&clientConfig)
	if err != nil{
		return nil, err
	}
	return NewClient(client, clientConfig.BaseURL),nil
}

var Module = fx.Options(
	fx.Provide(
		fxNewClient,
	),
)
