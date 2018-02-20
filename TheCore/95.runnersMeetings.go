package main

func runnersMeetings(startPosition []int, speed []int) int {
	maxCross := 0
	for i := 0; i < len(startPosition); i++ {
		m := map[float64]int{}

		for j := i + 1; j < len(startPosition); j++ {
			if speed[i] == speed[j] {
				continue
			}
			position, time := getCrossing(startPosition[i], startPosition[j], speed[i], speed[j])
			if time < 0 {
				continue
			}
			m[position]++
			if maxCross < m[position] {
				maxCross = m[position]
			}
		}
	}

	if maxCross != 0 {
		return maxCross + 1
	}
	return -1
}

func getCrossing(startPositionA, startPositionB, speedA, speedB int) (float64, float64) {
	time := float64(startPositionB-startPositionA) / (float64(speedA-speedB) / 60)
	position := float64(speedA)*time/60 + float64(startPositionA)

	if time < 0 && position < 0 {
		return -position, -time
	}
	return position, time
}
