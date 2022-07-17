package repository

import (
	"errors"
	"go-auth/internal/models"
)

type AuthenticateRepository interface {
	GetClientCredential(id string, secret string) (models.ClientCredential, error)
}

type BasicAuthenticateRepository struct {
	clients []models.ClientCredential
}

func CreateMockRepository() MockRepository {
	return MockRepository{}
}

func CreateYamlRepository(configFile string) (YAMLRepository, error) {
	repo := YAMLRepository{}

	err := repo.Setup(configFile)

	return repo, err
}

func (r BasicAuthenticateRepository) GetClientCredential(id string, secret string) (models.ClientCredential, error) {
	for _, el := range r.clients {
		if el.Id == id && el.Secret == secret {
			return el, nil
		}
	}
	return models.ClientCredential{}, errors.New("client credential not found")
}
