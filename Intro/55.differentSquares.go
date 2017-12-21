package main

func differentSquares(matrix [][]int) int {
	m := map[int]struct{}{}

	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			m[matrixToInt(matrix, i, j)] = struct{}{}
		}
	}
	return len(m)
}

func matrixToInt(matrix [][]int, x, y int) int {
	return matrix[x][y] + 10*matrix[x+1][y] + matrix[x][y+1]*100 + matrix[x+1][y+1]*1000
}
