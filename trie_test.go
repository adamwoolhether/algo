package algo

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTrie(t *testing.T) {
	newTrie := NewTrie()

	if newTrie.root.children == nil {
		t.Fatalf("exp non-nil children map")
	}
	if len(newTrie.root.children) > 0 {
		t.Fatalf("exp empty initialized map, got %d", len(newTrie.root.children))
	}
}

func TestTrie_Insert(t *testing.T) {
	newTrie := NewTrie()

	newTrie.Insert("cat")
	if len(newTrie.root.children) != 1 {
		t.Fatalf("exp map len %d, got %d", 1, len(newTrie.root.children))
	}

	newTrie.Insert("bat")
	if len(newTrie.root.children) != 2 {
		t.Fatalf("exp map len %d, got %d", 2, len(newTrie.root.children))
	}
}

func TestTrie_AutoComplete(t *testing.T) {
	newTrie := NewTrie()
	newTrie.Insert("cat")
	newTrie.Insert("cab")
	newTrie.Insert("bat")

	exp := []string{"at", "ab"}

	res := newTrie.AutoComplete("c")
	if diff := cmp.Diff(exp, res); diff != "" {
		t.Errorf("exp %v, got %v; diff: %v", exp, res, diff)
	}
}

func TestTrie_Autocorrect(t *testing.T) {
	newTrie := NewTrie()

	newTrie.Insert("ace")
	newTrie.Insert("act")
	newTrie.Insert("bad")
	newTrie.Insert("bake")
	newTrie.Insert("bat")
	newTrie.Insert("batter")
	newTrie.Insert("cat")
	newTrie.Insert("catnip")
	newTrie.Insert("catnap")

	testCases := []struct {
		input string
		exp   string
	}{
		{"catnar", "catnap"},
		{"catnip", "catnip"},
		{"caxasfdij", "cat"},
		{"bakp", "bake"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.input), func(t *testing.T) {
			if res := newTrie.Autocorrect(tc.input); res != tc.exp {
				t.Errorf("exp %s, got %s", tc.exp, res)
			}
		})
	}
}
