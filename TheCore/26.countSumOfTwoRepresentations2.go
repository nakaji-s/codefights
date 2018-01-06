package main

func countSumOfTwoRepresentations2(n int, l int, r int) int {
	if n/2 < l || n/2 > r {
		return 0
	}

	minDiff := n/2 - l
	if r-n/2 < minDiff {
		minDiff = r - n/2
	}
	return minDiff + 1 - n%2
}
