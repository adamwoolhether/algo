package algo

import "testing"

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		max([]int{1, 2, 3, 4, 5, 6, 7})
	}
}

func BenchmarkMaxInefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxInefficient([]int{1, 2, 3, 4, 5, 6, 7})
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
