package algorithms

import (
	"algorithms"
	"testing"
)

func TestSum(t *testing.T) {
	sum := algorithms.Sum([]int{-1, 3, 89, 10})

	if sum != 101 {
		t.Error("Wrong amount")
	}
}

func TestSumEmpty(t *testing.T) {
	sum := algorithms.Sum([]int{})

	if sum != 0 {
		t.Error("Wrong amount")
	}
}

func TestSliceCount(t *testing.T) {
	count := algorithms.SliceCount([]int{1, 2, 3, 4, 5, 6})

	if count != 6 {
		t.Error("Wrong quantity")
	}
}

func TestSliceCountEmpty(t *testing.T) {
	count := algorithms.SliceCount([]int{})

	if count != 0 {
		t.Error("Wrong quantity")
	}
}

func TestMax(t *testing.T) {
	max, err := algorithms.Max([]int{1, 87, -1, 6, 27})

	if max != 87 || err != nil {
		t.Error("Wrong max int")
	}
}

func TestMaxEmpty(t *testing.T) {
	_, err := algorithms.Max([]int{})

	if err == nil {
		t.Error("Wrong max int")
	}
}
