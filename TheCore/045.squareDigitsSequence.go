package main

import "fmt"

func squareDigitsSequence(a0 int) int {
	m := map[int]struct{}{}
	m[a0] = struct{}{}
	return countSquareDigitsSequence(m, a0, 1)
}

func countSquareDigitsSequence(m map[int]struct{}, a int, count int) int {
	if _, ok := m[a]; ok {
		if count != 1 {
			return count
		}
	}
	m[a] = struct{}{}

	sum := 0
	aStr := fmt.Sprint(a)
	for _, num := range aStr {
		sum += int(num-'0') * int(num-'0')
	}

	return countSquareDigitsSequence(m, sum, count+1)
}
