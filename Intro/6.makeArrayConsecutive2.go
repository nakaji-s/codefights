package main

import "sort"

func makeArrayConsecutive2(statues []int) int {
	sort.Ints(statues)
	sum := 0
	for i := 0; i < len(statues)-1; i++ {
		sum += statues[i+1] - statues[i] - 1
	}
	return sum
}
