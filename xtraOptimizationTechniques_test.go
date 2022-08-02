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

func TestSumSwapInefficient(t *testing.T) {
	testCases := []struct {
		name   string
		input1 []int
		input2 []int
		exp    [2]int
	}{
		{"test1", []int{5, 3, 3, 7}, []int{4, 1, 1, 6}, [2]int{3, 0}},
		{"test2", []int{1, 2, 3, 4, 5}, []int{6, 7, 8}, [2]int{2, 0}},
		{"test3", []int{10, 15, 20}, []int{5, 30}, [2]int{0, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := sumSwapInefficient(tc.input1, tc.input2); result != tc.exp {
				t.Errorf("exp %v, got %v", tc.exp, result)
			}
		})
	}
}
func TestSumSwap(t *testing.T) {
	testCases := []struct {
		name   string
		input1 []int
		input2 []int
		exp    [2]int
	}{
		{"test1", []int{5, 3, 3, 7}, []int{4, 1, 1, 6}, [2]int{3, 0}},
		{"test2", []int{1, 2, 3, 4, 5}, []int{6, 7, 8}, [2]int{2, 0}},
		{"test3", []int{10, 15, 20}, []int{5, 30}, [2]int{0, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := sumSwap(tc.input1, tc.input2); result != tc.exp {
				t.Errorf("exp %v, got %v", tc.exp, result)
			}
		})
	}
}

func BenchmarkSumSwapInefficientShortArray(b *testing.B) {
	input1 := []int{5, 3, 3, 7}
	input2 := []int{4, 1, 1, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumSwapInefficient(input1, input2)
	}
}

func BenchmarkSumSwapShortArray(b *testing.B) {
	input1 := []int{5, 3, 3, 7}
	input2 := []int{4, 1, 1, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumSwap(input1, input2)
	}
}

func BenchmarkSumSwapInefficientLongArray(b *testing.B) {
	input1 := []int{5, 0, 0, 0, 0, 0, 0, 0, 00, 00, 3, 3, 7}
	input2 := []int{4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumSwapInefficient(input1, input2)
	}
}

func BenchmarkSumSwapLongArray(b *testing.B) {
	input1 := []int{5, 0, 0, 0, 0, 0, 0, 0, 00, 00, 3, 3, 7}
	input2 := []int{4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumSwap(input1, input2)
	}
}

func TestMaxGreedy(t *testing.T) {
	testCases := []struct {
		input []int
		exp   int
	}{
		{[]int{3, 7, 99, 2, 88, 2, 4}, 99},
		{[]int{-9, 5, 2, 9, -88, 0}, 9},
		{[]int{-99, -44, -29, -3}, -3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := maxGreedy(tc.input); result != tc.exp {
				t.Errorf("max() exp %d, got %d", tc.exp, result)
			}
		})
	}
}

func BenchmarkMaxGreedy(b *testing.B) {
	// input := []int{3, 7, 99, 2, 88, 2, 4}
	for i := 0; i < b.N; i++ {
		maxGreedy(ints)
	}
}

func TestMaxSum(t *testing.T) {
	input := []int{3, -4, 4, -3, 5, -9}
	exp := 6

	if result := maxSum(input); result != exp {
		t.Errorf("exp %d, got %d", exp, result)
	}
}

func TestIncreasingTriplets(t *testing.T) {
	testCases := []struct {
		input []float64
		exp   bool
	}{
		{[]float64{5, 2, 8, 4, 3, 7}, true},
		{[]float64{8, 9, 7, 10}, true},
		{[]float64{8, 9, 4, 5}, false},
		{[]float64{22, 25, 21, 18, 19.6, 17, 16, 20.5}, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			if result := increasingTriplets(tc.input); result != tc.exp {
				t.Errorf("exp %t, got %t", tc.exp, result)
			}
		})
	}
}

func BenchmarkAreAnagramsNested(b *testing.B) {
	string1 := "William Shakespeare"
	string2 := "I am a weakish speller"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		areAnagramsNested(string1, string2)
	}
}

func BenchmarkAreAnagramsSorted(b *testing.B) {
	string1 := "William Shakespeare"
	string2 := "I am a weakish speller"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		areAnagramsSorted(string1, string2)
	}
}

func BenchmarkAreAnagrams(b *testing.B) {
	string1 := "William Shakespeare"
	string2 := "I am a weakish speller"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		areAnagrams(string1, string2)
	}

}
