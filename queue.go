package algo

import (
	"fmt"
	"io"
	"sync"
)

// Queue uses an array as the underlying data structure.
type Queue struct {
	mu   sync.RWMutex
	data []any
}

func newQueue() *Queue {
	return &Queue{
		mu:   sync.RWMutex{},
		data: []any{},
	}
}

func (q *Queue) enqueue(v any) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, v)
}

func (q *Queue) dequeue() any {
	if len(q.data) == 0 {
		return nil
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	v := q.data[0]
	// q.data[0] = nil
	q.data = q.data[1:]

	return v
}

func (q *Queue) peak() any {
	if len(q.data) == 0 {
		return nil
	}

	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.data[0]
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Usecase: print manager

type PrintManager struct {
	queue *Queue
}

func newPrintMgr() *PrintManager {
	return &PrintManager{queue: newQueue()}
}

func (p *PrintManager) queuePrintJob(document string) {
	p.queue.enqueue(document)
}

func (p *PrintManager) run(out io.Writer) {
	for len(p.queue.data) > 0 {
		printer(out, p.queue.dequeue().(string))
	}
}

func printer(out io.Writer, doc string) {
	fmt.Fprintln(out, doc)
}
