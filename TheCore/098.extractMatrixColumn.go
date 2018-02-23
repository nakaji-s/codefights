package main

func extractMatrixColumn(matrix [][]int, column int) []int {
	ret := []int{}
	for _, line := range matrix {
		ret = append(ret, line[column])
	}

	return ret
}
