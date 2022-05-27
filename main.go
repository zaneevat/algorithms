package main

import (
	"algorithms/algorithms"
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(algorithms.BinarySearch(list, 9))
	fmt.Println(algorithms.BinarySearch(list, -1))

	list1 := []int{8, 5, 4, 3, 2, 0, 7, 6}
	fmt.Println(algorithms.SelectionSort(list1))

	fmt.Println(algorithms.Sum([]int{1, 2, 3, 4}))
	fmt.Println(algorithms.SliceCount([]int{1, 2, 3, 4}))
	fmt.Println(algorithms.Max([]int{1, 2, 3, 4}))
}
