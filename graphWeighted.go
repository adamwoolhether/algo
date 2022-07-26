package algo

import (
	"sync"

	"golang.org/x/exp/constraints"
)

// weightedGraphVertex represents a node in a weighted graph.
type weightedGraphVertex[T constraints.Ordered] struct {
	mu        sync.RWMutex
	value     T
	neighbors map[*weightedGraphVertex[T]]int
}

// NewWeightedVertex returns a new vertex in for use in a weighted graph.
func NewWeightedVertex[T constraints.Ordered](value T) *weightedGraphVertex[T] {
	return &weightedGraphVertex[T]{
		value:     value,
		neighbors: make(map[*weightedGraphVertex[T]]int),
	}
}

// AddNeighbor adds a new vertex to the graph with the given weight.
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

// DijkstraShortestPath will find the shortest/cheapest path
// from an origin to destination vertex in  a weighted graph
func DijkstraShortestPath(origin, destination *City) []string {
	cheapestPrices := make(map[string]int)
	cheapestPreviousStopover := make(map[string]string)

	// Track known cities that haven't been visited yet.
	// TODO: A better implementation would use a priority queue/heap.
	unvisitedCities := []*City{}

	// Keep track of visited cities with a hash table.
	visitedCities := make(map[string]bool)

	// Add starting city's name as first key in cheapestPrices.
	cheapestPrices[origin.name] = 0

	currentCity := origin

	// Core of the algorithm starts here:
	for currentCity != nil {
		// Add current city's name to visitedCities and
		// remove it from list of unvisited cities.
		visitedCities[currentCity.name] = true
		unvisitedCities = removeCurrentCity(unvisitedCities, currentCity.name)

		// Iterate over the current city's adjacent cities.
		for neighbor, price := range currentCity.routes {
			// Add newly discovered cities to unvisitedCities.
			if !visitedCities[neighbor.name] {
				unvisitedCities = append(unvisitedCities, neighbor)
			}

			// Calculate cost of getting from origin city to neighbor
			// using the current city as the second-to-last stop.
			priceThroughCurrentCity := cheapestPrices[currentCity.name] + price

			// If price from origin city to neighbor is the cheapest
			// price found thus far.
			if currentCheapestPrice, ok := cheapestPrices[neighbor.name]; !ok ||
				priceThroughCurrentCity < currentCheapestPrice {
				// Update the two tables.
				cheapestPrices[neighbor.name] = priceThroughCurrentCity
				cheapestPreviousStopover[neighbor.name] = currentCity.name
			}
		}
		// Vist the next unvisited city, choosing the one that's
		// cheapest to get to from the origin city.
		currentCity = cheapestCityFromOrigin(unvisitedCities, cheapestPrices)
	}

	// Core algorithm completed. cheapestPrices holds all cheapest prices
	// to get from origin to each city. We'll now calculate the price path
	// to take from origin to destination.
	shortestPath := []string{}

	// Work backwards from the final destination to create the shortest path.
	currentCityName := destination.name

	// Loop until starting city is reached.
	for currentCityName != origin.name {
		// Add each currentCityName encountered to shortestPath slice.
		shortestPath = append(shortestPath, currentCityName)
		// Use cheapestPreviousStopover city to follow each city
		// to its previous stopover.
		currentCityName = cheapestPreviousStopover[currentCityName]
	}

	// Cap it off by adding the starting city to the shortest path.
	shortestPath = append(shortestPath, origin.name)

	// Reverse the output to see path from beginning to end.
	return reverseShortestPath(shortestPath)
}

// removeCurrentCity removes the current city from the list of current cities.
func removeCurrentCity(unvisitedCities []*City, currentCityName string) []*City {
	for i := 0; i < len(unvisitedCities); i++ {
		if unvisitedCities[i].name == currentCityName {
			copy(unvisitedCities[i:], unvisitedCities[i+1:])
			unvisitedCities[len(unvisitedCities)-1] = nil
			unvisitedCities = unvisitedCities[:len(unvisitedCities)-1]
		}
	}

	return unvisitedCities
}

// cheapestCityFromOrigin finds the city with the cheapest flight from the origin.
// This is needed because this implementation doesn't use a heap.
func cheapestCityFromOrigin(unvisitedCities []*City, cheapestPrices map[string]int) *City {
	if len(unvisitedCities) <= 0 {
		return nil
	}

	cheapestCity := unvisitedCities[0]
	lowestPrice := cheapestPrices[unvisitedCities[0].name]

	for i := 1; i < len(unvisitedCities); i++ {
		if cheapestPrices[unvisitedCities[i].name] < lowestPrice {
			cheapestCity = unvisitedCities[i]
			lowestPrice = cheapestPrices[unvisitedCities[i].name]
		}
	}

	return cheapestCity
}

// reverseShortestPath flips shortestPath slice, so it is in the proper order.
func reverseShortestPath[T constraints.Ordered](shortestPath []T) []T {
	for i, j := 0, len(shortestPath)-1; i < j; i, j = i+1, j-1 {
		shortestPath[i], shortestPath[j] = shortestPath[j], shortestPath[i]
	}

	return shortestPath
}
