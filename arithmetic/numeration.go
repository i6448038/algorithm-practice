package main

import (
	"strings"
)

func trans2Binary(num int) string {
	ret := make([]string, 0)
	for num > 0 {
		if num%2 == 1 {
			ret = append(ret, "1")
		} else {
			ret = append(ret, "0")
		}
		num = num / 2
	}

	tmp := make([]string, 0)
	for i := len(ret) - 1; i >= 0; i-- {
		tmp = append(tmp, ret[i])
	}

	return strings.Join(tmp, "")
}
