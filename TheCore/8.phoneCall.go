package main

func phoneCall(min1 int, min2_10 int, min11 int, s int) int {
	for i := 0; i < 60; i++ {
		if i == 0 {
			s = s - min1
		} else if i < 10 {
			s = s - min2_10
		} else {
			s = s - min11
		}
		if s < 0 {
			return i
		}
	}
	return -1
}
