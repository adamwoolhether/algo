package algo

import "testing"

// gotest -bench Anagram -benchtime=100000x -run=^$
var anagramBenchString = "abcdef"

func BenchmarkAnagramOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anagramsOf(anagramBenchString)
	}
}

func BenchmarkAnagramOfStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anagramsOfStringsBuilder(anagramBenchString)
	}
}
