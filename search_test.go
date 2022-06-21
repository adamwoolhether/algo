package algo

import (
	"fmt"
	"testing"
)

func TestLinearSearchInt(t *testing.T) {
	testCases := []struct {
		input  []int
		search int
		exp    int
	}{
		{[]int{3, 17, 75, 80, 202}, 22, -1},
		{[]int{3, 24, 33, 64, 79, 300, 455}, 64, 3},
	}

	for _, tc := range testCases {
		input := fmt.Sprintf("%v", tc.input)
		t.Run(input, func(t *testing.T) {
			if got := linearSearch(tc.input, tc.search); got != tc.exp {
				t.Errorf("Exp %d, got %d", tc.exp, got)
			}

		})
	}
}

func TestLinearSearchFloat(t *testing.T) {
	testCases := []struct {
		input  []float64
		search float64
		exp    int
	}{
		{[]float64{3, 5.6, 78, 99.6, 300.789, 563.4}, 563.4, 5},
		{[]float64{3, 5.6, 78, 99.6, 300.789, 563.4}, 99, -1},
	}

	for _, tc := range testCases {
		input := fmt.Sprintf("%v", tc.input)
		t.Run(input, func(t *testing.T) {
			if got := linearSearch(tc.input, tc.search); got != tc.exp {
				t.Errorf("Exp %d, got %d", tc.exp, got)
			}

		})
	}
}

func TestBinarySearchInt(t *testing.T) {
	testCases := []struct {
		input  []int
		search int
		exp    int
	}{
		{[]int{3, 17, 75, 80, 202}, 22, -1},
		{[]int{3, 24, 33, 64, 79, 300, 455}, 64, 3},
	}

	for _, tc := range testCases {
		input := fmt.Sprintf("%v", tc.input)
		t.Run(input, func(t *testing.T) {
			if got := binarySearch(tc.input, tc.search); got != tc.exp {
				t.Errorf("Exp %d, got %d", tc.exp, got)
			}

		})
	}
}

func TestBinarySearchFloat(t *testing.T) {
	testCases := []struct {
		input  []float64
		search float64
		exp    int
	}{
		{[]float64{3, 5.6, 78, 99.6, 300.789, 563.4}, 563.4, 5},
		{[]float64{3, 5.6, 78, 99.6, 300.789, 563.4}, 99, -1},
	}

	for _, tc := range testCases {
		input := fmt.Sprintf("%v", tc.input)
		t.Run(input, func(t *testing.T) {
			if got := binarySearch(tc.input, tc.search); got != tc.exp {
				t.Errorf("Exp %d, got %d", tc.exp, got)
			}

		})
	}
}
