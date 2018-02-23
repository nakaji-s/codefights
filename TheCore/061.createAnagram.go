package main

func createAnagram(s string, t string) int {
	diffMap := map[rune]int{}

	for _, rune := range s {
		diffMap[rune]++
	}
	for _, rune := range t {
		diffMap[rune]--
	}

	diffCount := 0
	for _, num := range diffMap {
		if num < 0 {
			diffCount -= num
		}
	}

	return diffCount
}
