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

func reverse2(s string) string { // do some test with runes instead

	end := len(s) - 1

	var recurse func(s string, idx int) string
	recurse = func(s string, idx int) string {
		if idx < 0 {
			return ""
		}
		return string(s[idx]) + recurse(s, idx-1)
	}

	return recurse(s, end)
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
