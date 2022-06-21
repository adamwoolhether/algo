package algo

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBubbleSortInt(t *testing.T) {
	testCases := []struct {
		input []int
		exp   []int
	}{
		{[]int{65, 55, 45, 35, 25, 15, 10}, []int{10, 15, 25, 35, 45, 55, 65}},
	}

	for _, tc := range testCases {
		input := fmt.Sprintf("%v", tc.input)
		t.Run(input, func(t *testing.T) {
			got := bubbleSort(tc.input)
			if !cmp.Equal(got, tc.exp) {
				t.Errorf("Exp %v, got %v", tc.exp, got)
			}
		})
	}
}

func TestBubbleSortFloat(t *testing.T) {

	testCases := []struct {
		input []float64
		exp   []float64
	}{
		{[]float64{13, 11.7, 13.99, 11.5, 13.9, 11, 13.8}, []float64{11, 11.5, 11.7, 13, 13.8, 13.9, 13.99}},
	}

	for _, tc := range testCases {
		input := fmt.Sprintf("%v", tc.input)
		t.Run(input, func(t *testing.T) {
			got := bubbleSort(tc.input)

			if !cmp.Equal(got, tc.exp) {
				t.Errorf("Exp %v, got %v", tc.exp, got)
			}
		})
	}
}
