package algo

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var ints []int

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	ints = make([]int, 1_000_000)

	for i := 0; i < 1_000_000; i++ {
		ints[i] = rand.Int()
	}

	m.Run()
}

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		input []int
		exp   []int
	}{
		{[]int{0, 5, 2, 1, 6, 3}, []int{0, 1, 2, 3, 5, 6}},
		{[]int{0, 5, 2, -1, 6, 3}, []int{-1, 0, 2, 3, 5, 6}},
		{[]int{-5, -99, 0, -100, -999}, []int{-999, -100, -99, -5, 0}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			sortArray := newSortableArray(tc.input)

			sortArray.quicksort(0, len(sortArray.data)-1)

			if diff := cmp.Diff(sortArray.data, tc.exp); diff != "" {
				t.Errorf("exp %v, got %v", tc.exp, sortArray.data)
			}
		})
	}
}

// gotest -bench QuickSort -benchtime=100000x -run=^$
func BenchmarkQuickSort(b *testing.B) {
	slice := ints[0:10]
	sortArray := newSortableArray(slice)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sortArray.partition(0, len(sortArray.data)-1)
	}
}

func BenchmarkQuickSortNative(b *testing.B) {
	slice := ints[0:10]
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	}
}

func TestQuickSelect(t *testing.T) {
	testCases := []struct {
		input []int
		idx   int
		exp   int
	}{
		{[]int{0, 50, 20, 10, 60, 30}, 1, 10},
		{[]int{0, 50, 20, 10, 60, 30}, 3, 30},
		{[]int{0, 5, 2, -1, 6, 3}, 2, 2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			sortArray := newSortableArray(tc.input)

			if result := sortArray.quickselect(tc.idx, 0, len(sortArray.data)-1); result != tc.exp {
				t.Errorf("exp %d, got %d", tc.exp, result)
			}

		})
	}
}

func TestHasDuplicateValue(t *testing.T) {
	testCases := []struct {
		input []int
		exp   bool
	}{
		{[]int{0, 5, 2, -1, 6, 3}, false},
		{[]int{0, 5, 2, -1, 6, 5, 3}, true},
		{[]int{0, 5, 2, -1, 6, 5, 3, 0}, true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v", tc.input), func(t *testing.T) {
			if result := hasDuplicateValue(tc.input); result != tc.exp {
				t.Errorf("exp %t, got %t", result, tc.exp)
			}
		})
	}
}
