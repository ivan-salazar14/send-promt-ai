package main

import (
	"log"

	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/api"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/config"
)

func main() {
	cfg := config.Load()

	if err := api.RunServer(cfg); err != nil {
		log.Fatal(err)
	}
}
