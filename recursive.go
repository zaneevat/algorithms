package algorithms

import "errors"

func Sum(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + Sum(list[1:])
}

func SliceCount(list []int) int {
	switch len(list) {
	case 0:
		return 0
	default:
		return 1 + SliceCount(list[1:])
	}
}

func Max(list []int) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("slice is empty")
	}

	max := list[0]

	if len(list) > 1 {
		newMax, err := Max(list[1:])
		if err == nil && newMax > max {
			max = newMax
		}
	}

	return max, nil
}
