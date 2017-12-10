package main

func extractEachKth(inputArray []int, k int) []int {
	for i := len(inputArray) - 1; i >= 0; i-- {
		if (i+1)%k == 0 {
			inputArray = append(inputArray[:i], inputArray[i+1:]...)
		}
	}
	return inputArray
}
