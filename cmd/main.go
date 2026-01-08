package main

import (
	"log"

	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/api"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/config"
)

func main() {
	// 1. Cargar configuración al inicio
	cfg := config.Load()

	// 2. Ejecutar servidor pasando la configuración
	if err := api.RunServer(cfg); err != nil {
		log.Fatal(err)
	}
}
