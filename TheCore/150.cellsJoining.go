package main

func cellsJoining(table []string, coords [][]int) []string {
	currentRow := -1
	firstLineFlag := false

	column := []int{}
	for i, character := range table[0] {
		if byte(character) == '+' {
			column = append(column, i)
		}
	}

	for i, line := range table {
		if line[0] == '+' {
			currentRow++
		}

		if currentRow >= coords[1][0] && currentRow <= coords[0][0]+1 {
			if i == 0 || i == len(table)-1 {
				tmp := []byte(line)
				for j, character := range tmp {
					if j > column[coords[0][1]] && j < column[coords[1][1]+1] {
						if character == '+' {
							tmp[j] = '-'
						}
					}
				}
				table[i] = string(tmp)
			}

			if currentRow == coords[0][0]+1 {
				break
			}

			if !firstLineFlag {
				firstLineFlag = true
				continue
			}

			tmp := []byte(line)
			for j, character := range tmp {
				if j == column[coords[0][1]] && j == 0 {
					tmp[j] = '|'
				} else if j > column[coords[0][1]] && j < column[coords[1][1]+1] {
					if character == '-' || character == '+' || character == '|' {
						tmp[j] = ' '
					}
				} else if j == column[coords[1][1]+1] && j == len(tmp)-1 {
					tmp[j] = '|'
				}
			}
			table[i] = string(tmp)
		}
	}

	return table
}
