package main

func stringsRearrangement(inputArray []string) bool {
	field := make([][]bool, len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		field[i] = make([]bool, len(inputArray))
	}

	for i := 0; i < len(inputArray); i++ {
		for j := i + 1; j < len(inputArray); j++ {
			if isDiffOneCharacter(inputArray[i], inputArray[j]) {
				field[i][j] = true
				field[j][i] = true
			}
		}
	}

	for i := 0; i < len(inputArray); i++ {
		remaining := []int{}
		for j := 0; j < len(inputArray); j++ {
			if j != i {
				remaining = append(remaining, j)
			}
		}
		if findPath(i, remaining, field) == true {
			return true
		}
	}
	return false
}

func isDiffOneCharacter(s1, s2 string) bool {
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff == 1
}

func findPath(current int, remaining []int, field [][]bool) bool {
	for _, n := range remaining {
		if field[current][n] == true {
			newRemaining := []int{}
			for _, n2 := range remaining {
				if n != n2 {
					newRemaining = append(newRemaining, n2)
				}
			}
			if len(newRemaining) == 0 {
				return true
			}
			if findPath(n, newRemaining, field) == true {
				return true
			}
		}
	}
	return false
}
