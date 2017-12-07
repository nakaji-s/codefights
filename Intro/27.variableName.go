package main

import "unicode"

func variableName(name string) bool {
	if unicode.IsDigit(rune(name[0])) {
		return false
	}
	for _, rune := range name {
		if !(unicode.IsDigit(rune) ||
			unicode.IsLetter(rune) ||
			rune == '_') {
			return false
		}
	}
	return true
}
