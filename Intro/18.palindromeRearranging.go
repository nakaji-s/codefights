package main

func palindromeRearranging(inputString string) bool {
	m := map[string]int{}
	for _, s := range inputString {
		m[string(s)]++
	}

	odd := 0
	for _, n := range m {
		if n%2 == 1 {
			odd++
			if odd > 1 {
				return false
			}
		}
	}
	return true
}
