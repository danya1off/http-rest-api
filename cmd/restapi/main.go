package main

import (
	"log"

	"github.com/danya1off/http-rest-api/internal/app/config"
	"github.com/danya1off/http-rest-api/internal/app/restapi"
)

func main() {
	config := config.NewConfig()
	api := restapi.New(config)
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}
