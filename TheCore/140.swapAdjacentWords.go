package main

import "regexp"

func swapAdjacentWords(s string) string {
	return regexp.MustCompile(`(.+?) (.+?)( |$)`).ReplaceAllString(s, `$2 $1$3`)
}
