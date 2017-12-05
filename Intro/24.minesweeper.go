package main

func minesweeper(matrix [][]bool) [][]int {
	mines := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		mines[i] = make([]int, len(matrix[i]))
	}

	for i, line := range matrix {
		for j, _ := range line {
			if matrix[i][j] == true {
				add(mines, i-1, j-1)
				add(mines, i, j-1)
				add(mines, i+1, j-1)
				add(mines, i-1, j)
				add(mines, i+1, j)
				add(mines, i-1, j+1)
				add(mines, i, j+1)
				add(mines, i+1, j+1)
			}
		}
	}
	return mines
}

func add(mines [][]int, x int, y int) {
	if x >= 0 && x < len(mines) && y >= 0 && y < len(mines[x]) {
		mines[x][y]++
	}
}
