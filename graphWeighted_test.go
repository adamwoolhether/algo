package algo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDijkstraShortestPath(t *testing.T) {
	exp := []string{"Atlanta", "Denver", "Chicago", "El Paso"}

	atlanta := NewCity("Atlanta")
	boston := NewCity("Boston")
	chicago := NewCity("Chicago")
	denver := NewCity("Denver")
	elPaso := NewCity("El Paso")

	atlanta.AddRoute(boston, 100)
	atlanta.AddRoute(denver, 160)
	boston.AddRoute(chicago, 120)
	boston.AddRoute(denver, 180)
	chicago.AddRoute(elPaso, 80)
	denver.AddRoute(chicago, 40)
	denver.AddRoute(elPaso, 140)

	result := DijkstraShortestPath(atlanta, elPaso)

	if diff := cmp.Diff(exp, result); diff != "" {
		t.Errorf("Exp proper ordering, difference: %v", diff)
	}
}
