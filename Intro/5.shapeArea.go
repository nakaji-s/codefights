package main

func shapeArea(n int) int {
	if n == 1 {
		return 1
	}
	return shapeArea(n-1) + n*4 - 4
}
