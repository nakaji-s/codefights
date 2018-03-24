package main

func amazonCheckmate(king string, amazon string) []int {
	field := make([][]int, 8)
	for i := 0; i < 8; i++ {
		field[i] = make([]int, 8)
	}

	kingX := int(king[0] - 'a')
	kingY := int(king[1] - '1')
	amazonX := int(amazon[0] - 'a')
	amazonY := int(amazon[1] - '1')

	// amazon and king
	field[amazonY][amazonX] = -8
	field[kingY][kingX] = -9

	// king's field
	changeFieldValue(field, kingX-1, kingY-1, -1)
	changeFieldValue(field, kingX, kingY-1, -1)
	changeFieldValue(field, kingX+1, kingY-1, -1)
	changeFieldValue(field, kingX-1, kingY, -1)
	changeFieldValue(field, kingX+1, kingY, -1)
	changeFieldValue(field, kingX-1, kingY+1, -1)
	changeFieldValue(field, kingX, kingY+1, -1)
	changeFieldValue(field, kingX+1, kingY+1, -1)

	// check queen's field
	checkQueensField(field, amazonX, amazonY, 1)

	// check knight's field
	changeFieldValue(field, amazonX-1, amazonY-2, 1)
	changeFieldValue(field, amazonX+1, amazonY-2, 1)
	changeFieldValue(field, amazonX+2, amazonY-1, 1)
	changeFieldValue(field, amazonX+2, amazonY+1, 1)
	changeFieldValue(field, amazonX+1, amazonY+2, 1)
	changeFieldValue(field, amazonX-1, amazonY+2, 1)
	changeFieldValue(field, amazonX-2, amazonY-1, 1)
	changeFieldValue(field, amazonX-2, amazonY+1, 1)

	ret := []int{0, 0, 0, 0}
	for y, line := range field {
		for x, value := range line {
			if value == 0 {
				if cantMove(field, x, y) {
					ret[2]++
					//field[y][x] = 7
				} else {
					ret[3]++
				}
			}
			if value == 1 {
				if cantMove(field, x, y) {
					ret[0]++
					//field[y][x] = 8
				} else {
					ret[1]++
				}
			}
		}
	}

	// for i := 7; i >= 0; i-- {
	//     for j := 0; j < 8; j++ {
	//         fmt.Printf("%3d", field[i][j])
	//     }
	//     fmt.Println()
	// }

	return ret
}

func cantMove(field [][]int, x, y int) bool {
	if isFieldValue(field, x-1, y+1, 0) ||
		isFieldValue(field, x, y+1, 0) ||
		isFieldValue(field, x+1, y+1, 0) ||
		isFieldValue(field, x+1, y, 0) ||
		isFieldValue(field, x+1, y-1, 0) ||
		isFieldValue(field, x, y-1, 0) ||
		isFieldValue(field, x-1, y-1, 0) ||
		isFieldValue(field, x-1, y, 0) ||

		isFieldValue(field, x-1, y+1, -8) ||
		isFieldValue(field, x, y+1, -8) ||
		isFieldValue(field, x+1, y+1, -8) ||
		isFieldValue(field, x+1, y, -8) ||
		isFieldValue(field, x+1, y-1, -8) ||
		isFieldValue(field, x, y-1, -8) ||
		isFieldValue(field, x-1, y-1, -8) ||
		isFieldValue(field, x-1, y, -8) {
		return false
	}

	return true
}

func checkQueensField(field [][]int, amazonX, amazonY, value int) {
	for i := 1; ; i++ {
		if !changeFieldValue(field, amazonX-i, amazonY, value) {
			break
		}
	}
	for i := 1; ; i++ {
		if !changeFieldValue(field, amazonX-i, amazonY+i, value) {
			break
		}
	}
	for i := 1; ; i++ {
		if !changeFieldValue(field, amazonX, amazonY+i, value) {
			break
		}
	}
	for i := 1; ; i++ {
		if !changeFieldValue(field, amazonX+i, amazonY+i, value) {
			break
		}
	}
	for i := 1; ; i++ {
		if !changeFieldValue(field, amazonX+i, amazonY, value) {
			break
		}
	}
	for i := 1; ; i++ {
		if !changeFieldValue(field, amazonX+i, amazonY-i, value) {
			break
		}
	}
	for i := 1; ; i++ {
		if !changeFieldValue(field, amazonX, amazonY-i, value) {
			break
		}
	}
	for i := 1; ; i++ {
		if !changeFieldValue(field, amazonX-i, amazonY-i, value) {
			break
		}
	}
}

func isFieldValue(field [][]int, x, y, value int) bool {
	if y >= 0 && y < 8 && x >= 0 && x < 8 {
		if field[y][x] == value {
			return true
		}
	}

	return false
}

func changeFieldValue(field [][]int, x, y, value int) bool {
	if y >= 0 && y < 8 && x >= 0 && x < 8 {
		if field[y][x] == -9 {
			return false
		}
		if field[y][x] != -1 {
			field[y][x] = value
		}
		return true
	}
	return false
}
