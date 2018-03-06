package main

func polygonPerimeter(matrix [][]bool) int {
	ret := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == true {
				ret += 4
				if i > 0 && matrix[i-1][j] == true {
					ret--
				}
				if i < len(matrix)-1 && matrix[i+1][j] == true {
					ret--
				}
				if j > 0 && matrix[i][j-1] == true {
					ret--
				}
				if j < len(matrix[0])-1 && matrix[i][j+1] == true {
					ret--
				}
			}
		}
	}

	return ret
}
