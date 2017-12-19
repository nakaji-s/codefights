package main

import (
	"fmt"
	"strconv"
)

func deleteDigit(num int) int {
	max := 0
	str := fmt.Sprint(num)

	for i := 0; i < len(str); i++ {
		s := fmt.Sprint(str[:i] + str[i+1:])
		n, _ := strconv.Atoi(s)
		if max < n {
			max = n
		}
	}
	return max
}
