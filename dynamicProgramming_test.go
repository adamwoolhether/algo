package algo

/*
- Handle cases for negative numbers
*/
import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
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
			if result := max(tc.input); result != tc.exp {
				t.Errorf("max() exp %d, got %d", tc.exp, result)
			}
		})
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := maxInefficient(tc.input); result != tc.exp {
				t.Errorf("maxInefficient() exp %d, got %d", tc.exp, result)
			}
		})
	}
}

func TestFib(t *testing.T) {
	testCases := []struct {
		input int
		exp   int
	}{
		{15, 610},
		{8, 21},
		// {-3, 2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := fib(tc.input); result != tc.exp {
				t.Errorf("fib() exp %d, got %d", tc.exp, result)
			}
		})
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			memo := make(map[int]int, 2*tc.input-1)
			if result := fibMemoization(tc.input, memo); result != tc.exp {
				t.Errorf("fibMemoization() exp %d, got %d", tc.exp, result)
			}
		})
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := fibBottomUp(tc.input); result != tc.exp {
				t.Errorf("fibBottomUp() exp %d, got %d", tc.exp, result)
			}
		})
	}
}

func TestAddUntil100(t *testing.T) {
	testCases := []struct {
		input []int
		exp   int
	}{
		{[]int{4, 9, 200, 30, 99, 123, 4}, 47},
		// {[]int{1, 5, -2, 5, 100}, 9},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := addUntil100(tc.input); result != tc.exp {
				t.Errorf("addUntil100() exp %d, got %d", tc.exp, result)
			}
		})
	}
}

func TestGolomb(t *testing.T) {
	testCases := []struct {
		input int
		exp   int
	}{
		{7, 4},
		{3, 2},
		// {-3, -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := golomb(tc.input); result != tc.exp {
				t.Errorf("golomb() exp %d, got %d", tc.exp, result)
			}
		})
	}
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		max(ints)
	}
}

func BenchmarkMaxInefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxInefficient(ints)
	}
}

func BenchmarkFibInefficient10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibInefficient(10)
	}
}
func BenchmarkFibMemoization10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memo := make(map[int]int, 10*2-1)
		fibMemoization(10, memo)
	}
}
func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(10)
	}
}
func BenchmarkFibBottomUp10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibBottomUp(10)
	}
}
func BenchmarkFibInefficient20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibInefficient(20)
	}
}
func BenchmarkFibMemoization20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memo := make(map[int]int, 10*2-1)
		fibMemoization(20, memo)
	}
}
func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(20)
	}
}
func BenchmarkFibBottomUp20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibBottomUp(20)
	}
}
func BenchmarkFibInefficient40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibInefficient(40)
	}
}
func BenchmarkFibMemoization40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memo := make(map[int]int, 10*2-1)
		fibMemoization(40, memo)
	}
}
func BenchmarkFib40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(40)
	}
}
func BenchmarkFibBottomUp40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibBottomUp(40)
	}
}
