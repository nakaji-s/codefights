package main

import "strings"
import "regexp"

func repetitionEncryption(letter string) int {
	pattern := regexp.MustCompile(`[\W\s\d]+`)
	words := pattern.Split(strings.ToLower(letter), -1)
	res := 0
	for i := range words {
		if i > 0 && words[i] == words[i-1] {
			res++
		}
	}
	return res
}
