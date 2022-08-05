package service

import "fmt"

type MockService struct {
}

func (s MockService) GenerateJwtToken(id string, secret string) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU4MDY1ODY0fQ._Q9NIu1anMwPzZ3w0gvQbRQVlHRyZUnyd60LzhfNyL0", nil
}

func (s MockService) RefreshJwtToken(jwt string) (string, error) {
	return "ayJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU4MDY1ODY0fQ._Q9NIu1anMwPzZ3w0gvQbRQVlHRyZUnyd60LzhfNyL0", nil
}

func (s MockService) WellKnown() OpenIdConfig {
	issuer := "http://example.com"
	return OpenIdConfig{
		Issuer:                 issuer,
		TokenEndpoint:          fmt.Sprintf("%[1]s/oauth/oauth20/token", issuer),
		JwksUri:                fmt.Sprintf("%[1]s/oauth/jwks", issuer),
		ScopesSupported:        []string{"openid"},
		ResponseTypesSupported: []string{"token"},
		GrantTypesSupported:    []string{"client_credentials"},
		TokenEndpointsEndpointAuthSigningAlgValuesSupported: []string{"RS256"},
	}
}

func (s MockService) Jwks() (JwksModel, error) {
	return JwksModel{
		Keys: []JwkModel{
			{
				X5T: "SOMETHING",
				Use: "Sig",
				Kty: "RSA",
			},
		},
	}, nil
}
