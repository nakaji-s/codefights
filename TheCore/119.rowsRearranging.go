package main

import "sort"

func rowsRearranging(matrix [][]int) bool {
	sort.Slice(matrix, func(i, j int) bool {
		return matrix[i][0] < matrix[j][0]
	})

	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i+1][j] <= matrix[i][j] {
				return false
			}
		}
	}

	return true
}
