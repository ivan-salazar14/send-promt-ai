package ports

import "context"

type AIServicePort interface {
	GenerateText(ctx context.Context, prompt string) (string, error)
}
