package main

func alphabetSubsequence(s string) bool {
	lastRune := 'a' - 1
	for _, rune := range s {
		if rune <= lastRune {
			return false
		}
		lastRune = rune
	}

	return true
}
