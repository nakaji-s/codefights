package main

import (
	"fmt"
	"strconv"
)

func mirrorBits(a int) int {
	s := fmt.Sprintf("%b", a)
	rs := Reverse(s)
	n, _ := strconv.ParseInt(rs, 2, 0)
	return int(n)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
