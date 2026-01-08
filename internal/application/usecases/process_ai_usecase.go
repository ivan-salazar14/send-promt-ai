package usecases

import (
	"context"
	"log"

	"github.com/ivan-salazar14/send-promt-ai/internal/domain/ports"
)

type ProcessAIUseCase struct {
	aiService  ports.AIServicePort
	jobQueue   chan job
	maxWorkers int
}

type job struct {
	prompt string
	result chan string
	err    chan error
	ctx    context.Context
}

func NewProcessAIUseCase(ai ports.AIServicePort, maxWorkers, queueSize int) *ProcessAIUseCase {
	uc := &ProcessAIUseCase{
		aiService:  ai,
		jobQueue:   make(chan job, queueSize),
		maxWorkers: maxWorkers,
	}
	uc.startWorkers()
	return uc
}

func (uc *ProcessAIUseCase) startWorkers() {
	for i := 0; i < uc.maxWorkers; i++ {
		go func() {
			for j := range uc.jobQueue {
				log.Printf("Worker processing prompt: %s", j.prompt)
				res, err := uc.aiService.GenerateText(j.ctx, j.prompt)
				j.result <- res
				j.err <- err
			}
		}()
	}
}

func (uc *ProcessAIUseCase) Execute(ctx context.Context, prompt string) (string, error) {
	resChan := make(chan string)
	errChan := make(chan error)

	log.Printf("Queuing prompt: %s", prompt)
	uc.jobQueue <- job{prompt: prompt, result: resChan, err: errChan, ctx: ctx}

	select {
	case res := <-resChan:
		return res, <-errChan
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
