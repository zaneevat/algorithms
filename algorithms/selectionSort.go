package algorithms

func SelectionSort(list []int) []int {
	var newList []int

	for range list {
		smallest := findSmallest(list)
		newList = append(newList, list[smallest])
		list = append(list[:smallest], list[smallest+1:]...)
	}

	return newList
}

func findSmallest(list []int) int {
	index := 0
	value := 0

	for i, v := range list {
		if i == 0 || v < value {
			index = i
			value = v
		}
	}

	return index
}
