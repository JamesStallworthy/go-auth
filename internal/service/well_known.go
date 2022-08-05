package service

type OpenIdConfig struct {
	Issuer                                              string   `json:"issuer"`
	TokenEndpoint                                       string   `json:"token_endpoint"`
	JwksUri                                             string   `json:"jwks_uri"`
	ScopesSupported                                     []string `json:"scopes_supported"`
	ResponseTypesSupported                              []string `json:"response_types_supported"`
	GrantTypesSupported                                 []string `json:"grant_types_supported"`
	TokenEndpointsEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported"`
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
