package main

func isBeautifulString(inputString string) bool {
	m := map[rune]int{}

	for _, s := range inputString {
		m[s]++
	}
	num := 50
	for i := 'a'; i <= 'z'; i++ {
		if num >= m[i] {
			num = m[i]
		} else {
			return false
		}
	}
	return true
}
