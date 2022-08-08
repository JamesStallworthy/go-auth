package goauthdocdisco

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeOpenIdConfig(t *testing.T) {
	output, err := DecodeOpenIdConfig(`{
		"issuer": "http://localhost",
		"token_endpoint": "http://localhost/oauth/oauth20/token",
		"jwks_uri": "http://localhost/oauth/jwks",
		"scopes_supported": [
			"openid"
		],
		"response_types_supported": [
			"token"
		],
		"grant_types_supported": [
			"client_credentials"
		],
		"token_endpoint_auth_signing_alg_values_supported": [
			"RS256"
		]
	}`)

	assert.Equal(t, "http://localhost", output.Issuer)
	assert.Equal(t, nil, err)
}

func TestDecodeOpenIdConfigInvalidJson(t *testing.T) {
	_, err := DecodeOpenIdConfig(`{
		"something": "true"
	{`)

	assert.Equal(t, "invalid character '{' after object key:value pair", err.Error())
}

func TestDecodeOpenIdConfigMissingIssuer(t *testing.T) {
	_, err := DecodeOpenIdConfig(`{
		"token_endpoint": "http://localhost/oauth/oauth20/token",
		"jwks_uri": "http://localhost/oauth/jwks",
		"scopes_supported": [
			"openid"
		],
		"response_types_supported": [
			"token"
		],
		"grant_types_supported": [
			"client_credentials"
		],
		"token_endpoint_auth_signing_alg_values_supported": [
			"RS256"
		]
	}`)

	assert.Equal(t, "invalid Issuer", err.Error())
}
