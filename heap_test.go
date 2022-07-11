package algo

import (
	"fmt"
	"testing"
)

func TestNewHeap(t *testing.T) {
	root := NewHeap(55)

	fmt.Println(root)
}
