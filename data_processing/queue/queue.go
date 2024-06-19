package queue

import "errors"

type Queue interface {
	//New(int) (*Queue, error)
	Push(interface{})
	Pop() interface{}
	Size() int
}

var ErrNonPositiveQueueSize = errors.New("queue size must be greater than zero")

type FifoQueue struct {
	queue chan interface{}
}

func New(size int) (Queue, error) {
	if size <= 0 {
		return nil, ErrNonPositiveQueueSize
	}
	return &FifoQueue{
		queue: make(chan interface{}, size),
	}, nil
}

func (fq *FifoQueue) Push(elem interface{}) {
	fq.queue <- elem
}

func (fq *FifoQueue) Pop() interface{} {
	return <-fq.queue
}

func (fq *FifoQueue) Size() int {
	return len(fq.queue)
}
