package main

import "fmt"

func firstOperationCharacter(expr string) int {
	priotiry := 0
	maxIndex := -1
	maxChar := '+'
	maxPriority := -1

	for i, rune := range expr {
		switch byte(rune) {
		case '(':
			priotiry++
		case ')':
			priotiry--
		case '+':
			if maxPriority < priotiry {
				maxPriority = priotiry
				maxIndex = i
				maxChar = '+'
			}
		case '*':
			if maxPriority < priotiry {
				maxPriority = priotiry
				maxIndex = i
				maxChar = '*'
			} else if maxPriority == priotiry && maxChar == '+' {
				maxPriority = priotiry
				maxIndex = i
				maxChar = '*'
			}
		}
	}

	fmt.Println(string(maxChar), maxIndex, maxPriority)
	return maxIndex
}
