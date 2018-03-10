package main

import "sort"

func shuffledArray(shuffled []int) []int {
	sum := 0
	m := map[int]int{}
	for i, num := range shuffled {
		sum += num
		m[num] = i
	}

	original := sum / 2
	ret := append(shuffled[:m[original]], shuffled[m[original]+1:]...)
	sort.Ints(ret)

	return ret
}
