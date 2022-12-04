package algo

import "golang.org/x/exp/constraints"

/*
Current implementation returns T's zero-value if there is no data in the
heap's array. A production setting would probably benefit with another
way to handle this (bool, err, panic).
Also, as always, concurrency safety would be needed! See stack.go for use of RWMutex,
which could easily be implemented here.
*/

// heap represents an array-backed heap data structure.
type heap[T constraints.Ordered] struct {
	data []T
}

func NewHeap[T constraints.Ordered]() *heap[T] {
	return &heap[T]{
		data: []T{},
	}
}

// RootNode returns the root node's value.
func (h *heap[T]) RootNode() T {
	if len(h.data) < 1 {
		var zeroVal T
		return zeroVal
	}

	return h.data[0]
}

// LastNode returns the value of the last node.
func (h *heap[T]) LastNode() T {
	if len(h.data) < 1 {
		var zeroVal T
		return zeroVal
	}

	return h.data[len(h.data)-1]
}

// Insert will add a new node with the given value to the heap and
// trickle it up to the proper position.
func (h *heap[T]) Insert(value T) {
	// Append new value as last node in heap.
	h.data = append(h.data, value)

	// Track the index of the newly inserted node.
	newNodeIndex := len(h.data) - 1

	// Execute the 'trickle up' algorithm:
	// If new node isn't in root position is greater than its parent node.
	for newNodeIndex > 0 && h.data[newNodeIndex] > h.data[h.parentIndex(newNodeIndex)] {
		// Swap new node with its parent.
		h.data[newNodeIndex], h.data[h.parentIndex(newNodeIndex)] =
			h.data[h.parentIndex(newNodeIndex)], h.data[newNodeIndex]

		// Update index of the new node:
		newNodeIndex = h.parentIndex(newNodeIndex)
	}
}

// Remove deletes the root node from the index and returns its value.
func (h *heap[T]) Remove() T {
	if len(h.data) < 1 {
		var zeroVal T
		return zeroVal
	}

	dequeueNode := h.data[0]

	// Pop the last node and make it the root node.
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	// No need to heapify if length is less than 3.
	if len(h.data) < 3 {
		return dequeueNode
	}

	// Track the index of the trickle node.
	trickleNodeIndex := 0

	// Execute the 'trickle down' algorithm:
	// If the trickle node has a child that is greater than it.
	for h.hasGreaterChild(trickleNodeIndex) {
		// Save the large child index in a var.
		largerChildIndex := h.greaterChildIndex(trickleNodeIndex)

		// Swap trickle node with its larger child.
		h.data[trickleNodeIndex], h.data[largerChildIndex] = h.data[largerChildIndex], h.data[trickleNodeIndex]
	}

	return dequeueNode
}

// heap.leftChildIndex calculates the index of the given node's left child.
func (h *heap[T]) leftChildIndex(index int) int {
	return index*2 + 1
}

// heap.rightChildIndex calculates the index of the given node's right child.
func (h *heap[T]) rightChildIndex(index int) int {
	return index*2 + 2
}

// heap.parentIndex calculates the index of the given node's parent.
func (h *heap[T]) parentIndex(index int) int {
	return (index - 1) / 2
}

// heap.hasGreaterChild determines if the node at the given index
// has a child with a value greater than it.
func (h *heap[T]) hasGreaterChild(index int) bool {
	// Check if node at index has left/right children and
	// if their value is greater than the node at index.
	return h.data[h.leftChildIndex(index)] > h.data[index] || h.data[h.rightChildIndex(index)] > h.data[index]
}

// heap.greaterChildIndex returns the index of the child with a greater
// value than the node at the given index.
func (h *heap[T]) greaterChildIndex(index int) int {
	// If there is no right child.
	// if h.data[h.rightChildIndex(index)] == nil {
	// 	return h.leftChildIndex(index)
	// }

	// If right child value is greater than left child value, return right child index.
	if h.data[h.rightChildIndex(index)] > h.data[h.leftChildIndex(index)] {
		return h.rightChildIndex(index)
	}

	// Otherwise if left child value is greater than or equal to right child,
	// return left child index.
	return h.leftChildIndex(index)
}
