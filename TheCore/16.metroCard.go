package main

func metroCard(lastNumberOfDays int) []int {
	m := map[int][]int{
		28: []int{31},
		30: []int{31},
		31: []int{28, 30, 31},
	}
	n, _ := m[lastNumberOfDays]
	return n
}
