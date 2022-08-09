package goauthclient

import (
	goauthdocdisco "go-auth/pkg/go-auth-doc-disco"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestCreateClient(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("http://example.com").
		Get("/.well-known/openid-configuration").
		Reply(200).
		JSON(goauthdocdisco.OpenIdConfig{
			Issuer:                 "http://example.com",
			TokenEndpoint:          "http://example.com/oauth/oauth20/token",
			JwksUri:                "http://example.com/oauth/jwks",
			ScopesSupported:        []string{"openid"},
			GrantTypesSupported:    []string{"client_credentials"},
			ResponseTypesSupported: []string{"token"},
			TokenEndpointsEndpointAuthSigningAlgValuesSupported: []string{"RS256"},
		})

	client, err := CreateClient("http://example.com")

	assert.Equal(t, nil, err)
	assert.Equal(t, "http://example.com", client.Config.Issuer)
}

func TestCreateClientInvalid(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("http://example.com").
		Get("/.well-known/openid-configuration").
		Reply(200).
		JSON(goauthdocdisco.OpenIdConfig{
			//Issuer:                 "http://example.com",
			TokenEndpoint:          "http://example.com/oauth/oauth20/token",
			JwksUri:                "http://example.com/oauth/jwks",
			ScopesSupported:        []string{"openid"},
			GrantTypesSupported:    []string{"client_credentials"},
			ResponseTypesSupported: []string{"token"},
			TokenEndpointsEndpointAuthSigningAlgValuesSupported: []string{"RS256"},
		})

	_, err := CreateClient("http://example.com")

	assert.Equal(t, "invalid Issuer", err.Error())
}
