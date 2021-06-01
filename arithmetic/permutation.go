package main

import "fmt"

func main(){
	fmt.Println(permute([]int{1,2,3,4}))
}


func permute(nums []int) [][]int {
	var (
		iterator func(reserve []int)
		res = make([][]int, 0)
	)


	iterator = func(reserve []int) {
		tmp := getDiff(reserve, nums)
		if len(reserve) == len(nums) {
			tmpArray := make([]int, len(reserve))
			copy(tmpArray, reserve)
			res = append(res, tmpArray)
			return
		}


		for _, e := range tmp {
			reserve = append(reserve, e)

			if len(reserve) == len(nums) && hasExist(reserve, res) {
				reserve = reserve[:len(reserve)-1]
				continue
			}

			iterator(reserve)
			reserve = reserve[:len(reserve)-1]
		}
	}
	iterator([]int{})
	return res
}

func hasExist(array []int, res [][]int) bool {
	if len(res) == 0 {
		return false
	}

	if len(array) != len(res[0]) {
		return false
	}

	for _, elems := range res {
		flag := true
		for idx, e := range elems {
			if array[idx] != e {
				flag = false
				break
			}
		}
		if flag == true {
			return true
		}
	}
	return false
}

func getDiff(reserve, array []int)(res []int){
	elemMap := make(map[int]int)
	res = make([]int, 0)

	for _, e := range reserve {
		elemMap[e] = e
	}

	for _, e := range array {
		if _, ok := elemMap[e]; !ok {
			res = append(res, e)
		}
	}

	return
}
