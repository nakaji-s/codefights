package main

import "fmt"

func digitDegree(n int) int {
	count := 0
	for n > 9 {
		n = getSum(n)
		count++
	}
	return count
}

func getSum(n int) int {
	str := fmt.Sprint(n)
	sum := 0
	for _, s := range str {
		num := s - '0'
		sum += int(num)
	}
	return sum
}
