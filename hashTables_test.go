package algo

import "testing"

func TestIsSubsetInefficient(t *testing.T) {
	testCases := map[string]struct {
		input  []string
		input2 []string
		exp    bool
	}{
		"isSubset":    {[]string{"a", "b", "c", "d", "e", "f"}, []string{"b", "d", "f"}, true},
		"isNotSubset": {[]string{"a", "b", "c", "d", "e", "f"}, []string{"b", "d", "f", "h"}, false},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := isSubsetInefficient(tc.input, tc.input2); got != tc.exp {
				t.Errorf("Exp %t, got %t", tc.exp, got)
			}
		})
	}
}

func TestIsSubset(t *testing.T) {
	testCases := map[string]struct {
		input  []string
		input2 []string
		exp    bool
	}{
		"isSubset":    {[]string{"a", "b", "c", "d", "e", "f"}, []string{"b", "d", "f"}, true},
		"isNotSubset": {[]string{"a", "b", "c", "d", "e", "f"}, []string{"b", "d", "f", "h"}, false},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := isSubset(tc.input, tc.input2); got != tc.exp {
				t.Errorf("Exp %t, got %t", tc.exp, got)
			}
		})
	}
}

func BenchmarkIsSubsetInefficient(b *testing.B) {
	arr1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	arr2 := []string{"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for i := 0; i < b.N; i++ {
		isSubsetInefficient(arr1, arr2)
	}
}

func BenchmarkIsSubset(b *testing.B) {
	arr1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	arr2 := []string{"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for i := 0; i < b.N; i++ {
		isSubset(arr1, arr2)
	}
}
