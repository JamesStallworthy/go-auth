package goauthdocdisco

import (
	"encoding/json"
)

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
