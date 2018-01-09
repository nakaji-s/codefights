package main

import "fmt"

func increaseNumberRoundness(n int) bool {
	str := fmt.Sprint(n)
	isZeroExist := false

	for _, digit := range str {
		if digit == '0' {
			isZeroExist = true
		} else {
			if isZeroExist {
				return true
			}
		}
	}

	return false
}
