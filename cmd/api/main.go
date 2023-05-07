package main

import (
	"fmt"
	"golang_project_ecommerce/pkg/config"
	"golang_project_ecommerce/pkg/di"
	"golang_project_ecommerce/pkg/verification"
	"log"
)

func main() {
	config, configErr := config.LoadConfig()
	verification.InitTwilio(config)

	fmt.Println(config)
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
