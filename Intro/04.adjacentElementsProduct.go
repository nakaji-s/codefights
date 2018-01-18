package main

func adjacentElementsProduct(inputArray []int) int {
	max := -10000
	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i]*inputArray[i+1] > max {
			max = inputArray[i] * inputArray[i+1]
		}
	}
	return max
}
