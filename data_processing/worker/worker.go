// Worker wait for data in a queue to transform using a transformer function.
package worker

import (
	"errors"

	"github.com/torbendury/aracotaebwsc/data_processing/queue"
)

var ErrNoProcessorFunctionProvided = errors.New("must provide processor function to worker")

type Worker interface {
	Start()
	Stop()
	Continue()
	Shutdown()
}

type ProductWorker struct {
	inQueue   queue.Queue
	outQueue  queue.Queue
	processor *func(input chan interface{}, output chan interface{}, errs chan error)
}

func New(in queue.Queue, out queue.Queue, processor interface{}) (*ProductWorker, error) {
	procFunc, ok := processor.(*func(input chan interface{}, output chan interface{}, errs chan error))
	if !ok {
		return nil, ErrNoProcessorFunctionProvided
	}
	return &ProductWorker{
		inQueue:   in,
		outQueue:  out,
		processor: procFunc,
	}, nil
}

// Start listening to the inQueue and process stuff with processor, afterwards publish the result to outQueue
// TODO: provide error channel for worker
func (pw *ProductWorker) Start() {
	//go *pw.processor(pw.inQueue, pw.outQueue)
}
