package service

import (
	"go-auth/internal/config"
	"go-auth/internal/repository"

	"github.com/golang-jwt/jwt/v4"
)

type AuthenticateService interface {
	GenerateJwtToken(id string, secret string) (string, error)
	RefreshJwtToken(jwt string) (string, error)
	WellKnown() OpenIdConfig
	Jwks() (JwksModel, error)
}

type Claims struct {
	jwt.RegisteredClaims
}

func CreateClientCredentialService(repo repository.AuthenticateRepository, config config.Config) (ClientCredentialService, error) {
	s := ClientCredentialService{AuthRepo: repo}
	err := s.Init(config.KeyLocation, config.Url)
	return s, err
}
