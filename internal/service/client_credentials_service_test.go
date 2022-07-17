package service

import (
	"go-auth/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Setup() *ClientCredentialService {
	mockRepo := repository.CreateMockRepository()

	mockRepo.CreateClient("Test", "Secret")

	serv := CreateClientCredentialService(mockRepo)
	return &serv
}

func TestGetJwtToken(t *testing.T) {
	serv := Setup()

	token, err := serv.GenerateJwtToken("Test", "Secret")

	assert.NotEmpty(t, token)
	assert.Equal(t, nil, err)
}

func TestGetJwtTokenInvalidClient(t *testing.T) {
	serv := Setup()

	token, err := serv.GenerateJwtToken("Invalid", "Secret")

	assert.Empty(t, token)
	assert.Equal(t, "client credential not found", err.Error())
}

func TestGetJwtTokenInvalidSecret(t *testing.T) {
	serv := Setup()

	token, err := serv.GenerateJwtToken("Test", "Invalid")

	assert.Empty(t, token)
	assert.Equal(t, "client credential not found", err.Error())
}
