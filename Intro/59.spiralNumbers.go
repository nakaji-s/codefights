package main

func spiralNumbers(n int) [][]int {
	direction := 0
	x := 0
	y := 0

	field := make([][]int, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, n)
	}
	field[y][x] = 1
	for i := 0; i < n*n-1; i++ {
		switch direction {
		case 0:
			if x < n-1 && field[y][x+1] == 0 {
				x++
			} else {
				i--
				direction = 1
				continue
			}
		case 1:
			if y < n-1 && field[y+1][x] == 0 {
				y++
			} else {
				i--
				direction = 2
				continue
			}
		case 2:
			if x > 0 && field[y][x-1] == 0 {
				x--
			} else {
				i--
				direction = 3
				continue
			}
		case 3:
			if y > 0 && field[y-1][x] == 0 {
				y--
			} else {
				i--
				direction = 0
				continue
			}
		}
		field[y][x] = i + 2

	}

	return field
}
