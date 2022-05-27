package algorithms

import (
	"algorithms"
	"testing"
)

func TestOk(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	v, err := algorithms.BinarySearch(list, 9)

	if v != 8 || err != nil {
		t.Error("Expected 9, got ", v)
	}
}

func TestEmptySlice(t *testing.T) {
	v, err := algorithms.BinarySearch([]int{}, 9)

	if err == nil {
		t.Error("Search is incorrect", v)
	}
}

func TestNotFound(t *testing.T) {
	v, err := algorithms.BinarySearch([]int{102, 84, 93, 12, 56}, -1)

	if err == nil {
		t.Error("Search is incorrect", v)
	}
}
