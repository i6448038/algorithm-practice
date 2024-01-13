package main

import "fmt"

func main() {
	testArray := []int{111, 6, 13, 4, 15, 6}
	fmt.Println(QuickSort(testArray))
}

func QuickSort(messArray []int) []int {
	if len(messArray) == 0 {
		return messArray
	}

	var leftElems []int
	var rightElems []int
	flagElem := messArray[0]
	for _, e := range messArray[1:] {
		if flagElem <= e {
			rightElems = append(rightElems, e)
		} else {
			leftElems = append(leftElems, e)
		}
	}
	result := append(QuickSort(leftElems), flagElem)
	result = append(result, QuickSort(rightElems)...)
	return result
}
