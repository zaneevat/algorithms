package algorithms

import (
	"algorithms"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	sortSlice := algorithms.QuickSort([]int{5, 1, 10, 26, 105, -1})

	if !reflect.DeepEqual([]int{-1, 1, 5, 10, 26, 105}, sortSlice) {
		t.Error("Sorting is incorrect")
	}
}

func TestQuickSortOneElement(t *testing.T) {
	sortSlice := algorithms.QuickSort([]int{5})

	if !reflect.DeepEqual([]int{5}, sortSlice) {
		t.Error("Sorting is incorrect")
	}

	if reflect.DeepEqual([]int{10}, sortSlice) {
		t.Error("Sorting is incorrect")
	}
}

func TestQuickSortSortedSlice(t *testing.T) {
	list := []int{-5, -3, -1, -1, 0, 1, 2, 3}
	sortSlice := algorithms.QuickSort(list)

	if !reflect.DeepEqual(list, sortSlice) {
		t.Error("Sorting is incorrect")
	}

	if reflect.DeepEqual([]int{10}, sortSlice) {
		t.Error("Sorting is incorrect")
	}
}

func TestQuickSortEmpty(t *testing.T) {
	var list []int
	sortSlice := algorithms.QuickSort(list)

	if !reflect.DeepEqual(list, sortSlice) {
		t.Error("Sorting is incorrect")
	}
}
