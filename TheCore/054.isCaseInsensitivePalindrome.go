package main

import "strings"

func isCaseInsensitivePalindrome(inputString string) bool {
	lowerStr := strings.ToLower(inputString)
	s1 := lowerStr[0 : len(lowerStr)/2]
	s2 := lowerStr[(len(lowerStr)+1)/2 : len(lowerStr)]
	s2Reverse := Reverse(s2)

	return s1 == s2Reverse
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
