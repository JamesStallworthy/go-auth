package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"go-auth/internal/repository"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ClientCredentialService struct {
	AuthRepo           repository.AuthenticateRepository
	privateKey         []byte
	publicKey          []byte
	publicKeyLocation  string
	privateKeyLocation string
}

func (s *ClientCredentialService) Init(keyLocation string) error {
	s.publicKeyLocation = filepath.Join(keyLocation, "/public.pem")
	s.privateKeyLocation = filepath.Join(keyLocation, "/private.pem")

	if _, err := os.Stat(s.privateKeyLocation); err == nil {
		err = s.LoadKeys()
		if err != nil {
			return err
		}

		return nil
	} else if errors.Is(err, os.ErrNotExist) {
		err := s.GenerateRSAKey()
		if err != nil {
			return err
		}
		s.LoadKeys()
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func (s *ClientCredentialService) LoadKeys() error {
	var err error
	s.privateKey, err = os.ReadFile(s.privateKeyLocation)

	if err != nil {
		return err
	}

	s.publicKey, err = os.ReadFile(s.publicKeyLocation)

	if err != nil {
		return err
	}
	return nil
}

func (s ClientCredentialService) GenerateJwtToken(id string, secret string) (string, error) {
	_, err := s.AuthRepo.GetClientCredential(id, secret)

	if err != nil {
		return "", err
	}

	tokenString, err := s.generateJwtTokenImpl()

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s ClientCredentialService) generateJwtTokenImpl() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	key, err := jwt.ParseRSAPrivateKeyFromPEM(s.privateKey)

	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func (s ClientCredentialService) RefreshJwtToken(tokenString string) (string, error) {
	claims := &Claims{}

	key, err := jwt.ParseRSAPublicKeyFromPEM(s.publicKey)

	if err != nil {
		return "", err
	}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", errors.New("token is not valid")
	}

	newTokenString, err2 := s.generateJwtTokenImpl()

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

	privatePemWriter, err := os.Create(s.privateKeyLocation)

	if err != nil {
		return err
	}

	err = pem.Encode(privatePemWriter, privateKeyBlock)

	if err != nil {
		return err
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicPemWriter, err := os.Create(s.publicKeyLocation)

	if err != nil {
		return err
	}

	err = pem.Encode(publicPemWriter, publicKeyBlock)

	if err != nil {
		return err
	}

	return nil
}
