package main

import "strings"

func buildPalindrome(st string) string {
	rv := Reverse(st)
	for i := 0; i < len(rv); i++ {
		if strings.HasSuffix(st, rv[:len(rv)-i]) {
			return st + Reverse(strings.TrimSuffix(st, rv[:len(rv)-i]))
		}
	}
	return ""
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
