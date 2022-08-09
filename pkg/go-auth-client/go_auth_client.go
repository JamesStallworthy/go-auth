package goauthclient

import (
	goauthdocdisco "go-auth/pkg/go-auth-doc-disco"
)

type GoAuthClient struct {
	authority string
	Config    goauthdocdisco.OpenIdConfig
}

func CreateClient(authority string) (GoAuthClient, error) {
	config, err := goauthdocdisco.RequestConfig(authority)

	if err != nil {
		return GoAuthClient{}, err
	}

	decodedConfig, err := goauthdocdisco.DecodeOpenIdConfig(config)

	if err != nil {
		return GoAuthClient{}, err
	}

	return GoAuthClient{
		authority: authority,
		Config:    decodedConfig,
	}, nil
}

// func (c GoAuthClient) LoginClientCredentials(clientId string, clientSecret string) (string, error) {

// }
