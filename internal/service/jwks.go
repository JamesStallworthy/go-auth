package service

type JwksModel struct {
	Keys []JwkModel `json:"keys"`
}

type JwkModel struct {
	X5T string `json:"x5t"`
	Use string `json:"use"`
	Kty string `json:"kty"`
}
