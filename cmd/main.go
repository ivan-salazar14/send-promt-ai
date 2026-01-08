package main

import (
	"log"

	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/api"
)

func main() {
	log.Println("Starting AI Gateway service...")
	// El main solo invoca el arranque del servidor
	if err := api.RunServer(); err != nil {
		log.Fatal(err)
	}
}
