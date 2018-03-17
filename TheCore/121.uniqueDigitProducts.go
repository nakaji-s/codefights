package main

import "fmt"

func uniqueDigitProducts(a []int) int {
	m := map[int]struct{}{}
	for _, num := range a {
		m[getProduct(num)] = struct{}{}
	}

	return len(m)
}

func getProduct(num int) int {
	ret := 1

	tmp := fmt.Sprint(num)
	for _, rune := range tmp {
		ret *= int(rune) - '0'
	}

	return ret
}
