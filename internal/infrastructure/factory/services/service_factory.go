package services

import (
	"os"

	"github.com/ivan-salazar14/send-promt-ai/internal/domain/ports"
	"github.com/ivan-salazar14/send-promt-ai/internal/infrastructure/adapters"
)

func NewAIService() ports.AIServicePort {
	return &adapters.OpenAIClient{
		APIKey: os.Getenv("OPENAI_API_KEY"),
		URL:    "https://api.openai.com/v1/chat/completions",
	}
}
