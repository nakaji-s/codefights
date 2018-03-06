package main

func gravitation(rows []string) []int {
	turns := make([]int, len(rows[0]))
	for i := 0; i < len(turns); i++ {
		turns[i] = 0
	}

	for i := 0; i < len(rows[0]); i++ {
		count := 0
		maxBlank := 0
		blank := 0
		foundFirstBlock := false

		for j := len(rows) - 1; j >= 0; j-- {
			if rows[j][i] == '.' {
				if foundFirstBlock == false {
					count++
				} else {
					blank++
				}
			} else {
				if maxBlank < blank {
					maxBlank = blank
				}
				turns[i] = count
				foundFirstBlock = true
				blank = 0
			}
		}
		turns[i] = turns[i] + maxBlank
	}

	minTurns := []int{}
	minTurn := 10
	for i := 0; i < len(turns); i++ {
		if minTurn > turns[i] {
			minTurns = []int{i}
			minTurn = turns[i]
		} else if minTurn == turns[i] {
			minTurns = append(minTurns, i)
		}
	}

	return minTurns
}
