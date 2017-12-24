package main

func sudoku(grid [][]int) bool {
	for i := 0; i < len(grid); i += 3 {
		for j := 0; j < len(grid[i]); j += 3 {
			if isCorrect3x3(grid, i, j) == false {
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		prodX := 1
		prodY := 1
		for j := 0; j < 9; j++ {
			prodX *= grid[i][j]
			prodY *= grid[j][i]
		}

		if prodX != 362880 {
			return false
		}
		if prodY != 362880 {
			return false
		}
	}
	return true
}

func isCorrect3x3(grid [][]int, x, y int) bool {
	prod := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			prod *= grid[x+i][y+j]
		}
	}
	return prod == 362880
}
