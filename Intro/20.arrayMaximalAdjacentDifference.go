package main

import "math"

func arrayMaximalAdjacentDifference(inputArray []int) int {
	maxDiff := 0
	for i := 1; i < len(inputArray); i++ {
		num := int(math.Abs(float64(inputArray[i] - inputArray[i-1])))
		if num > maxDiff {
			maxDiff = num
		}
	}
	return maxDiff
}
