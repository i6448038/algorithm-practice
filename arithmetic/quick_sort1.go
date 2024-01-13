package main

import (
	"fmt"
)

func main() {
	testArray := []int{111, 6, 13, 4, 7, 9, 15, 6}
	fmt.Println(QuickSort1(testArray))
}

func QuickSort1(list []int) []int {
	quick(0, len(list)-1, list)
	return list
}

func quick(l, r int, array []int) {
	if l == r || r < 0 || l >= len(array) {
		return
	}
	flag := array[l]
	i := l
	j := r
	for i < j {
		for array[j] >= flag && i < j {
			j--
		}

		for array[i] <= flag && i < j {
			i++
		}

		array[i], array[j] = array[j], array[i]
	}
	array[l], array[i] = array[i], array[l]
	quick(l, i, array)
	quick(i+1, r, array)
}
