package main

import (
	"fmt"
	"sort"
)

func digitDifferenceSort(a []int) []int {
	m := map[int][]int{}
	for _, num := range a {
		dNum := differenceNum(num)
		m[dNum] = append([]int{num}, m[dNum]...)
	}

	diffsArray := []int{}
	for key, _ := range m {
		diffsArray = append(diffsArray, key)
	}
	sort.Ints(diffsArray)

	ret := []int{}
	for _, num := range diffsArray {
		ret = append(ret, m[num]...)
	}
	return ret
}

func differenceNum(num int) int {
	min := 10
	max := -1

	for _, rune := range fmt.Sprint(num) {
		tmp := int(rune) - '0'
		if tmp > max {
			max = tmp
		}
		if tmp < min {
			min = tmp
		}
	}
	return max - min
}
