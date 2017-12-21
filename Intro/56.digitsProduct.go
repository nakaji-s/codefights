package main

import (
	"fmt"
	"strconv"
)

func digitsProduct(product int) int {
	if product == 0 {
		return 10
	}
	if product == 1 {
		return 1
	}
	remain := product
	nums := ""
	for i := 9; i >= 2; i-- {
		if remain%i == 0 {
			remain = remain / i
			nums = fmt.Sprint(i) + nums
			i++
		}
	}
	if remain > 9 {
		return -1
	}

	num, _ := strconv.Atoi(nums)
	return num
}
