package main

import "fmt"

type GoldHill struct {
	PersonNum int
	GoldNum   int
}

func main() {
	list := []GoldHill{
		{3, 11},
		{5, 12},
		{5, 13},
		{10, 20},
	}
	fmt.Println(getMostGold(20, list))
}

func getMostGold(num int, list []GoldHill) int {
	if num <= 0 {
		return 0
	}

	if num >= getGoldHillPersonNum(list) {
		return getGoldHillGoldNum(list)
	}

	reduce := list[1:]
	reduceNum := num - list[0].PersonNum
	goldNum := 0
	if reduceNum < 0 {
		return getMostGold(reduceNum, reduce)
	}
	goldNum = list[0].GoldNum

	return max(getMostGold(reduceNum, reduce)+goldNum, getMostGold(num, reduce))
}

func getGoldHillPersonNum(list []GoldHill) (res int) {
	for _, e := range list {
		res += e.PersonNum
	}
	return
}

func getGoldHillGoldNum(list []GoldHill) (res int) {
	for _, e := range list {
		res += e.GoldNum
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
