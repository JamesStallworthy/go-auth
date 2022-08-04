package service

import (
	"go-auth/internal/repository"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Setup() *ClientCredentialService {
	mockRepo := repository.CreateMockRepository()

	mockRepo.CreateClient("Test", "Secret")

	serv, _ := CreateClientCredentialService(mockRepo, "../../test_certs/")

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

	newToken, err := serv.RefreshJwtToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU5NjM2NzA0fQ.Ns_PZV1vtfmgcd0M7JyIkF11IQT94eAiLh64L9EGCew7XB4QWLmKCfL_CWfLzLVng7vEGgFBioR1hX-Ruq9gylOn5lfMolOzzd-ckhVb27YHTTLNdL6N21rHTfaN1AkBf74V05vd1jAH5oWIbAeSfohxnfCGHlGps4ef9A9P-zHwS3LAtROaDM-IaWRXQUvVgYf1jto1bDlHh5mONdKy9-EtftoH4qIUgmxsajwvi4Y3GKQO32hmqontjWa_IHUBQXSmb08W39PSvlJI1wPSVLPdbNbUvbJigHiyoykFcVSZFzzWp8TBwcIHKFNmjdukrwqUkuhU6qsOrx-3UUtF5Q")

	assert.Empty(t, newToken)
	assert.Contains(t, err.Error(), "token is expired by")
}

func TestRSAKeyGen(t *testing.T) {
	mockRepo := repository.CreateMockRepository()

	mockRepo.CreateClient("Test", "Secret")

	serv, _ := CreateClientCredentialService(mockRepo, "./")

	err := serv.GenerateRSAKey()
	assert.Equal(t, nil, err)

	publicFile, err1 := ioutil.ReadFile("public.pem")
	privateFile, err2 := ioutil.ReadFile("private.pem")

	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)

	assert.Contains(t, string(publicFile), "PUBLIC KEY")
	assert.Contains(t, string(privateFile), "RSA PRIVATE KEY")
}
