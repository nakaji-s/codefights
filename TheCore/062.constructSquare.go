package main

import (
	"fmt"
	"math"
)

func constructSquare(s string) int {
	digits := len(s)
	maxNum := math.Pow10(digits) - 1
	maxSqrt := math.Sqrt(maxNum)

	for i := int(maxSqrt); i > 0; i-- {
		squareNumber := i * i
		if len(fmt.Sprint(squareNumber)) < digits {
			break
		}

		if stringsToMapNum(s) == stringsToMapNum(fmt.Sprint(squareNumber)) {
			return squareNumber
		}
	}

	return -1
}

func stringsToMapNum(s string) int {
	sMap := map[rune]int{}
	for _, rune := range s {
		sMap[rune]++
	}
	numMap := map[int]int{}
	for _, num := range sMap {
		numMap[num]++
	}

	mapNum := 0
	for key, value := range numMap {
		mapNum += value * int(math.Pow10(key))
	}

	return mapNum
}
