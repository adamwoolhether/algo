package algo

import (
	"strings"
)

func reverse(s string) string {
	if len(s) == 0 {
		return ""
	}

	return reverse(s[1:]) + string(s[0])
}

func reverse2(str string) string { // do some test with runes instead

	end := len(str) - 1

	var recurse func(s string, idx int) string
	recurse = func(s string, idx int) string {
		if idx < 0 {
			return ""
		}
		return string(s[idx]) + recurse(s, idx-1)
	}

	return recurse(str, end)
}

func countX(str string) int {
	if len(str) == 0 {
		return 0
	}

	if str[0] == 'x' {
		return 1 + countX(str[1:])
	}

	return countX(str[1:])
}

func numberOfPaths(n int) int {
	switch {
	case n < 0: // base case
		return 0
	case n == 0 || n == 1: // also base case
		return 1
	default:
		return numberOfPaths(n-1) + numberOfPaths(n-2) + numberOfPaths(n-3)
	}
}

/*func numberOfPathsV1(n int) int {
	switch {
	case n <= 0:
		return 0
	case n == 1:
		return 1
	case n == 2:
		return 2
	case n == 3:
		return 4
	default:
		return numberOfPathsV1(n-1) + numberOfPathsV1(n-2) + numberOfPathsV1(n-3)
	}
}*/

func anagramsOf(str string) []string {
	if len(str) == 1 {
		return []string{string(str[0])}
	}

	collection := []string{}

	substrAnagrams := anagramsOf(str[1:])

	for _, subStrAnagram := range substrAnagrams {

		for i := 0; i < len(subStrAnagram)+1; i++ {

			cpy := subStrAnagram[:i] + string(str[0]) + subStrAnagram[i:]

			collection = append(collection, cpy)
		}
	}

	return collection
}

func anagramsOfStringsBuilder(str string) []string {
	var sb strings.Builder

	// Base case:
	if len(str) == 1 {
		return []string{string(str[0])}
	}

	var collection []string

	substrAnagrams := anagramsOfStringsBuilder(str[1:])

	for _, subStrAnagram := range substrAnagrams {

		for i := 0; i < len(subStrAnagram)+1; i++ {
			// sb.Grow(len(subStrAnagram[:i] + string(str[0]) + subStrAnagram[i:]))

			sb.WriteString(subStrAnagram[:i])
			sb.WriteString(string(str[0]))
			sb.WriteString(subStrAnagram[i:])

			collection = append(collection, sb.String())
			sb.Reset()
		}
	}

	return collection
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exercises

// totalChars returns the total number of chars in a given slice of strings.
func totalChars(strings []string) int {
	if len(strings) == 1 {
		return len(strings[0])
	}

	return len(strings[0]) + totalChars(strings[1:])
}

// onlyEvens returns only even integers in a slice of int
func onlyEvens(intSlice []int) []int {
	if len(intSlice) == 0 {
		return nil
	}

	if intSlice[0]%2 == 0 {
		return append(intSlice[0:1], onlyEvens(intSlice[1:])...)
	}

	return onlyEvens(intSlice[1:])
}

// triangle adds numbers, incremented by each sequence up to the given integer.
//
// https://en.wikipedia.org/wiki/Triangular_number
func triangle(n int) int {
	if n == 0 {
		return 0
	}

	return n + triangle(n-1)
}

// firstX returns the index of the first 'x' in a given string.
//
// input "abcdefghijklmnopqrstuvwxyz" should return '23'
func firstX(str string) int {
	idx := 0

	var recurse func(s string, i int) int
	recurse = func(s string, i int) int {
		if len(s) == 0 {
			return -1
		}

		if s[0] == 120 {
			return i
		}
		i++

		return recurse(s[1:], i)
	}

	return recurse(str, idx)
}

// firstXrune is same as aboce, but uses runes. see benchmarks for comparison.
func firstXrune(str string) int {
	idx := 0

	var recurse func(s []rune, i int) int
	recurse = func(s []rune, i int) int {
		if len(s) == 0 {
			return -1
		}

		if s[0] == 'x' {
			return i
		}
		i++

		return recurse(s[1:], i)
	}

	return recurse([]rune(str), idx)
}

// firstXrune is same as aboce, but uses runes. see benchmarks for comparison.
func firstXString(str string) int {
	idx := 0

	var recurse func(s string, i int) int
	recurse = func(s string, i int) int {
		if len(s) == 0 {
			return -1
		}

		if string(s[0]) == "x" {
			return i
		}
		i++

		return recurse(s[1:], i)
	}

	return recurse(str, idx)
}

// uniquePaths calculates number of possible shortest paths needed to get from the
// top-left to the bottom-right of a grid with a given amount of columns/rows.
func uniquePaths(rows, columns int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	return uniquePaths(rows-1, columns) + uniquePaths(rows, columns-1)
}
