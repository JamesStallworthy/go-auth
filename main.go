package main

import (
	"go-auth/internal/api"
	"go-auth/internal/repository"
	"go-auth/internal/service"
	"log"
)

var port string = "5001"
var serv service.AuthenticateService

func main() {
	repo, err := repository.CreateYamlRepository("test_config/clients.yaml")

	if err != nil {
		log.Fatal(err)
		return
	}

	serv, err = service.CreateClientCredentialService(repo)

	if err != nil {
		log.Fatal(err)
		return
	}

	_ = api.CreateRestAPIHandler(serv, port)
}
