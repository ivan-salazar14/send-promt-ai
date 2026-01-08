package api

import (
	"log"
	"net/http"

	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/config"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/factory/handlers"
)

func RunServer(cfg *config.Config) error {
	mux := http.NewServeMux()

	aiHandler := handlers.NewAIHandler(cfg)

	mux.HandleFunc("POST /process", aiHandler.Handle)

	log.Printf("Server starting on port %s with %d workers", cfg.Port, cfg.MaxWorkers)
	return http.ListenAndServe(":"+cfg.Port, mux)
}
