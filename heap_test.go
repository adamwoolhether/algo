package algo

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHeap_Insert(t *testing.T) {

	input := []int{55, 22, 34, 10, 2, 9, 68}
	exp := []int{68, 22, 55, 10, 2, 9, 34}

	testHeap := NewHeap[int]()

	for _, v := range input {
		testHeap.Insert(v)
	}

	if diff := cmp.Diff(testHeap.data, exp); diff != "" {
		t.Errorf("exp %v, got %v; diff: %v", exp, testHeap.data, diff)
	}
}

func TestHeap_Delete(t *testing.T) {
	testCases := []struct {
		input     []int
		runRemove int
		exp       []int
	}{
		{input: []int{55, 22, 34, 10, 2, 9, 68}, runRemove: 1, exp: []int{55, 22, 34, 10, 2, 9}},
		{input: []int{55, 22, 34, 10, 2, 9, 68}, runRemove: 2, exp: []int{34, 22, 9, 10, 2}},
		{input: []int{55, 22, 34, 10, 2, 9, 68}, runRemove: 3, exp: []int{22, 2, 9, 10}},
		{input: []int{55, 22, 34, 10, 2, 9, 68}, runRemove: 4, exp: []int{10, 2, 9}},
		{input: []int{55, 22, 34, 10, 2, 9, 68}, runRemove: 5, exp: []int{9, 2}},
		{input: []int{55, 22, 34, 10, 2, 9, 68}, runRemove: 6, exp: []int{2}},
		{input: []int{55, 22, 34, 10, 2, 9, 68}, runRemove: 7, exp: []int{}},
		// {input: []int{55, 22, 34, 10, 2, 9, 68}, runRemove: 8, exp: []int{}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("remove %d times", tc.runRemove), func(t *testing.T) {
			testHeap := NewHeap[int]()

			for _, v := range tc.input {
				testHeap.Insert(v)
			}

			for i := 0; i < tc.runRemove; i++ {
				poppedVal := testHeap.data[0]
				if val := testHeap.Remove(); val != poppedVal {
					t.Fatalf("exp popped val %d, got %d", poppedVal, val)
				}
			}

			if diff := cmp.Diff(testHeap.data, tc.exp); diff != "" {
				t.Errorf("exp %v, got %v; diff: %v", tc.exp, testHeap.data, diff)
			}
		})
	}
}

func TestHeap_DeleteEmpty(t *testing.T) {
	testHeap := NewHeap[string]()
	if val := testHeap.Remove(); val != "" {
		t.Errorf("Exp empty string, got: %s", val)
	}

	testHeapInt := NewHeap[int]()
	if val := testHeapInt.Remove(); val != 0 {
		t.Errorf("Exp zero val, got: %d", val)
	}
}
