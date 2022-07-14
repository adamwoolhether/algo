package algo

import "sync"

// vertex represents a node in a graph.
type vertex[T any] struct {
	mu        sync.RWMutex
	value     T
	neighbors []vertex[T]
}

// NewVertex instantiates a new graph vertex.
func NewVertex[T any](value T) vertex[T] {
	return vertex[T]{
		value:     value,
		neighbors: []vertex[T]{},
	}
}

// AddNeighbor appends a given vertex to the calling vertex's list of neighbors.
func (v *vertex[T]) AddNeighbor(vertex vertex[T]) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.neighbors = append(v.neighbors, vertex)
}
