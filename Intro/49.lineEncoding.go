package main

import "fmt"

func lineEncoding(s string) string {
	out := ""
	lastStr := ""
	streak := 0
	for _, rune := range s {
		if string(rune) != lastStr {
			if streak > 0 {
				out += fmt.Sprint(streak + 1)
				streak = 0
			}
			out += lastStr
		} else {
			streak++
		}
		lastStr = string(rune)
	}

	if streak > 0 {
		out += fmt.Sprint(streak + 1)
	}
	out += lastStr

	return out
}
