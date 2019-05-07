package main

import "fmt"

func main()  {
	fmt.Println(getFibonacci(9))
}

//1、1、2、3、5、8、13、21、34、……
func getFibonacci(index int)int {
	if index <= 0 {
		return -1
	}

	if index == 1 {
		return 1
	} else if index == 2{
		return 1
	} else {
		return getFibonacci(index - 1) + getFibonacci(index - 2)
	}
}
