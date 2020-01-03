package main

import "strconv"

func c(array []int, m int)[][m]string {
	ret := make([][m]string, 0)
	if m <= 1{
		tmp := [m]string{}
		for _, e := range array{
			tmp[0] = e
		}
		ret = append(ret, strconv.Itoa(tmp))
	}else{

	}
	return ret
}
