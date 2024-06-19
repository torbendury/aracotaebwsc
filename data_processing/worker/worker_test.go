package worker

import (
	"errors"
	"testing"

	"github.com/torbendury/aracotaebwsc/data_processing/product"
	"github.com/torbendury/aracotaebwsc/data_processing/queue"
	"github.com/torbendury/aracotaebwsc/data_processing/xmljson"
)

func MustEqual(want interface{}, got interface{}, t *testing.T) {
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func spawnQueue() queue.Queue {
	q, _ := queue.New(1)
	return q
}

func processorFunc() func(input chan interface{}, output chan interface{}, errs chan error) {
	return func(input chan interface{}, output chan interface{}, errs chan error) {
		i := <-input
		inputBytes, ok := i.([]byte)
		if !ok {
			errs <- errors.New("no bytes input provided into worker function")
			return
		}
		p := product.Product{}
		res, err := xmljson.XmlToJson(inputBytes, p)

		if err != nil {
			errs <- err
		}
		output <- res
	}
}

func TestNew(t *testing.T) {
	inQueue := spawnQueue()
	outQueue := spawnQueue()
	processor := processorFunc()

	_, err := New(inQueue, outQueue, &processor)

	MustEqual(nil, err, t)
}

func TestStar(t *testing.T) {
	inQueue := spawnQueue()
	outQueue := spawnQueue()
	processor := processorFunc()

	pw, _ := New(inQueue, outQueue, &processor)

	pw.Start()
	// TODO finish test - what should happen?
}
