package main

import (
	"go-auth/internal/api"
	"go-auth/internal/config"
	"go-auth/internal/repository"
	"go-auth/internal/service"
	"log"
)

var port string = "5001"
var serv service.AuthenticateService

func main() {
	config := config.LoadConfig("config.yaml")

	repo, err := repository.CreateYamlRepository(config.ClientConfigLocation)

	if err != nil {
		log.Fatal(err)
		return
	}

	serv, err = service.CreateClientCredentialService(repo, config.KeyLocation)

	if err != nil {
		log.Fatal(err)
		return
	}

	_ = api.CreateRestAPIHandler(serv, port)
}
