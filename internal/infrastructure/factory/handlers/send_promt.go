package handlers

import (
	"github.com/ivan-salazar14/send-promt-ai/internal/application/usecases"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/api/handlers"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/config"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/factory/services"
)

func NewAIHandler(cfg *config.Config) *handlers.AIHandler {
	// Inyectamos la config en la factory de servicios
	aiService := services.NewAIService(cfg)

	// Inyectamos valores de la config en el Caso de Uso
	useCase := usecases.NewProcessAIUseCase(aiService, cfg.MaxWorkers, cfg.QueueSize)

	return &handlers.AIHandler{
		UseCase: useCase,
	}
}
