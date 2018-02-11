package main

import "unicode"

func stolenLunch(note string) string {
	ret := []rune{}
	for _, runeValue := range note {
		if unicode.IsDigit(runeValue) {
			ret = append(ret, rune(int(runeValue)-'0'+'a'))
		} else {
			if int(runeValue) >= 'a' && int(runeValue) <= 'j' {
				ret = append(ret, rune(int(runeValue)-'a'+'0'))
			} else {
				ret = append(ret, runeValue)
			}
		}
	}

	return string(ret)
}
