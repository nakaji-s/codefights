package main

func countBlackCells(n int, m int) int {
	return n + m + GCDEuclidean(n, m) - 2
}

func GCDEuclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
