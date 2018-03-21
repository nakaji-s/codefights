package main

type situation struct {
	positionY  int
	positionX  int
	directionY int
	directionX int
}

func chessBishopDream(boardSize []int, initPosition []int, initDirection []int, k int) []int {
	currentPosition := initPosition
	currentDirection := initDirection

	m := map[situation]int{}
	situations := []situation{}

	for i := 0; i < k; i++ {
		currentSituation := situation{
			currentPosition[0],
			currentPosition[1],
			currentDirection[0],
			currentDirection[1],
		}
		if _, ok := m[currentSituation]; ok {
			tmp := situations[k%len(situations)]
			return []int{tmp.positionY, tmp.positionX}
		}
		m[currentSituation] = k
		situations = append(situations, currentSituation)

		nextY := currentPosition[0] + currentDirection[0]
		if nextY >= boardSize[0] || nextY < 0 {
			nextY = currentPosition[0]
			currentDirection[0] *= -1
		}
		nextX := currentPosition[1] + currentDirection[1]
		if nextX >= boardSize[1] || nextX < 0 {
			nextX = currentPosition[1]
			currentDirection[1] *= -1
		}

		currentPosition = []int{nextY, nextX}
	}

	return currentPosition
}
