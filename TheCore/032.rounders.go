package main

import "math"

func rounders(value int) int {
	num := float64(value)
	i := 0
	for ; num >= 10; i++ {
		num /= 10
		num = math.Floor(num + 0.5)
	}

	return int(num * math.Pow10(i))
}
