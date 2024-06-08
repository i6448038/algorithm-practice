package main

import "fmt"

func main() {
	array := []int{5, 3, 4, 1, 6, 2}
	fmt.Println(InsertionSort(array))
}

//插入排序
//5               3            将 3 插入只有一个元素 5 的有序表中
//3 5             4            将 4 插入有两个元素 3 5 的有序表中
//3 4 5           1            将 1 插入有两个元素 3 4 5 的有序表中
//1 3 4 5         6            将 6 插入有两个元素 1 3 4 5 的有序表中
//1 3 4 5 6       2            将 2 插入有两个元素 1 3 4 5 6 的有序表中
//1 2 3 4 5 6
func InsertionSort(messArray []int) []int {
	for i := 1; i < len(messArray); i++ {
		key := messArray[i]
		j := i - 1
		for j >= 0 && messArray[j] > key {
			messArray[j+1] = messArray[j]
			j--
		}
		messArray[j+1] = key

	}
	return messArray
}
