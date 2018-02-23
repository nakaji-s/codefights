package main

func lineUp(commands string) int {
	countLR := 0
	countSameDirection := 0

	for _, command := range commands {
		switch command {
		case 'L', 'R':
			countLR++
		}

		countSameDirection += (countLR + 1) % 2
	}

	return countSameDirection
}
