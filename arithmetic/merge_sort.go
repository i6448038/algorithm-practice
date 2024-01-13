package main

import "fmt"

func main() {
	array := []int{3, 15, 3, 4, 5}
	fmt.Println(MergeSort(array))
}

func MergeSort(array []int) []int {
	l := len(array) - 1
	if l < 2 {
		return array
	}

	mid := l / 2

	left := MergeSort(array[:mid])
	right := MergeSort(array[mid:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	ret := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			ret[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			ret[k] = left[i]
			i++
		} else if left[i] < right[j] {
			ret[k] = left[i]
			i++
		} else {
			ret[k] = right[j]
			j++
		}
	}
	return ret
}
