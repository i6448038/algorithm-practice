package main

import "fmt"

// 方法：利用回溯模板
//  res array[]
//  def dfs	(list array[], xxxx) {
//      if 边界条件 {
//			res.push
//			return
// 		}
//		for 循环条件 {
//          list.push(xxx)
//			if 条件语句
//			dfs(list, xxxx)
//			list.pop
//		}
//
// }
func eightQueens() {
	var (
		res = [8][8]int{}
		dfs func(int, [8][8]int)
	)

	dfs = func(row int, array [8][8]int) {
		if row > 7 {
			printRes(array)
			fmt.Println("-------")
			return
		}

		for j := 0; j < 8; j++ {
			array[row][j] = 1
			if !isAttracked(row, j, array) {
				dfs(row+1, array)
			}
			array[row][j] = 0
		}

	}
	dfs(0, res)
	return
}

func isAttracked(row, column int, list [8][8]int) bool {
	for i := 0; i < row; i++ {
		// 上下
		if list[i][column] == 1 {
			return true
		}
		// 左上
		left := column - (row - i)
		if left >= 0 && left <= 7 && list[i][left] == 1 {
			return true
		}
		// 右上
		right := column + (row - i)
		if right >= 0 && right <= 7 && list[i][right] == 1 {
			return true
		}
	}
	return false
}

func printRes(array [8][8]int) {
	for _, elems := range array {
		for _, e := range elems {
			fmt.Print(e)
		}
		fmt.Println()
	}
}
