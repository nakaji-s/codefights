package main

func removeArrayPart(inputArray []int, l int, r int) []int {
	return append(inputArray[:l], inputArray[r+1:]...)
}
