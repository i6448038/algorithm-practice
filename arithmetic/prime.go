package main

import (
	"fmt"
	"math"
)


func main(){
	ret := GetPrimes(100)
	fmt.Println(ret)
}

func GetPrimes(n int) []int {
	var ret []int
	for i := 1; i <= n; i++ {
		if IsPrime(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func IsPrime(n int) bool{
	temp := math.Floor(math.Sqrt(float64(n)))
	for i := 1; i <= int(temp); i ++ {
		if i != 1 && n % i == 0 {
			return false
		}
	}
	return true
}

