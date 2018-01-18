package main

func matrixElementsSum(matrix [][]int) int {
	sum := 0
	for i, x := range matrix {
		for j, num := range x {
			if i > 0 && matrix[i-1][j] == 0 {
				matrix[i][j] = 0
			} else {
				sum += num
			}
		}
	}

	return sum
}
