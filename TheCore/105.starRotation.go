package main

func starRotation(matrix [][]int, width int, center []int, t int) [][]int {
	for r := 1; r <= width/2; r++ {
		nums := []int{
			matrix[center[0]-r][center[1]-r],
			matrix[center[0]-r][center[1]],
			matrix[center[0]-r][center[1]+r],
			matrix[center[0]][center[1]+r],
			matrix[center[0]+r][center[1]+r],
			matrix[center[0]+r][center[1]],
			matrix[center[0]+r][center[1]-r],
			matrix[center[0]][center[1]-r],
		}

		matrix[center[0]-r][center[1]-r] = nums[(8+(0-t)%8)%8]
		matrix[center[0]-r][center[1]] = nums[(8+(1-t)%8)%8]
		matrix[center[0]-r][center[1]+r] = nums[(8+(2-t)%8)%8]
		matrix[center[0]][center[1]+r] = nums[(8+(3-t)%8)%8]
		matrix[center[0]+r][center[1]+r] = nums[(8+(4-t)%8)%8]
		matrix[center[0]+r][center[1]] = nums[(8+(5-t)%8)%8]
		matrix[center[0]+r][center[1]-r] = nums[(8+(6-t)%8)%8]
		matrix[center[0]][center[1]-r] = nums[(8+(7-t)%8)%8]
	}

	return matrix
}
