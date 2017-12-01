package main

import "regexp"

func reverseParentheses(str string) string {
	re, _ := regexp.Compile(`\([^\(\)]*\)`)

	cb := func(s string) string {
		return Reverse(s[1 : len(s)-1])
	}
	for {
		tmp := str
		str = re.ReplaceAllStringFunc(str, cb)
		if tmp == str {
			break
		}
	}

	return str
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
