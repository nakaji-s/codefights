package main

func contoursShifting(matrix [][]int) [][]int {
	ret := make([][]int, len(matrix))
	for i := range matrix {
		ret[i] = make([]int, len(matrix[i]))
		copy(ret[i], matrix[i])
	}

	y := len(matrix)
	x := len(matrix[0])

	if y < x && y%2 == 1 {
		k := y / 2
		if k%2 == 0 {
			for i := k + 1; i < x-k; i++ {
				ret[k][i] = matrix[k][i-1]
			}
			ret[k][k] = matrix[k][x-1-k]
		} else {
			for i := k; i < x-k-1; i++ {
				ret[k][i] = matrix[k][i+1]
			}
			ret[k][x-1-k] = matrix[k][k]
		}

	} else if x < y && x%2 == 1 {
		k := x / 2
		if k%2 == 0 {
			for i := 1 + k; i < y-k; i++ {
				ret[i][k] = matrix[i-1][k]
			}
			ret[k][k] = matrix[y-1-k][k]
		} else {
			for i := k; i < y-k-1; i++ {
				ret[i][k] = matrix[i+1][k]
			}
			ret[y-1-k][k] = matrix[k][k]
		}
	}

	contours := 0
	if y < x {
		contours = y / 2
	} else {
		contours = x / 2
	}

	for k := 0; k < contours; k++ {
		if k%2 == 0 {
			for i := k; i < y-k; i++ {
				if i == k {
					for j := 1 + k; j < x-k; j++ {
						ret[k][j] = matrix[k][j-1]
					}
				} else {
					ret[i][x-k-1] = matrix[i-1][x-k-1]
				}

				if i == y-1-k {
					for j := k; j < x-k-1; j++ {
						ret[i][j] = matrix[i][j+1]
					}
				} else {
					ret[i][k] = matrix[i+1][k]
				}
			}
		} else {
			for i := k; i < y-k; i++ {
				if i == k {
					for j := k; j < x-k-1; j++ {
						ret[k][j] = matrix[k][j+1]
					}
				} else {
					ret[i][k] = matrix[i-1][k]
				}

				if i == y-1-k {
					for j := k + 1; j < x-k; j++ {
						ret[i][j] = matrix[i][j-1]
					}
				} else {
					ret[i][x-k-1] = matrix[i+1][x-k-1]
				}
			}
		}
	}

	return ret
}
