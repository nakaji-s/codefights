package main

func chessKnightMoves(cell string) int {
	sum := 0
	x := int(cell[0] - 'a')
	y := int(cell[1] - '1')

	if inBoard(x+1, y+2) {
		sum++
	}
	if inBoard(x+1, y-2) {
		sum++
	}
	if inBoard(x+2, y+1) {
		sum++
	}
	if inBoard(x+2, y-1) {
		sum++
	}
	if inBoard(x-1, y+2) {
		sum++
	}
	if inBoard(x-1, y-2) {
		sum++
	}
	if inBoard(x-2, y+1) {
		sum++
	}
	if inBoard(x-2, y-1) {
		sum++
	}
	return sum
}

func inBoard(x, y int) bool {
	return x >= 0 && x <= 7 && y >= 0 && y <= 7
}
