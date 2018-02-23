package main

import "math"

func numberOfClans(divisors []int, k int) int {
	m := map[int]int{}
	for i := 1; i <= k; i++ {
		m[getHashCode(divisors, i)]++
	}

	return len(m)
}

func getHashCode(divisors []int, targetNum int) int {
	ret := 0
	for i, num := range divisors {
		if targetNum%num == 0 {
			ret += int(math.Exp2(float64(i)))
		}
	}
	return ret
}
