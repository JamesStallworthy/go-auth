package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"go-auth/internal/repository"
	"os"
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

	tokenString, err := generateJwtTokenImpl()

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func generateJwtTokenImpl() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func (s ClientCredentialService) RefreshJwtToken(tokenString string) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", errors.New("token is not valid")
	}

	newTokenString, err2 := generateJwtTokenImpl()

	if err2 != nil {
		return "", err
	}

	return newTokenString, nil
}

func (s ClientCredentialService) GenerateRSAKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	publicKey := &privateKey.PublicKey

	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privatePemWriter, err := os.Create("private.pem")

	if err != nil {
		return err
	}

	err = pem.Encode(privatePemWriter, privateKeyBlock)

	if err != nil {
		return err
	}

	var publicKeyBytes []byte = x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicPemWriter, err := os.Create("public.pem")

	if err != nil {
		return err
	}

	err = pem.Encode(publicPemWriter, publicKeyBlock)

	if err != nil {
		return err
	}

	return nil
}
