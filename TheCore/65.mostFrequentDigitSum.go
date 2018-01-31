package main

import "fmt"

func mostFrequentDigitSum(n int) int {
	m := map[int]int{}
	firstSequenceNum := n
	for firstSequenceNum > 0 {
		tmp := fmt.Sprint(firstSequenceNum)
		secondSequenceNum := 0
		for _, rune := range tmp {
			secondSequenceNum += int(rune) - '0'
		}

		firstSequenceNum -= secondSequenceNum
		m[secondSequenceNum]++
	}

	mostFrequentNum := 0
	mostFrequentCount := 0
	for num, count := range m {
		if mostFrequentCount <= count {
			mostFrequentCount = count
			mostFrequentNum = num
		}
	}

	return mostFrequentNum
}
