package main

func stringsConstruction(a string, b string) int {
	aMap := map[rune]int{}
	bMap := map[rune]int{}

	for _, rune := range a {
		aMap[rune]++
	}
	for _, rune := range b {
		bMap[rune]++
	}

	minConstruct := 50
	for rune, num := range aMap {
		if bMap[rune] == 0 {
			return 0
		}

		tmp := bMap[rune] / num
		if tmp < minConstruct {
			minConstruct = tmp
		}
	}

	return minConstruct
}
