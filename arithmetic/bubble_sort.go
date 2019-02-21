package main

import "fmt"

func main() {
	array := []int{3, 13, 2, 11, 4, 6, 13}
	fmt.Println(BubbleSort(array))
}

func BubbleSort(array []int)[]int {
	l := len(array)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
