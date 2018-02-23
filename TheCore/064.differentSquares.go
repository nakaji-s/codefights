package main

func differentSquares(matrix [][]int) int {
	m := map[int]struct{}{}
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			sum := 0
			sum += matrix[i][j]
			sum += matrix[i][j+1] * 10
			sum += matrix[i+1][j] * 100
			sum += matrix[i+1][j+1] * 1000
			m[sum] = struct{}{}
		}
	}

	return len(m)
}
