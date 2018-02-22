package main

func reverseOnDiagonals(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			if i == j {
				matrix[i][j], matrix[len(matrix)-1-i][len(matrix)-1-j] = matrix[len(matrix)-1-i][len(matrix)-1-j], matrix[i][j]
			} else if i+j == len(matrix)-1 {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
	return matrix
}
