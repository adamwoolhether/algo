package algo

import "fmt"

// trieNode represents a node in a trie tree. Uint8 is used instead of string for efficiency
type trieNode struct {
	children map[uint8]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[uint8]*trieNode),
	}
}

// trie represents a trie data structure.
type trie struct {
	root *trieNode
}

// NewTrie instantiates a new trie.
func NewTrie() *trie {
	return &trie{
		root: newTrieNode(),
	}
}

// Search searches for a given word in the trie.
func (t *trie) Search(word string) *trieNode {
	currentNode := t.root

	for i := 0; i < len(word); i++ {
		char := word[i]
		// Current node has a child key with current char.
		if n, ok := currentNode.children[char]; ok {
			currentNode = n
		} else {
			// If current char isn't found amount the node's children.
			return nil
		}
	}

	return currentNode
}

// Insert adds a new word to the trie.
func (t *trie) Insert(word string) {
	currentNode := t.root

	for i := 0; i < len(word); i++ {
		char := word[i]
		// Current node has a child key with current char.
		if n, ok := currentNode.children[char]; ok {
			currentNode = n
		} else {
			// If current char isn't found amount the node's children,
			// add it as a new child node.
			newNode := newTrieNode()
			currentNode.children[char] = newNode

			// Follow this node.
			currentNode = newNode
		}
	}

	// After inserting the entire word, add a '*' key to the end.
	currentNode.children['*'] = nil
}

// collectAllWords returns a list of all the trie's words, starting with the calling node.
func (t *trie) collectAllWords(node *trieNode) []string {
	if node == nil {
		node = t.root
	}
	var words []string
	var word []uint8

	var recurse func(*trieNode, []uint8) []string
	recurse = func(currentNode *trieNode, w []uint8) []string {
		// Iterate through all the current node's children.
		for key, childNode := range currentNode.children {
			// A key of '*' means we've found a complete word,
			// add it to the array.
			if key == '*' {
				words = append(words, string(w))
			} else {
				// Still in the middle of a word, so recursively call this function on child node.
				recurse(childNode, append(w, key))
			}
		}

		return words
	}

	return recurse(node, word)
}

// AutoComplete searches the trie for the existence of the given prefix,
// returning the node in the trie that represents the final char of the prefix.
func (t *trie) AutoComplete(prefix string) []string {
	currentNode := t.Search(prefix)

	if currentNode == nil {
		return nil
	}

	return t.collectAllWords(currentNode)
}

// Traverse will iterate through and print all children keys.
// Some additional formatting would be ideal.
func (t *trie) Traverse(node *trieNode) {
	if node == nil {
		node = t.root
	}

	var recurse func(*trieNode)
	recurse = func(n *trieNode) {
		for key, childNode := range n.children {
			fmt.Print(string(key))
			if key != '*' {
				recurse(childNode)
			}
		}
	}

	recurse(node)
}

// Autocorrect will check a given word against words in the trie and
// return a spelling suggestion if it isn't found.
func (t *trie) Autocorrect(word string) string {
	currentNode := t.root

	// Track how much of the given word is found in the trie
	// so far to be concatenated with the best suffix found.
	var wordMatch []uint8

	for i := 0; i < len(word); i++ {
		char := word[i]
		// If the current node has a child key with the current char.
		if n, found := currentNode.children[char]; found {
			wordMatch = append(wordMatch, char)
			// Follow the child node.
			currentNode = n
		} else {
			// If char isn't found in current node's children, collect all
			// suffixes that descend from the current node and get the first one.
			// Concatenate it with the previously found prefix for suggestion.
			wordMatch = append(wordMatch, t.collectAllWords(currentNode)[0][0])

			return string(wordMatch)
		}
	}

	// If complete word is found, return it.
	return word
}
