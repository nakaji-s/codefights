package main

import "math"

func largestNumber(n int) int {
	if n == 1 {
		return 9
	}
	return 9*int(math.Pow10(n-1)) + largestNumber(n-1)
}
