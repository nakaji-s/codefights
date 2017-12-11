package main

import "unicode"

func firstDigit(inputString string) string {
	for _, rune := range inputString {
		if unicode.IsDigit(rune) {
			return string(rune)
		}
	}
	return ""
}
