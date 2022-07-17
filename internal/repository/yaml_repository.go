package repository

import (
	"go-auth/internal/models"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type YAMLRepository struct {
	BasicAuthenticateRepository
}

type YAMLClients struct {
	Id     string
	Secret string
}

func (y *YAMLRepository) Setup(configFile string) error {
	yfile, err := ioutil.ReadFile(configFile)

	if err != nil {
		return err
	}

	data := make(map[string]YAMLClients)

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		return err
	}

	for _, v := range data {
		y.clients = append(y.clients, models.ClientCredential{Id: v.Id, Secret: v.Secret})
	}

	println(len(y.clients))
	return nil
}
