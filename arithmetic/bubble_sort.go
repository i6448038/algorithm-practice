package main

import "fmt"

func main() {
	array := []int{3, 13, 2, 11, 4, 6, 13}
	fmt.Println(BubbleSort(array))
}

func BubbleSort(array []int)[]int {
	l := len(array)
	for i := 0; i < l; i++ {
		for j := 0; j < l - i -1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
