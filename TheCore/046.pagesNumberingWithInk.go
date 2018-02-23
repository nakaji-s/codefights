package main

import "fmt"

func pagesNumberingWithInk(current int, numberOfDigits int) int {
	for {
		numLen := len(fmt.Sprint(current))
		if numberOfDigits-numLen < 0 {
			break
		}

		numberOfDigits -= numLen
		current++
	}

	return current - 1
}
