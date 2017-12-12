package main

func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
	height := 0
	for i := 1; ; i++ {
		height += upSpeed
		if height >= desiredHeight {
			return i
		}
		height -= downSpeed
	}
	return -1

}
