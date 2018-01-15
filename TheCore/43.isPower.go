package main

import "math"

func isPower(n int) bool {
	m := map[int]int{}

	for j := 2; j < 9; j++ { // because 2^9 == 512 > 400
		for i := 1; ; i++ {
			ret := int(math.Pow(float64(i), float64(j)))
			if ret > 400 {
				break
			}

			m[ret] = n
		}
	}

	if _, ok := m[n]; ok {
		return true
	}
	return false
}
