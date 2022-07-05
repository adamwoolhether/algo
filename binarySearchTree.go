package algo

import "golang.org/x/exp/constraints"

/*
Test duplicate values and implement better handling if needed.
*/

// treeNode represents a node in a Binary Search Tree.
type treeNode[T constraints.Ordered] struct {
	value      T
	leftChild  *treeNode[T]
	rightChild *treeNode[T]
}

// newTReeNode returns a new Binary Search Tee node.
func newTreeNode[T constraints.Ordered](value T, left, right *treeNode[T]) *treeNode[T] {
	return &treeNode[T]{
		value:      value,
		leftChild:  left,
		rightChild: right,
	}
}

// search will find and return the node with the given value.
// O(log N)
func (tn *treeNode[T]) search(searchValue T) *treeNode[T] {
	// Base case: If node is nil or matches the search value.
	if tn == nil || tn.value == searchValue {
		return tn
	}

	if searchValue < tn.value {
		return tn.leftChild.search(searchValue)
	}

	// if searchValue > node.value
	return tn.rightChild.search(searchValue)
}

// insert creates a node with the given value at the correct
// location in a Binary Search Tree.
// O(log N) + 1, or O(log N)
func (tn *treeNode[T]) insert(value T) {
	if value < tn.value {
		// If left child doesn't exist, insert the value as left child.
		if tn.leftChild == nil {
			tn.leftChild = &treeNode[T]{value: value}
			// return
		} else {
			tn.leftChild.insert(value)
		}
	} else { // value > node.value
		if tn.rightChild == nil {
			tn.rightChild = &treeNode[T]{value: value}
		} else {
			tn.rightChild.insert(value)
		}
	}
}

// remove deletes the node with the given value from a Binary Search Tree.
// 'remove' is used, as 'delete' clashes with Go's 'delete' built-in.
// O(log N)
func (tn *treeNode[T]) remove(valueToDelete T) *treeNode[T] {
	if tn == nil {
		return nil
	}

	switch {
	// If value to be deleted is less or greater than current node, set the left or right child
	// respectively to be the return value of a recursive call of this func on the current node's
	// left or right subtree.
	case valueToDelete < tn.value:
		tn.leftChild = tn.leftChild.remove(valueToDelete)
		// Return the current node (and its subtree if it exists) to be
		// used as the new value of its parent's left or right child.
		return tn
	case valueToDelete > tn.value:
		tn.rightChild = tn.rightChild.remove(valueToDelete)
		return tn
	default: // valueToDelete == tn.value
		// If the current node has no left child, delete it by returning its right child
		// (and its subtree if it exists) to be its parent's new subtree.
		if tn.leftChild == nil {
			return tn.rightChild
			// If the current node has no left OR right child,
			// this returns nil, per first line of this func.
		} else if tn.rightChild == nil {
			return tn.leftChild
		} else {
			// If the current node has two children, delete the current node by calling the lift
			// function which changes the current node's value to the value of its successor node.
			tn.rightChild = tn.rightChild.lift(tn)
			return tn
		}
	}
}

// life finds the successor node, returning either the original right child
// passed into it, or nil if the original right child ends up as the successor
// node, which happens if it had no left children of its own.
func (tn *treeNode[T]) lift(nodeToDelete *treeNode[T]) *treeNode[T] {
	// If the current node in this method has a left child, we recursively
	// call this func to continue down the left subtree to find the successor.
	if tn.leftChild != nil {
		tn.leftChild = tn.lift(nodeToDelete)
		return tn
	}

	// If the current node has no left child, that means the current node of
	// this method is the successor node, it's value is taken and made be the
	// new value of the node that is to be deleted.
	nodeToDelete.value = tn.value
	// Return the successor node's right child to be now used
	// as its parent's left child.

	return tn.rightChild
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Example use case: Maintaining a list of Book titles.
// 1. Can print the list of book title alphabetically.
// 2. Allows for constant changes to the list.
// 3. Allows searching for a title within the list.
