package main

func volleyballPositions(formation [][]string, k int) [][]string {
	playerName := []string{
		formation[3][2],
		formation[1][2],
		formation[0][1],
		formation[1][0],
		formation[3][0],
		formation[2][1],
	}

	formation[3][2] = playerName[(6-k%6)%6]
	formation[1][2] = playerName[(6-(k-1)%6)%6]
	formation[0][1] = playerName[(6-(k-2)%6)%6]
	formation[1][0] = playerName[(6-(k-3)%6)%6]
	formation[3][0] = playerName[(6-(k-4)%6)%6]
	formation[2][1] = playerName[(6-(k-5)%6)%6]

	return formation
}
