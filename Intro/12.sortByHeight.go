package main

import "sort"

func sortByHeight(a []int) []int {
	b := []int{}
	for _, x := range a {
		if x != -1 {
			b = append(b, x)
		}
	}

	sort.Ints(b)

	j := 0
	for i := 0; i < len(a); i++ {
		if a[i] != -1 {
			a[i] = b[j]
			j++
		}
	}

	return a
}
