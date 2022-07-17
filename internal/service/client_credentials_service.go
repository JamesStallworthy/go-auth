package service

import (
	"go-auth/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ClientCredentialService struct {
	AuthRepo repository.AuthenticateRepository
}

func (s ClientCredentialService) GenerateJwtToken(id string, secret string) (string, error) {
	_, err := s.AuthRepo.GetClientCredential(id, secret)

	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}
