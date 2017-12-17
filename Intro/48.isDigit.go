package main

import "unicode"

func isDigit(symbol string) bool {
	for _, rune := range symbol {
		return unicode.IsDigit(rune)
	}
	return false
}
