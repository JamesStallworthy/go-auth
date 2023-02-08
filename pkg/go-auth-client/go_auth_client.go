package goauthclient

import (
	"encoding/json"
	"fmt"
	goauthdocdisco "go-auth/pkg/go-auth-doc-disco"
	"io"
	"net/http"
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

type TokenResponse struct {
	Token string `json:"token"`
}

func (c GoAuthClient) LoginClientCredentials(clientId string, clientSecret string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s?clientId=%s&clientSecret=%s", c.Config.TokenEndpoint, clientId, clientSecret))

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("auth service returned: %s", resp.Status)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	tokenResponse := TokenResponse{}
	err2 := json.Unmarshal([]byte(body), &tokenResponse)

	if err2 != nil {
		return "", err
	}

	return tokenResponse.Token, nil
}
