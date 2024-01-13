package main

import "fmt"

func main() {
	array := []int{3, 13, 14, 18, 11, 12, 8, 8, 7}

	fmt.Println(SelectSort(array))
}

func SelectSort(array []int) []int {
	l := len(array)
	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
	return array
}
