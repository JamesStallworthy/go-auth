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

	newToken, err := serv.RefreshJwtToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU5NDc2ODk5fQ.4BJrS-p1oxewhyOT5kf5Zc6gs7B1wK3QdhTofLnPwVIr0gACBEmnxStQn7CC6BjYktaWAImizClId6_Os8RnRML9E2vFA5ZCt2ZevMkFSJqRWFOEXwTfKj-AnwxV6Pxt_VzYm_RkImV0JE50C2Kkg1A6ps8OAOPGCg6VZDiLNPr7JZZYj1dtjgn6DZzFBeJSRI1bXp8NuBlW_YmqJzcgSSN_TB6ztVwKKqoE782y5xYoZ-z20qtUQrfadX1b0PszkMRJJrbBbvUpe_pInkc3qeyH1Ib7uiNqZjPNNlxVQ4lihzliVHjdu1nGTlOrMIMR4d32Tp2kwXfNcZXiyVofPw")

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
