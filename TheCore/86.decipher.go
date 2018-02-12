package main

import "strconv"

func decipher(cipher string) string {
	out := ""
	tmp := []rune{}

	for _, runeValue := range cipher {
		tmp = append(tmp, runeValue)
		num, _ := strconv.Atoi(string(tmp))
		if num >= 'a' && num <= 'z' {
			tmp = []rune{}
			out = out + string(num)
		}
	}

	return out
}
