package main

func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	max := 0
	if maxW >= weight1 {
		max = value1
	}
	if maxW >= weight2 {
		if max < value2 {
			max = value2
		}
	}
	if maxW >= weight1+weight2 {
		max = value1 + value2
	}

	return max
}
