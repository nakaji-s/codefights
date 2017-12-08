package main

func chessBoardCellColor(cell1 string, cell2 string) bool {
	return getColor(cell1) == getColor(cell2)
}

func getColor(cell string) bool {
	return (int((cell[0]-'A'+1)+(cell[1]-'0')) % 2) == 1
}
