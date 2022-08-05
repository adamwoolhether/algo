package algo

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestGroupArray(t *testing.T) {
	input := []string{"a", "c", "d", "b", "b", "c", "a", "d", "c", "b", "a", "d"}

	result := groupArray(input)

	if len(result) != len(input) {
		t.Errorf("exp same length. input len=%d, got=%d", len(input), len(result))
	}

	for i := 0; i < len(result)-2; i += 3 {
		if (result[i] != result[i+1]) || (result[i] != result[i+2]) {
			t.Errorf("chars aren't grouped together, result: %v", result)
		}
	}
}

var groupArrayTests = []struct {
	name  string
	input []string
}{
	{"short", []string{"a", "c", "d", "b", "b", "c", "a", "d", "c", "b", "a", "d"}},
	{"long", []string{"a", "c", "d", "b", "b", "c", "a", "d", "c", "b", "a", "d", "a", "c", "d", "b", "b", "c", "a", "d", "c", "b", "a", "d", "a", "c", "d", "b", "b", "c", "a", "d", "c", "b", "a", "d"}},
}

func BenchmarkGroupArray(b *testing.B) {
	for _, tt := range groupArrayTests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				groupArray(tt.input)
			}
		})
	}
}

func BenchmarkGroupArraySort(b *testing.B) {
	for _, tt := range groupArrayTests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				groupArraySort(tt.input)
			}
		})
	}
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exercises

// Exercise 1

func TestFindMultisportAthletes(t *testing.T) {
	bastketballPlayers, footballPlayers := generateAthletes()

	exp := []string{"Jill Huang", "Wanda Vakulskas"}

	result := findMultisportAthletes(bastketballPlayers, footballPlayers)

	if !cmp.Equal(result, exp) {
		t.Errorf("Exp slice %v, got %v", exp, result)
	}
}

func BenchmarkFindMultisportAthletes(b *testing.B) {
	bastketballPlayers, footballPlayers := generateAthletes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findMultisportAthletes(bastketballPlayers, footballPlayers)
	}
}

func TestFindMultisportAthletesInefficient(t *testing.T) {
	bastketballPlayers, footballPlayers := generateAthletes()

	exp := []string{"Jill Huang", "Wanda Vakulskas"}

	result := findMultisportAthletesInefficient(bastketballPlayers, footballPlayers)

	if !cmp.Equal(result, exp) {
		t.Errorf("Exp slice %v, got %v", exp, result)
	}
}

func BenchmarkFindMultisportAthletesInefficient(b *testing.B) {
	bastketballPlayers, footballPlayers := generateAthletes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findMultisportAthletesInefficient(bastketballPlayers, footballPlayers)
	}
}

func generateAthletes() ([]Athletes, []Athletes) {
	basketBallPlayers := []Athletes{
		{FirstName: "Jill", LastName: "Huang", Team: "Gators"},
		{FirstName: "Janko", LastName: "Barton", Team: "Sharks"},
		{FirstName: "Wanda", LastName: "Vakulskas", Team: "Gators"},
		{FirstName: "Jill", LastName: "Moloney", Team: "Gators"},
		{FirstName: "Luuk", LastName: "Watkins", Team: "Gators"},
	}

	footballPlayers := []Athletes{
		{FirstName: "Hanzla", LastName: "Radosti", Team: "32ers"},
		{FirstName: "Tina", LastName: "Watkins", Team: "Barleycorns"},
		{FirstName: "Alex", LastName: "Patel", Team: "32ers"},
		{FirstName: "Jill", LastName: "Huang", Team: "Barleycorns"},
		{FirstName: "Wanda", LastName: "Vakulskas", Team: "Barleycorns"},
	}

	return basketBallPlayers, footballPlayers
}

// 2

func TestFindMissingNum(t *testing.T) {
	tests := []struct {
		input []int
		exp   int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 0},
		{[]int{0, 2, 3, 4, 5, 6}, 1},
		{[]int{0, 1, 3, 4, 5, 6}, 2},
		{[]int{0, 1, 2, 4, 5, 6}, 3},
		{[]int{0, 1, 2, 3, 5, 6}, 4},
		{[]int{0, 1, 2, 3, 4, 6}, 5},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.exp), func(t *testing.T) {
			if res := findMissingNum(tt.input); res != tt.exp {
				t.Errorf("exp %d, got %d", tt.exp, res)
			}
		})
	}
}

func BenchmarkFindMissingNum(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6}
	// 4, 6, 8, 2, 3, 1, 5, 9, 7, 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findMissingNum(input)
	}
}
func BenchmarkFindMissingNumber(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6}
	// 4, 6, 8, 2, 3, 1, 5, 9, 7, 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findMissingNumber(input)
	}
}

func TestFindGreatestProfit(t *testing.T) {
	prices := []int{10, 7, 5, 8, 11, 2, 6}
	exp := 6

	if res := findGreatestProfit(prices); res != exp {
		t.Errorf("Exp %d, got %d", exp, res)
	}
}

func TestGreatestProduct(t *testing.T) {
	input := []int{5, -10, -6, 9, 4}
	exp := 60

	if res := greatestProduct(input); res != exp {
		t.Errorf("Exp %d, got %d", exp, res)
	}
}

func TestSortTemperatures(t *testing.T) {
	input := []float64{98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1}
	exp := []float64{97.1, 97.1, 97.8, 98, 98, 98.2, 98.5, 98.6, 98.9, 99}

	result := sortTemperatures(input)

	if diff := cmp.Diff(exp, result); diff != "" {
		t.Errorf("Unexpected result, diff: %v", diff)
	}

	fmt.Println(result)
}
