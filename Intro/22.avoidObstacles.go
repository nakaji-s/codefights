package main

func avoidObstacles(inputArray []int) int {
	for i := 2; i <= 41; i++ {
		passed := 0
		for j := 0; j < len(inputArray); j++ {
			if inputArray[j]%i != 0 {
				passed++
				continue
			}
		}
		if passed == len(inputArray) {
			return i
		}
	}
	return 0
}
