package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ivan-salazar14/send-promt-ai/internal/domain/ports"
)

type OpenAIAdapter struct {
	APIKey string
	URL    string
}

// Aseguramos que OpenAIAdapter implementa la interfaz AIServicePort
var _ ports.AIServicePort = (*OpenAIAdapter)(nil)

func (a *OpenAIAdapter) GenerateText(ctx context.Context, prompt string) (string, error) {
	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, "POST", a.URL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+a.APIKey)
	req.Header.Set("Content-Type", "application/json")

	log.Printf("Sending request to OpenAI with prompt: %s", prompt)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	log.Printf("Response status: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("openai api error: status %d", resp.StatusCode)
	}
	log.Printf("Response body: %s", resp.Body)
	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Choices[0].Message.Content, nil
}
