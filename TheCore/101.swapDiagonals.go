package main

func swapDiagonals(matrix [][]int) [][]int {
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix)/2; j++ {
			if i == j {
				matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
			}
		}
	}

	for i := len(matrix)/2 + len(matrix)%2; i < len(matrix); i++ {
		for j := len(matrix)/2 + len(matrix)%2; j < len(matrix); j++ {
			if i == j {
				matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
			}
		}
	}

	return matrix
}
