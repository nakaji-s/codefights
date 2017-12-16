package main

func electionsWinners(votes []int, k int) int {
	max := 0
	maxNum := 0
	for _, n := range votes {
		if max < n {
			max = n
			maxNum = 1
		} else if max == n {
			maxNum++
		}
	}

	sum := 0
	for _, n := range votes {
		if n+k > max || (n == max && maxNum == 1) {
			sum++
		}
	}
	return sum
}
