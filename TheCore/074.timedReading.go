package main

import "strings"
import "unicode"

func timedReading(maxLength int, text string) int {
	count := 0
	for _, word := range strings.Split(text, " ") {
		characters := 0
		for _, rune := range word {
			if unicode.IsLetter(rune) {
				characters++
			}
		}

		if characters <= maxLength && characters > 0 {
			count++
		}
	}

	return count
}
