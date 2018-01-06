package main

func leastFactorial(n int) int {
	sum := 1
	for i := 1; sum < n; i++ {
		sum *= i
	}

	return sum
}
