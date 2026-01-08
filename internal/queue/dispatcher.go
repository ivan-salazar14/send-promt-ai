package queue

import (
	"context"
)

// Job represents a unit of work to be processed by a worker
type Job struct {
	Payload string
	Result  chan JobResult
	Ctx     context.Context
}

// JobResult carries the AI response or an error back to the handler
type JobResult struct {
	Response string
	Err      error
}

// Dispatcher manages the worker pool
type Dispatcher struct {
	jobQueue   chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int, queueSize int) *Dispatcher {
	return &Dispatcher{
		jobQueue:   make(chan Job, queueSize),
		maxWorkers: maxWorkers,
	}
}

// Start initializes the workers
func (d *Dispatcher) Start(workerFunc func(Job)) {
	for i := 0; i < d.maxWorkers; i++ {
		go func() {
			for job := range d.jobQueue {
				workerFunc(job)
			}
		}()
	}
}

// Submit pushes a new job into the queue
func (d *Dispatcher) Submit(job Job) {
	d.jobQueue <- job
}
