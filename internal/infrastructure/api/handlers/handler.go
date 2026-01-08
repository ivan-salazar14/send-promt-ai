package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ivan-salazar14/send-promt-ai/internal/application/usecases"
)

type AIHandler struct {
	UseCase *usecases.ProcessAIUseCase
}

func (h *AIHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Prompt string `json:"prompt"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// El Caso de Uso maneja la concurrencia internamente
	response, err := h.UseCase.Execute(r.Context(), body.Prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": response})
}
