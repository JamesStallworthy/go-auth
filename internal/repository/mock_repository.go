package repository

import (
	"go-auth/internal/models"
)

type MockRepository struct {
	BasicAuthenticateRepository
}

func (r *MockRepository) CreateClient(id string, secret string) {
	r.clients = append(r.clients, models.ClientCredential{Id: id, Secret: secret})
}
