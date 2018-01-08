package main

func appleBoxes(k int) int {
	sum := 0
	for i := 1; i <= k; i++ {
		tmp := i * i
		if i%2 == 1 {
			tmp *= -1
		}
		sum += tmp
	}

	return sum
}
