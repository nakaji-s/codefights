package main

import "math"

func rectangleRotation(a int, b int) int {
	m := float64(b) * math.Cos(math.Pi/4)
	n := float64(a) * math.Sin(math.Pi/4)

	count := 0
	for i := -50; i <= 50; i++ {
		for j := -50; j <= 50; j++ {
			if isInnerY(i, j, m) && isInnerY2(i, j, n) {
				count++
			}
		}
	}
	return count
}

// y = x + m
func isInnerY(x int, y int, m float64) bool {
	return float64(y-x) <= m && float64(y-x) >= -m
}

// y = -x + n
func isInnerY2(x int, y int, n float64) bool {
	return float64(y+x) <= n && float64(y+x) >= -n
}
