package goauthdocdisco

import "errors"

type OpenIdConfig struct {
	Issuer                                              string   `json:"issuer"`
	TokenEndpoint                                       string   `json:"token_endpoint"`
	JwksUri                                             string   `json:"jwks_uri"`
	ScopesSupported                                     []string `json:"scopes_supported"`
	ResponseTypesSupported                              []string `json:"response_types_supported"`
	GrantTypesSupported                                 []string `json:"grant_types_supported"`
	TokenEndpointsEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported"`
}

func (o OpenIdConfig) Validate() (bool, error) {
	if o.Issuer == "" {
		return false, errors.New("invalid Issuer")
	}

	if o.TokenEndpoint == "" {
		return false, errors.New("invalid Token Endpoint")
	}

	if o.JwksUri == "" {
		return false, errors.New("invalid Jwks endpoint")
	}

	if len(o.ScopesSupported) == 0 {
		return false, errors.New("no supported scopes found")
	}

	if len(o.ResponseTypesSupported) == 0 {
		return false, errors.New("no supported response types")
	}

	if len(o.GrantTypesSupported) == 0 {
		return false, errors.New("no supported grant types")
	}

	if len(o.TokenEndpointsEndpointAuthSigningAlgValuesSupported) == 0 {
		return false, errors.New("no supported token alg values supported")
	}

	return true, nil
}

// `{
// 	"issuer":"%[1]s",
// 	"token_endpoint":"%[1]s/oauth/oauth20/token",
// 	"jwks_uri":"%[1]s/oauth/jwks",
// 	"scopes_supported":[
// 		"openid",
// 	],
// 	"response_types_supported":[
// 		"token",
// 	],
// 	"grant_types_supported":[
// 		"client_credentials",
// 	],
// 	"token_endpoint_auth_signing_alg_values_supported":[
// 		"RS256"
// 	],
// }`
