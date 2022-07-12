package algo

// trieNode represents a node in a trie tree. Uint8 is used instead of string for efficiency
type trieNode struct {
	children map[rune]*trieNode
}

// trie represents a trie data structure.
type trie struct {
	root *trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode),
	}
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

	for _, r := range word {
		// Current node has a child key with current char.
		if n, ok := currentNode.children[r]; ok {
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

	for _, r := range word {
		// Current node has a child key with current char.
		if n, ok := currentNode.children[r]; ok {
			currentNode = n
		} else {
			// If current char isn't found amount the node's children,
			// add it as a new child node.
			newNode := newTrieNode()
			currentNode.children[r] = newNode

			// Follow this node.
			currentNode = newNode
		}
	}

	// After inserting the entire word, add a '*' key to the end.
	currentNode.children['*'] = nil
}

// CollectAllWords returns a list of all the trie's words, starting with the calling node.
func (t *trie) collectAllWords(node *trieNode) []string {
	if node == nil {
		node = t.root
	}
	var words []string
	var word string

	var recurse func(*trieNode, string) []string
	recurse = func(currentNode *trieNode, w string) []string {
		// Iterate through all the current node's children.
		for key, childNode := range currentNode.children {
			// A key of '*' means we've found a complete word,
			// add it to the array.
			if key == '*' {
				words = append(words, w)
			} else {
				// Still in the middle of a word, so recursively call this function on child node.
				recurse(childNode, w+string(key))
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
