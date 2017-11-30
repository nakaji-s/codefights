package main

func commonCharacterCount(s1 string, s2 string) int {
	m := map[string]int{}

	for _, rune := range s1 {
		m[string(rune)] += 1
	}

	num := 0
	for _, rune := range s2 {
		if m[string(rune)] > 0 {
			num++
			m[string(rune)]--
		}
	}

	return num
}
