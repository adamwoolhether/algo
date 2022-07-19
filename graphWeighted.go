package algo

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type weightedGraphVertex[T constraints.Ordered] struct {
	mu        sync.RWMutex
	value     T
	neighbors map[*weightedGraphVertex[T]]int
}

func NewWeightedVertex[T constraints.Ordered](value T) *weightedGraphVertex[T] {
	return &weightedGraphVertex[T]{
		value:     value,
		neighbors: make(map[*weightedGraphVertex[T]]int),
	}
}

func (wv *weightedGraphVertex[T]) AddNeighbor(vertex *weightedGraphVertex[T], weight int) {
	wv.mu.Lock()
	defer wv.mu.Unlock()

	wv.neighbors[vertex] = weight
}

// /////////////////////////////////////////////////////////////////
// Dijkstra's Algorithm Example

type City struct {
	name   string
	routes map[*City]int
}

func NewCity(name string) *City {
	return &City{
		name:   name,
		routes: make(map[*City]int),
	}
}

func (c *City) AddRoute(city *City, price int) {
	c.routes[city] = price
}
