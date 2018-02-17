package main

func stringsCrossover(inputArray []string, result string) int {
	count := 0
	for i := 0; i < len(inputArray); i++ {
		for j := i + 1; j < len(inputArray); j++ {
			if canForm(inputArray[i], inputArray[j], result) {
				count++
			}
		}
	}

	return count
}

func canForm(s1, s2, result string) bool {
	for i := 0; i < len(result); i++ {
		if result[i] != s1[i] && result[i] != s2[i] {
			return false
		}
	}

	return true
}
