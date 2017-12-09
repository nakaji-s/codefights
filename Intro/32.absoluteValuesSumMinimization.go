package main

import "math"

func absoluteValuesSumMinimization(a []int) int {
	min := a[0]
	max := a[len(a)-1]
	minSub := 100000000
	minNum := 0

	for i := min; i <= max; i++ {
		var sum float64 = 0
		for _, n := range a {
			sum += math.Abs(float64(n - i))
		}
		if minSub > int(sum) {
			minSub = int(sum)
			minNum = i
		}
	}
	return minNum
}
