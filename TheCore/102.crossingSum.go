package main

func crossingSum(matrix [][]int, a int, b int) int {
	sum := 0
	for column, num := range matrix[a] {
		if column == b {
			continue
		}
		sum += num
	}

	for _, row := range matrix {
		sum += row[b]
	}

	return sum
}
