package algo

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDepthFirstTraversal(t *testing.T) {
	alice := NewVertex("Alice")
	bob := NewVertex("Bob")
	candy := NewVertex("Candy")
	derek := NewVertex("Derek")
	elaine := NewVertex("Elaine")
	fred := NewVertex("Fred")
	gina := NewVertex("Gina")
	helen := NewVertex("Helen")
	irena := NewVertex("Irena")

	alice.AddNeighbor(bob)
	alice.AddNeighbor(candy)
	alice.AddNeighbor(derek)
	alice.AddNeighbor(elaine)
	bob.AddNeighbor(fred)
	fred.AddNeighbor(helen)
	derek.AddNeighbor(gina)
	gina.AddNeighbor(irena)

	bob.AddNeighbor(alice)
	candy.AddNeighbor(alice)
	derek.AddNeighbor(alice)
	elaine.AddNeighbor(alice)

	fred.AddNeighbor(bob)
	helen.AddNeighbor(fred)
	gina.AddNeighbor(derek)
	irena.AddNeighbor(gina)

	alice.DepthFirstTraversal(os.Stdout)
}

func TestDepthFirstSearch(t *testing.T) {
	alice := NewVertex("Alice")
	bob := NewVertex("Bob")
	candy := NewVertex("Candy")
	derek := NewVertex("Derek")
	elaine := NewVertex("Elaine")
	fred := NewVertex("Fred")
	gina := NewVertex("Gina")
	helen := NewVertex("Helen")
	irena := NewVertex("Irena")

	alice.AddNeighbor(bob)
	alice.AddNeighbor(candy)
	alice.AddNeighbor(derek)
	alice.AddNeighbor(elaine)
	bob.AddNeighbor(fred)
	fred.AddNeighbor(helen)
	derek.AddNeighbor(gina)
	gina.AddNeighbor(irena)

	bob.AddNeighbor(alice)
	candy.AddNeighbor(alice)
	derek.AddNeighbor(alice)
	elaine.AddNeighbor(alice)

	fred.AddNeighbor(bob)
	helen.AddNeighbor(fred)
	gina.AddNeighbor(derek)
	irena.AddNeighbor(gina)

	found := alice.DepthFirstSearch("FAKE")
	if found != nil {
		t.Errorf("exp <nil>, got %v", found)
	}

	found = alice.DepthFirstSearch("Derek")
	if found.value != "Derek" {
		t.Errorf("")
	}
	fmt.Println(found)
}

func TestBreadthFirstTraversal(t *testing.T) {

	alice := NewVertex("Alice")
	bob := NewVertex("Bob")
	candy := NewVertex("Candy")
	derek := NewVertex("Derek")
	elaine := NewVertex("Elaine")
	fred := NewVertex("Fred")
	gina := NewVertex("Gina")
	helen := NewVertex("Helen")
	irena := NewVertex("Irena")

	alice.AddNeighbor(bob)
	alice.AddNeighbor(candy)
	alice.AddNeighbor(derek)
	alice.AddNeighbor(elaine)
	bob.AddNeighbor(fred)
	fred.AddNeighbor(helen)
	derek.AddNeighbor(gina)
	gina.AddNeighbor(irena)

	bob.AddNeighbor(alice)
	candy.AddNeighbor(alice)
	derek.AddNeighbor(alice)
	elaine.AddNeighbor(alice)

	fred.AddNeighbor(bob)
	helen.AddNeighbor(fred)
	gina.AddNeighbor(derek)
	irena.AddNeighbor(gina)

	alice.BreadthFirstTraversal(os.Stdout)
}

func TestBreadthFirstSearch(t *testing.T) {
	alice := NewVertex("Alice")
	bob := NewVertex("Bob")
	candy := NewVertex("Candy")
	derek := NewVertex("Derek")
	elaine := NewVertex("Elaine")
	fred := NewVertex("Fred")
	gina := NewVertex("Gina")
	helen := NewVertex("Helen")
	irena := NewVertex("Irena")

	alice.AddNeighbor(bob)
	alice.AddNeighbor(candy)
	alice.AddNeighbor(derek)
	alice.AddNeighbor(elaine)
	bob.AddNeighbor(fred)
	fred.AddNeighbor(helen)
	derek.AddNeighbor(gina)
	gina.AddNeighbor(irena)

	bob.AddNeighbor(alice)
	candy.AddNeighbor(alice)
	derek.AddNeighbor(alice)
	elaine.AddNeighbor(alice)

	fred.AddNeighbor(bob)
	helen.AddNeighbor(fred)
	gina.AddNeighbor(derek)
	irena.AddNeighbor(gina)

	found := alice.BreadthFirstSearch("FAKE")
	if found != nil {
		t.Errorf("exp <nil>, got %v", found)
	}

	found = alice.BreadthFirstSearch("Derek")
	if found.value != "Derek" {
		t.Errorf("")
	}
	fmt.Println(found)
}

func TestFindShortestPath(t *testing.T) {
	idris := NewVertex("Idris")

	kamil := NewVertex("Kamil")
	lina := NewVertex("Lina")
	sasha := NewVertex("Sasha")

	talia := NewVertex("Talia")
	ken := NewVertex("Ken")
	marco := NewVertex("Marco")

	idris.AddNeighbor(kamil)
	kamil.AddNeighbor(lina)
	lina.AddNeighbor(sasha)
	idris.AddNeighbor(talia)
	talia.AddNeighbor(ken)
	ken.AddNeighbor(marco)
	marco.AddNeighbor(sasha)

	exp := []string{"Idris", "Kamil", "Lina"}

	result := FindShortestPath(idris, lina)

	if diff := cmp.Diff(exp, result); diff != "" {
		t.Errorf("Exp proper ordering, difference: %v", diff)
	}

	fmt.Println(result)
}
