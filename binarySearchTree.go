package algo

import (
	"fmt"
	"io"

	"golang.org/x/exp/constraints"
)

// treeNode represents a node in a Binary Search Tree.
type treeNode[T constraints.Ordered] struct {
	value      T
	leftChild  *treeNode[T]
	rightChild *treeNode[T]
}

// NewTreeNode returns a new Binary Search Tee node.
func NewTreeNode[T constraints.Ordered](value T, left, right *treeNode[T]) *treeNode[T] {
	return &treeNode[T]{
		value:      value,
		leftChild:  left,
		rightChild: right,
	}
}

// Search will find and return the node with the given value.
// O(log N)
func (tn *treeNode[T]) Search(searchValue T) *treeNode[T] {
	// Base case: If node is nil or matches the Search value.
	if tn == nil || tn.value == searchValue {
		return tn
	}

	if searchValue < tn.value {
		return tn.leftChild.Search(searchValue)
	}

	// if searchValue > node.value
	return tn.rightChild.Search(searchValue)
}

// Insert creates a node with the given value at the correct
// location in a Binary Search Tree.
// O(log N) + 1, or O(log N)
func (tn *treeNode[T]) Insert(value T) {
	// If value is already in tree, no further action is needed.
	if tn.value == value {
		return
	}

	if value < tn.value {
		// If left child doesn't exist, Insert the value as left child.
		if tn.leftChild == nil {
			tn.leftChild = &treeNode[T]{value: value}
			// return
		} else {
			tn.leftChild.Insert(value)
		}
	} else { // value > node.value
		if tn.rightChild == nil {
			tn.rightChild = &treeNode[T]{value: value}
		} else {
			tn.rightChild.Insert(value)
		}
	}
}

// Remove deletes the node with the given value from a Binary Search Tree.
// 'Remove' is used, as 'delete' clashes with Go's 'delete' built-in.
// O(log N)
func (tn *treeNode[T]) Remove(valueToDelete T) *treeNode[T] {
	if tn == nil {
		return nil
	}

	switch {
	// If value to be deleted is less or greater than current node, set the left or right child
	// respectively to be the return value of a recursive call of this func on the current node's
	// left or right subtree.
	case valueToDelete < tn.value:
		tn.leftChild = tn.leftChild.Remove(valueToDelete)
		// Return the current node (and its subtree if it exists) to be
		// used as the new value of its parent's left or right child.
		return tn
	case valueToDelete > tn.value:
		tn.rightChild = tn.rightChild.Remove(valueToDelete)
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

// lift finds the successor node, returning either the original right child
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

// FetchMin returns the value of the node in the minimum element position.
func (tn *treeNode[T]) FetchMin() T {
	if tn.leftChild == nil {
		return tn.value
	}

	return tn.leftChild.FetchMin()
}

// FetchMax returns the value of the node in the maximum element position
func (tn *treeNode[T]) FetchMax() T {
	if tn.rightChild == nil {
		return tn.value
	}

	return tn.rightChild.FetchMax()
}

// TraverseAndPrintInOrder traverse's the entire Binary Search Tree in
//  inorder and prints the value to the given writer.
func (tn *treeNode[T]) TraverseAndPrintInOrder(writer io.Writer) {
	if tn == nil {
		return
	}

	tn.leftChild.TraverseAndPrintInOrder(writer)
	_, err := fmt.Fprintln(writer, tn.value)
	if err != nil {
		panic(err)
	}
	tn.rightChild.TraverseAndPrintInOrder(writer)
}

func (tn *treeNode[T]) TraverseAndPrintPreOrder(writer io.Writer) {
	if tn == nil {
		return
	}

	_, err := fmt.Fprintln(writer, tn.value)
	if err != nil {
		panic(err)
	}
	tn.leftChild.TraverseAndPrintPreOrder(writer)
	tn.rightChild.TraverseAndPrintPreOrder(writer)
}

func (tn *treeNode[T]) TraverseAndPrintPostOrder(writer io.Writer) {
	if tn == nil {
		return
	}

	tn.leftChild.TraverseAndPrintPostOrder(writer)
	tn.rightChild.TraverseAndPrintPostOrder(writer)
	_, err := fmt.Fprintln(writer, tn.value)
	if err != nil {
		panic(err)
	}
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Example use case: Maintaining a list of Book titles.
// 1. Can print the list of book title alphabetically.
// 2. Allows for constant changes to the list.
// 3. Allows searching for a title within the list.
/*	root := NewTreeNode("Moby Dick", nil, nil)
	root.Insert("The Odyssey")
	root.Insert("Pride and Prejudice")
	root.Insert("Robinson Crusoe")
	root.Insert("Great Expectations")
	root.Insert("Alice in Wonderland")
	root.Insert("Lord of the Flies")
	root.Insert("Alice in Wonderland")
	root.Insert("Alice in Wonderland")
	root.Insert("Alice in Wonderland")
	root.TraverseAndPrintInOrder(os.Stdout)
*/
