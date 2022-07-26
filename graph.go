package algo

import (
	"fmt"
	"io"
	"sync"

	"golang.org/x/exp/constraints"
)

/*
Implement a factory pattern to return optional directional/non-directional or weighted/non-weighted graph.
*/

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

// AddNeighborUndirected adds a given vertex to the calling vertex's
// list of neighbors and vice-versa. For use in an undirected graph.
func (v *vertex[T]) AddNeighborUndirected(vertex *vertex[T]) {
	// Prevent an infinite loop.
	for _, n := range v.neighbors {
		if n == vertex {
			return
		}
	}

	v.neighbors = append(v.neighbors, vertex)
	v.AddNeighborUndirected(v)
}

// DepthFirstTraversal Uses the Depth First Search algorith to traverse over all
// neighbors of a vertex, printout their values out to the given io.Writer.
func (v *vertex[T]) DepthFirstTraversal(writer io.Writer) {
	visitedVertices := make(map[T]bool)

	var recurse func(vert *vertex[T], m map[T]bool)
	recurse = func(vert *vertex[T], m map[T]bool) {
		// Mark the vertex as visited.
		m[vert.value] = true

		// Print the vertex's value.
		fmt.Fprintf(writer, "%v", vert.value)

		// Iterate through the current vertex's neighbors.
		for _, neighbor := range vert.neighbors {
			// Ignore the neighbor if it's been visited already.
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

// BreadthFirstTraversal uses a Breadth First algorithm to traverse
// over all neighbors of the calling vertex, printing out their
// values to the given io.Writer. Neighbors to be visited are
// stored in a queue.
func (v *vertex[T]) BreadthFirstTraversal(writer io.Writer) {
	graphQueue := NewLinkedListQueue()
	visitedVertices := make(map[T]bool)

	visitedVertices[v.value] = true
	graphQueue.enqueue(v)

	for graphQueue.read() != nil {
		// Remove first vertex off queue and make it current vertex.
		currentVertex := graphQueue.dequeue().(*vertex[T])

		// Print the current vertex's value
		fmt.Fprintf(writer, "%v", currentVertex.value)

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

func (v *vertex[T]) BreadthFirstSearch(value T) *vertex[T] {
	graphQueue := NewLinkedListQueue()
	visitedVertices := make(map[T]bool)

	visitedVertices[v.value] = true
	graphQueue.enqueue(v)

	for graphQueue.read() != nil {
		currentVertex := graphQueue.dequeue().(*vertex[T])

		if currentVertex.value == value {
			return currentVertex
		}

		for _, neighbor := range currentVertex.neighbors {
			if !visitedVertices[neighbor.value] {
				visitedVertices[neighbor.value] = true
				graphQueue.enqueue(neighbor)
			}
		}
	}

	return nil
}
