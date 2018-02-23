package main

import (
	"fmt"
	"strconv"
)

func arrayPacking(a []int) int {
	s := ""
	for i := len(a) - 1; i >= 0; i-- {
		s += fmt.Sprintf("%08b", a[i])
	}
	ret, _ := strconv.ParseInt(s, 2, 0)
	return int(ret)
}
