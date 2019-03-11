package main

import "fmt"

func main() {
	fmt.Println(getNum(5, 4))
}


func getNum(m, n int) int {
	if m == 1 && n == 1 {
		return  1
	}else if m == 2 && (n == 1  || n == 2) {
		return 1
	} else if n == m || n == 1 {
		return 1
	} else {
		return getNum(m-1, n-1) + getNum(m-1, n)
	}

}
