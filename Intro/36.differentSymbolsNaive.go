package main

func differentSymbolsNaive(s string) int {
	m := map[rune]struct{}{}
	for _, rune := range s {
		m[rune] = struct{}{}
	}

	return len(m)
}
