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

// gotest -bench TwoSum -run=^$
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
	testCases := []struct {
		size int
		exp  string
	}{
		{1, "them"},
		{2, "you"},
		{3, "you"},
		{4, "them"},
		{5, "you"},
		{6, "you"},
		{7, "them"},
		{8, "you"},
		{9, "you"},
		{10, "them"},
	}

	startingPlayer := "you"
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.size), func(t *testing.T) {
			if result := gameWinnerInefficient(tc.size, startingPlayer); result != tc.exp {
				t.Errorf("exp %s, got %s", tc.exp, result)
			}
		})
	}
}

func TestGameWinner(t *testing.T) {
	testCases := []struct {
		size int
		exp  string
	}{
		{1, "them"},
		{2, "you"},
		{3, "you"},
		{4, "them"},
		{5, "you"},
		{6, "you"},
		{7, "them"},
		{8, "you"},
		{9, "you"},
		{10, "them"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.size), func(t *testing.T) {
			if result := gameWinner(tc.size); result != tc.exp {
				t.Errorf("exp %s, got %s", tc.exp, result)
			}
		})
	}
}

// gotest -bench GameWinner -run=^$
func BenchmarkGameWinnerInefficient(b *testing.B) {
	size := 20
	startingPlayer := "you"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gameWinnerInefficient(size, startingPlayer)
	}
}
func BenchmarkGameWinner(b *testing.B) {
	size := 20

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gameWinner(size)
	}
}
