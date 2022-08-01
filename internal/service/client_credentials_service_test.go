package service

import (
	"go-auth/internal/repository"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Setup() *ClientCredentialService {
	mockRepo := repository.CreateMockRepository()

	mockRepo.CreateClient("Test", "Secret")

	serv, _ := CreateClientCredentialService(mockRepo)
	serv.Init()
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

	newToken, err := serv.RefreshJwtToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU5Mzg3MTQxfQ.tKBKpfjtE938TvqwcQyFA7WzjPIvf9bD_JcMYltqQ2AoZkaDBKahnmJ2YobIQ3P4k7jUWtlsMtY507l08jQFMdhuYQYUDQtaEeaWRzVWy7osqgRXXc44CVfRVfNq9oAuKzsByisUqJvvtdimmYKkmcuVwkdqp-1Md1deVcqUvc_r4KpsFgQgW5k7uZ3PuwJ7jyNPefldltDfl-qN-tw0XwY5PdeTMjhrB8OpcBIXF17n_bIZVwTi99gWTSzxe6TfknYw_XjNXcje_xWdlQE6WZmlVLUArM5eJcbVeCS2f6J6lJArgGeGsojHewY0k9l3ujd4vdteZQg41T3Ssupsiw")

	assert.Empty(t, newToken)
	assert.Contains(t, err.Error(), "token is expired by")
}

func TestRSAKeyGen(t *testing.T) {
	serv := Setup()

	err := serv.GenerateRSAKey()
	assert.Equal(t, nil, err)

	publicFile, err1 := ioutil.ReadFile("public.pem")
	privateFile, err2 := ioutil.ReadFile("private.pem")

	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)

	assert.Contains(t, string(publicFile), "PUBLIC KEY")
	assert.Contains(t, string(privateFile), "RSA PRIVATE KEY")

	os.Remove("private.pem")
	os.Remove("public.pem")
}
