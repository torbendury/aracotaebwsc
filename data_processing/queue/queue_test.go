package queue

import (
	"sync"
	"testing"
)

func MustEqual(want interface{}, got interface{}, t *testing.T) {
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestNew(t *testing.T) {
	_, err := New(-1)
	MustEqual(ErrNonPositiveQueueSize, err, t)
}

func TestPush(t *testing.T) {
	fq, _ := New(3000)
	wg := sync.WaitGroup{}
	for i := 0; i < 3000; i++ {
		wg.Add(1)
		go func() {
			fq.Push("a")
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestPop(t *testing.T) {
	fq, _ := New(3000)

	wg := sync.WaitGroup{}
	for i := 0; i < 3000; i++ {
		wg.Add(1)
		go func() {
			fq.Push("a")
		}()
		go func() {
			wg.Done()
			fq.Pop()
		}()
	}
	wg.Wait()
}

func TestSize(t *testing.T) {
	want := 20
	queueSize := want + 10
	fq, _ := New(queueSize)

	for i := 0; i < want; i++ {
		fq.Push("a")
	}

	got := fq.Size()

	MustEqual(want, got, t)
}

func TestOrder(t *testing.T) {
	want := []int{1, 2, 3}

	fq, _ := New(3)

	for i := 1; i < 4; i++ {
		fq.Push(i)
	}

	got := make([]int, 3)

	for i := 0; i < len(got); i++ {
		elem := fq.Pop()
		elemInt, ok := elem.(int)
		if !ok {
			t.Errorf("could not convert %v to int", elem)
		}
		got[i] = int(elemInt)
	}

	for i, e := range got {
		MustEqual(e, want[i], t)
	}
}

func TestFifoQueueImplementsQueue(t *testing.T) {
	fq, _ := New(1)
	_, ok := interface{}(fq).(Queue)

	MustEqual(true, ok, t)
}
