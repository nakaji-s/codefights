package main

import (
	"strings"
	"unicode"
)

func longestWord(text string) string {
	word := []rune{}
	for _, rune := range text {
		if unicode.IsLetter(rune) {
			word = append(word, rune)
		} else {
			word = append(word, ',')
		}
	}
	str := string(word)

	s := strings.Split(str, ",")

	longestWord := ""
	for _, st := range s {
		if len(longestWord) < len(st) {
			longestWord = st
		}
	}

	return longestWord
}
