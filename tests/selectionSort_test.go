package algorithms

import (
	"algorithms"
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	list := algorithms.SelectionSort([]int{8, 9, 10, 1, 6, 26, 13, 88})

	if !reflect.DeepEqual([]int{1, 6, 8, 9, 10, 13, 26, 88}, list) {
		t.Error("Sorting is incorrect")
	}
}

func TestIsEmpty(t *testing.T) {
	list := algorithms.SelectionSort([]int{})
	if len(list) != len([]int{}) {
		t.Error("Sorting is incorrect")
	}
}

func TestOneElement(t *testing.T) {
	list := algorithms.SelectionSort([]int{9})

	if !reflect.DeepEqual([]int{9}, list) {
		t.Error("Sorting is incorrect")
	}
}
