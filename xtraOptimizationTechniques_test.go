package algo

import (
	"fmt"
	"testing"
)

/*
[]int{2, 0, 4, 1, 7, 9}
[]int{2, 0, 4, 5, 3, 9}
*/

func TestTwoSumInefficient(t *testing.T) {
	testCases := map[string]struct {
		input  []int
		target int
		exp    bool
	}{
		"TwoSumTrue":  {[]int{2, 0, 4, 1, 7, 9}, 10, true},
		"TwoSumFalse": {[]int{2, 0, 4, 5, 3, 9}, 10, false},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := twoSumInefficient(tc.input, tc.target); got != tc.exp {
				t.Errorf("exp bool %t, got %t", tc.exp, got)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	testCases := map[string]struct {
		input  []int
		target int
		exp    bool
	}{
		"TwoSumTrue":  {[]int{2, 0, 4, 1, 7, 9}, 10, true},
		"TwoSumFalse": {[]int{2, 0, 4, 5, 3, 9}, 10, false},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := twoSum(tc.input, tc.target); got != tc.exp {
				t.Errorf("exp bool %t, got %t", tc.exp, got)
			}
		})
	}
}

// gotest -bench Anagram -benchtime=100000x -run=^$
func BenchmarkTwoSumInefficient(b *testing.B) {
	input := []int{2, 0, 4, 1, 7, 9}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = twoSumInefficient(input, 100)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	input := []int{2, 0, 4, 1, 7, 9}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = twoSum(input, 100)
	}
}

func TestGameWinnerInefficient(t *testing.T) {
	size := 10
	startingPlayer := "you"

	result := gameWinnerInefficient(size, startingPlayer)

	fmt.Println(result)
}