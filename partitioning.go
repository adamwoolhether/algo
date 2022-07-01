package algo

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type SortableArray[T constraints.Ordered] struct {
	data []T
}

func newSortableArray[T constraints.Ordered](input []T) *SortableArray[T] {
	return &SortableArray[T]{
		data: input,
	}
}

// partition puts the rightmost element in the correct position, and returns the pivot index.
func (s *SortableArray[T]) partition(leftPtr, rightPtr int) int {
	// Choose right-most element as pivot.
	pivotIndex := rightPtr

	// Get the pivot value
	pivot := s.data[pivotIndex]

	// Start right pointer immediately to the left of the pivot.
	rightPtr--

	for true {
		// Move left pointer to right as long as it's value < pivot.
		for s.data[leftPtr] < pivot {
			leftPtr++
		}

		// Move right pointer to left as long as it's value > pivot.
		// Ensure right pointer doesn't get out of range(for negative numbers)
		for s.data[rightPtr] > pivot && rightPtr > 0 {
			rightPtr--
		}

		// Pointers have stopped moving. Check if left pointer has reached.
		// or gon beyond right pointer. If so, break out of loop.
		if leftPtr >= rightPtr {
			break
		}

		// If left pointer is still to left of right pointer, swap their values.
		s.data[leftPtr], s.data[rightPtr] = s.data[rightPtr], s.data[leftPtr]

		// Move left pointer right, preparing for not round of loop.
		leftPtr++
	}

	// Final partitioning step: swap value of left pointer with the pivot.
	s.data[leftPtr], s.data[pivotIndex] = s.data[pivotIndex], s.data[leftPtr]

	// Return the left pointer for the quicksort method.
	return leftPtr
}

// quicksort runs the quicksort algorithm on the SortableArray's data.
func (s *SortableArray[T]) quicksort(leftIndex, rightIndex int) {
	// 	Base case: subarray has 0 or 1 elements.
	if rightIndex-leftIndex <= 0 {
		return
	}

	// Partition range of elements and grab the index of the pivot.
	pivotIndex := s.partition(leftIndex, rightIndex)

	// Recursively call quicksort method on all elements left of the pivot.
	s.quicksort(leftIndex, pivotIndex-1)

	// Recursively call quicksort method on all elements right of the pivot.
	s.quicksort(pivotIndex+1, rightIndex)
}

func (s *SortableArray[T]) quickselect(nthLowestVal, leftIndex, rightIndex int) T {
	// Base case: subarray has one cell.
	if rightIndex-leftIndex <= 0 {
		return s.data[leftIndex]
	}

	// Partition the array and get pivot index.
	pivotIndex := s.partition(leftIndex, rightIndex)

	// If our desired target is to the left of the pivot.
	if nthLowestVal < pivotIndex {
		// Recursively quickselect on subarray to left of pivot.
		return s.quickselect(nthLowestVal, leftIndex, pivotIndex-1)
	}
	// If our desired target is to the right of the pivot.
	if nthLowestVal > pivotIndex {
		// Recursively quickselect on subarray to right of pivot.
		return s.quickselect(nthLowestVal, pivotIndex+1, rightIndex)
	}

	// nthLowestVal == pivotIndex
	// If desired target is the same as pivot index after partition, we've found the target.
	return s.data[pivotIndex]
}

func hasDuplicateValue[T constraints.Ordered](input []T) bool {
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}

	return false
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exercises

// greatestProductOf3 returns the greatest product of any three numbers in an array.
func greatestProductOf3(input []int) int {
	l := len(input)
	if l < 3 {
		return -1
	}
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })

	return input[l-1] * input[l-2] * input[l-3]
}

// findMissingNumber determines if a number is missing in a slice that beings with 0
func findMissingNumber(input []int) int {
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })

	for i := range input {
		if input[i] != i {
			return i
		}
	}

	return -1
}

// O(N^2) implementation
func findGreatest1(input []int) int {
	greatest := input[0]

	for i := 1; i < len(input); i++ {
		isGreatest := true
		greatest = input[i]

		for j := 2; j < len(input); j++ {
			if greatest < input[j] {
				isGreatest = false
			}
		}

		if isGreatest {
			break
		}
	}

	return greatest
}

// ON(N log N) implementation
func findGreatest2(input []int) int {
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })

	return input[len(input)-1]
}

// O(N) implementation.
func findGreatest3(input []int) int {
	greatest := input[0]

	for i := 1; i < len(input); i++ {
		if input[i] > greatest {
			greatest = input[i]
		}
	}

	return greatest
}
