package main

func cyclicString(s string) int {
	for i := 0; i < len(s); i++ {
		subStr := s[:i+1]
		for j, runeValue := range subStr {
			if i+j+1 >= len(s) {
				return i + 1
			}

			if s[i+j+1] != byte(runeValue) {
				break
			} else {
				if j+1 == len(subStr) {
					return i + 1
				}
			}
		}
	}

	return 0
}
