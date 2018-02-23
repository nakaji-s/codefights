package main

import "fmt"

func comfortableNumbers(l int, r int) int {
	m := map[int]int{}

	for i := l; i <= r; i++ {
		sxa := calcSx(i)
		for j := i + 1; j <= r && j <= i+sxa; j++ {
			sxb := calcSx(j)
			if j-sxb <= i {
				m[i]++
			}
		}
	}

	ret := 0
	for _, num := range m {
		ret += num
	}
	return ret
}

func calcSx(x int) int {
	sx := 0
	tmp := fmt.Sprint(x)
	for _, num := range tmp {
		sx += int(num - '0')
	}

	return sx
}
