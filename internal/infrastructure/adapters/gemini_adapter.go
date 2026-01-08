package adapters

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/ivan-salazar14/send-promt-ai/internal/domain/ports"
	"google.golang.org/api/option"
)

type GeminiAdapter struct {
	APIKey string
}

var _ ports.AIServicePort = (*GeminiAdapter)(nil)

func (a *GeminiAdapter) GenerateText(ctx context.Context, prompt string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(a.APIKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	var content string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if txt, ok := part.(genai.Text); ok {
					content += string(txt)
				}
			}
		}
	}

	log.Printf("Response from Gemini: %s", content)

	return content, nil
}
