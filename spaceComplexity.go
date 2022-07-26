package algo

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// makeUppercaseInefficient will convert a slice of words into uppercase.
// Space Complexity: O(N)
func makeUppercaseInefficient(words []string) []string {
	newSlice := make([]string, len(words))
	for i, word := range words {
		newSlice[i] = strings.ToUpper(word)
	}

	return newSlice
}

// makeUppercase will convert a slice of words into uppercase.
// Space Complexity: O(1)
func makeUppercase(words []string) []string {
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}

	return words
}

// /////////////////////////////////////////////////////////////////
// Exercises

// reverseSlice reverses a given array in place.
// Space Complexity: O(1)
func reverseSlice[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	return arr
}
