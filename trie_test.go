package algo

import (
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
