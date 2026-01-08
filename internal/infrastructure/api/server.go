package api

import (
	"log"
	"net/http"
	"os"

	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/factory/handlers"
)

func RunServer() error {
	// Usamos el ServeMux de la librería estándar (Go 1.22+)
	mux := http.NewServeMux()

	// Inicializamos el handler mediante su factory
	aiHandler := handlers.NewAIHandler()

	// Definimos la ruta
	mux.HandleFunc("POST /process", aiHandler.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	return http.ListenAndServe(":"+port, mux)
}
