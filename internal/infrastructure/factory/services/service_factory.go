package services

import (
	"github.com/ivan-salazar14/send-promt-ai/internal/domain/ports"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/adapters"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/config"
)

func NewAIService(cfg *config.Config) ports.AIServicePort {
	return &adapters.GeminiAdapter{
		APIKey: cfg.OpenAIKey,
	}
}
