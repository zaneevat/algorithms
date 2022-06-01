package algorithms

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[0]

	var less, greater, result []int

	for _, value := range list[1:] {
		if value <= pivot {
			less = append(less, value)
		} else {
			greater = append(greater, value)
		}
	}

	result = append(QuickSort(less), pivot)
	result = append(result, QuickSort(greater)...)

	return result
}
