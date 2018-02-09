package main

import "strconv"

func characterParity(symbol string) string {
	if num, err := strconv.Atoi(symbol); err == nil {
		if num%2 == 0 {
			return "even"
		} else {
			return "odd"
		}
	} else {
		return "not a digit"
	}
}
