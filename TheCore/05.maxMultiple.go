package main

func maxMultiple(divisor int, bound int) int {
	for i := bound; i > 0; i-- {
		if i%divisor == 0 {
			return i
		}
	}
	return -1
}
