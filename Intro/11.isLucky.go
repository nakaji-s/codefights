package main

import (
	"fmt"
	"strconv"
)

func isLucky(n int) bool {
	s := fmt.Sprint(n)

	first := 0
	for i := 0; i < len(s)/2; i++ {
		num, _ := strconv.Atoi(string(s[i]))
		first += num
	}

	latter := 0
	for i := len(s) / 2; i < len(s); i++ {
		num, _ := strconv.Atoi(string(s[i]))
		latter += num
	}

	return first == latter
}
