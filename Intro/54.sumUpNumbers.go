package main

import (
	"strconv"
	"strings"
	"unicode"
)

func sumUpNumbers(inputString string) int {
	word := []rune{}
	for _, rune := range inputString {
		if unicode.IsDigit(rune) {
			word = append(word, rune)
		} else {
			word = append(word, ',')
		}
	}
	str := string(word)
	sum := 0

	s := strings.Split(str, ",")
	for _, st := range s {
		n, err := strconv.Atoi(st)
		if err != nil {
			continue
		}
		sum += n
	}
	return sum
}
