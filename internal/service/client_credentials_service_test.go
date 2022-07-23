package service

import (
	"go-auth/internal/repository"
	"testing"
	"time"

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

func TestRefreshJwtToken(t *testing.T) {
	serv := Setup()

	token, err := serv.GenerateJwtToken("Test", "Secret")

	assert.NotEmpty(t, token)
	assert.Equal(t, nil, err)

	time.Sleep(1 * time.Second)

	newToken, err := serv.RefreshJwtToken(token)

	assert.NotEmpty(t, newToken)
	assert.NotEqual(t, token, newToken)
	assert.Equal(t, nil, err)
}

func TestRefreshJwtTokenInvalidToken(t *testing.T) {
	serv := Setup()

	newToken, err := serv.RefreshJwtToken("ayJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU4MDY1ODY0fQ._Q9NIu1anMwPzZ3w0gvQbRQVlHRyZUnyd60LzhfNyL0")

	assert.Empty(t, newToken)
	assert.Contains(t, err.Error(), "invalid character 'k'")
}

func TestRefreshJwtTokenExpiredToken(t *testing.T) {
	serv := Setup()

	newToken, err := serv.RefreshJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU4MDY1ODY0fQ._Q9NIu1anMwPzZ3w0gvQbRQVlHRyZUnyd60LzhfNyL0")

	assert.Empty(t, newToken)
	assert.Contains(t, err.Error(), "token is expired by")
}
