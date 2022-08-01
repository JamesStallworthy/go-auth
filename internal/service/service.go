package service

import (
	"go-auth/internal/repository"

	"github.com/golang-jwt/jwt/v4"
)

type AuthenticateService interface {
	GenerateJwtToken(id string, secret string) (string, error)
	RefreshJwtToken(jwt string) (string, error)
}

type Claims struct {
	jwt.RegisteredClaims
}

func CreateClientCredentialService(repo repository.AuthenticateRepository) (ClientCredentialService, error) {
	s := ClientCredentialService{AuthRepo: repo}
	err := s.Init()
	return s, err
}
