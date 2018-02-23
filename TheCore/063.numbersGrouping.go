package main

func numbersGrouping(a []int) int {
	m := map[int]int{}
	for _, num := range a {
		m[(num-1)/10000]++
	}

	return len(a) + len(m)
}
