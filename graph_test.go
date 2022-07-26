package algo

import (
	"fmt"
	"os"
	"testing"
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
