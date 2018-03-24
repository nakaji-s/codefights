package main

func pawnRace(white string, black string, toMove string) string {
	if white[0] == black[0] && white[1] < black[1] {
		return "draw"
	}

	pawn := [2][]int{}
	pawn[0] = []int{int(white[0] - 'a'), int(white[1] - '1')}
	pawn[1] = []int{int(black[0] - 'a'), int(black[1] - '1')}
	turn := 0
	if toMove == "b" {
		turn = 1
	}

	for {
		if pawn[0][1]+1 == pawn[1][1] &&
			(pawn[0][0]-pawn[1][0])*(pawn[0][0]-pawn[1][0]) == 1 {
			break
		}

		direction := 1
		if turn%2 == 1 {
			direction = -1
			if pawn[1][1] == 6 {
				if pawn[0][1] >= 4 || pawn[0][1] <= 1 {
					return "white"
				}
				return "black"
			}
		} else {
			if pawn[0][1] == 1 {
				if pawn[1][1] <= 3 || pawn[1][1] >= 7 {
					return "black"
				}
				return "white"
			}
		}

		pawn[turn%2][1] += direction
		if pawn[turn%2][1] == 0 || pawn[turn%2][1] == 7 {
			break
		}

		turn++
	}

	if turn%2 == 0 {
		return "white"
	}
	return "black"
}
