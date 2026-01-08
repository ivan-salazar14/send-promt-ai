package handlers

import (
	"os"
	"strconv"

	"github.com/ivan-salazar14/send-promt-ai/internal/application/usecases"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/factory/services"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/http/handlers"
)

func NewAIHandler() *handlers.AIHandler {
	aiService := services.NewAIService()

	workers, _ := strconv.Atoi(os.Getenv("MAX_WORKERS"))
	queueSize, _ := strconv.Atoi(os.Getenv("QUEUE_SIZE"))

	useCase := usecases.NewProcessAIUseCase(aiService, workers, queueSize)

	return &handlers.AIHandler{
		UseCase: useCase,
	}
}
