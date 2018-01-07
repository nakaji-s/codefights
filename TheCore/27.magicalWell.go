package main

func magicalWell(a int, b int, n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += (a + i) * (b + i)
	}

	return sum
}
