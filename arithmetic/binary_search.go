package main

import (
	"fmt"
)

func main() {
	sortedArray := []int{1, 2, 3, 4, 6, 7}
	position, ret := BinarySearch2(sortedArray, 7, 0, len(sortedArray) - 1)
	position, ret = BinarySearch1(sortedArray, 1)

	fmt.Println(position, ret)
}

//非递归
func BinarySearch1(sortedArray []int, searchKey int) (int, bool) {
	flag := len(sortedArray)/2 -1
	for flag >= 0 && flag <= len(sortedArray) - 1{
		if searchKey == sortedArray[flag] {
			return flag, true
		} else if searchKey < sortedArray[flag] {
			flag = (flag - 1)/2
		} else if searchKey > sortedArray[flag] {
			flag = (len(sortedArray)-1 - (flag + 1)) /2 + flag + 1
		}
	}
	return -1 , false
}

//递归
func BinarySearch2(sortedArray []int, searchKey, begin, end int) (int, bool) {
	flag := (end - begin) /2 + begin
	if searchKey == sortedArray[flag] {
		return flag, true
	} else if searchKey < sortedArray[flag] {
		return BinarySearch2(sortedArray, searchKey, begin, flag-1)
	} else {
		return BinarySearch2(sortedArray, searchKey, flag+1, end)
	}
}