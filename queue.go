package algo

import (
	"fmt"
	"io"
	"sync"
)

// Queue uses an array as the underlying data structure.
type Queue[T any] struct {
	mu   sync.RWMutex
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		mu:   sync.RWMutex{},
		data: []T{},
	}
}

func (q *Queue[T]) enqueue(v T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, v)
}

func (q *Queue[T]) dequeue() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isEmpty() {
		return q.nilType()
	}

	v := q.data[0]
	// q.data[0] = nil
	q.data = q.data[1:]

	return v
}

func (q *Queue[T]) peak() T {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.isEmpty() {
		return q.nilType()
	}

	return q.data[0]
}

func (q *Queue[T]) isEmpty() bool {
	// no locking, as this is internal and should be done by caller.
	return len(q.data) == 0
}

func (q *Queue[T]) nilType() T {
	var t T

	return t
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Usecase: print manager

type PrintManager struct {
	queue *Queue[string]
}

func newPrintMgr() *PrintManager {
	return &PrintManager{queue: NewQueue[string]()}
}

func (p *PrintManager) queuePrintJob(document string) {
	p.queue.enqueue(document)
}

func (p *PrintManager) run(out io.Writer) {
	for len(p.queue.data) > 0 {
		printer(out, p.queue.dequeue())
	}
}

func printer(out io.Writer, doc string) {
	fmt.Fprintln(out, doc)
}
