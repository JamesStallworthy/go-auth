package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlRepo(t *testing.T) {
	yamlRepo, err := CreateYamlRepository("../../test_config/clients.yaml")

	assert.Equal(t, nil, err)

	client, err := yamlRepo.GetClientCredential("TempuratureClient", "SuperSecretSecret")

	assert.Equal(t, nil, err)
	assert.Equal(t, "TempuratureClient", client.Id)
	assert.Equal(t, "SuperSecretSecret", client.Secret)
}

func TestYamlRepoInvalidId(t *testing.T) {
	yamlRepo, err := CreateYamlRepository("../../test_config/clients.yaml")

	assert.Equal(t, nil, err)

	client, err := yamlRepo.GetClientCredential("Invalid", "SuperSecretSecret")

	assert.Equal(t, "client credential not found", err.Error())
	assert.Equal(t, "", client.Id)
}

func TestYamlRepoInvalidSecret(t *testing.T) {
	yamlRepo, err := CreateYamlRepository("../../test_config/clients.yaml")

	assert.Equal(t, nil, err)

	client, err := yamlRepo.GetClientCredential("TempuratureClient", "invalid")

	assert.Equal(t, "client credential not found", err.Error())
	assert.Equal(t, "", client.Id)
}

func TestYamlRepoInvalidFile(t *testing.T) {
	_, err := CreateYamlRepository("invalid.yaml")

	assert.Equal(t, "open invalid.yaml: no such file or directory", err.Error())
}
