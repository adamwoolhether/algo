package algo

import (
	"fmt"
	"io"
	"sync"

	"golang.org/x/exp/constraints"
)

// vertex represents a node in a graph.
type vertex[T constraints.Ordered] struct {
	mu        sync.RWMutex
	value     T
	neighbors []*vertex[T]
}

// NewVertex instantiates a new graph vertex.
func NewVertex[T constraints.Ordered](value T) *vertex[T] {
	return &vertex[T]{
		value:     value,
		neighbors: []*vertex[T]{},
	}
}

// AddNeighbor appends a given vertex to the calling vertex's list of neighbors.
func (v *vertex[T]) AddNeighbor(vertex *vertex[T]) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.neighbors = append(v.neighbors, vertex)
}

// DepthFirstTraversal Uses the Depth First Search algorith to traverse over all
// neighbors of a vertex, printout their values out to the given io.Writer.
func (v *vertex[T]) DepthFirstTraversal(writer io.Writer) {
	visitedVertices := make(map[T]bool)

	var recurse func(vert *vertex[T], m map[T]bool)
	recurse = func(vert *vertex[T], m map[T]bool) {
		// Mark the vertex as visited.
		m[v.value] = true

		// Print the vertex's value.
		fmt.Fprintf(writer, "%v\n", vert.value)

		// Iterate through the current vertex's neighbors.
		for _, neighbor := range vert.neighbors {
			// Ignore the neighbor if it's been bisited already.
			if m[neighbor.value] {
				continue
			}

			// Recursively call this method on the neighbor.
			recurse(neighbor, m)
		}
	}

	recurse(v, visitedVertices)
}

// DepthFirstSearch uses the Depth First algorithm to search
// far a node that contains the given value.
func (v *vertex[T]) DepthFirstSearch(value T) *vertex[T] {
	// Return if given search value is the caller's own.
	if v.value == value {
		return v
	}

	visitedVertices := make(map[T]bool)

	var recurse func(vert *vertex[T], val T, m map[T]bool) *vertex[T]
	recurse = func(vert *vertex[T], val T, m map[T]bool) *vertex[T] {
		// Mark the vertex as visited.
		m[vert.value] = true

		// Iterate through the current vertex's neighbors.
		for _, neighbor := range v.neighbors {
			// Ignore the neighbor if it's been bisited already.
			if m[neighbor.value] {
				continue
			}

			// If the neighbor is the targeted vertex, return it.
			if neighbor.value == val {
				return neighbor
			}

			// Try finding the target vertex with recursion on the adjacent vertex,
			// retuning it if found. Otherwise, `nil` will be returned.
			return recurse(neighbor, val, m)

			/*// If found, return the target vertex.
			if targetVertex != nil {
				return nil
			}*/
		}

		return nil
	}

	return recurse(v, value, visitedVertices)
}

func (v *vertex[T]) BreadthFirstSearch(writer io.Writer) {
	graphQueue := NewLinkedListQueue()
	visitedVertices := make(map[T]bool)

	visitedVertices[v.value] = true
	graphQueue.enqueue(v)

	for graphQueue.read() != nil {
		// Remove first vertex off queue and make it current vertex.
		currentVertex := graphQueue.dequeue().(*vertex[T])

		// Print the current vertex's value
		fmt.Fprintf(writer, "%v\n", currentVertex.value)

		// Iterate over the current vertex's neighbors.
		for _, neighbor := range currentVertex.neighbors {
			// If it hasn't been visited yet...
			if !visitedVertices[neighbor.value] {
				// Mark as visited.
				visitedVertices[neighbor.value] = true

				// Add neighbor to the queue.
				graphQueue.enqueue(neighbor)
			}
		}
	}
}