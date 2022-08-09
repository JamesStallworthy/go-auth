package goauthdocdisco

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RequestConfig(authorityUrl string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", authorityUrl, "/.well-known/openid-configuration"))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func DecodeOpenIdConfig(s string) (OpenIdConfig, error) {
	openId := OpenIdConfig{}
	err := json.Unmarshal([]byte(s), &openId)

	if err != nil {
		return OpenIdConfig{}, err
	}

	valid, err := openId.Validate()

	if !valid {
		return openId, err
	}

	return openId, nil
}
