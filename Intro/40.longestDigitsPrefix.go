package main

import "unicode"

func longestDigitsPrefix(inputString string) string {
	max := 0
	streakStr := ""
	maxStreakStr := ""
	for i, rune := range inputString {
		if unicode.IsDigit(rune) {
			streakStr = streakStr + string(rune)
			if max < i {
				max = i
				maxStreakStr = streakStr
			}
		} else {
			return maxStreakStr
		}
	}
	return maxStreakStr
}
