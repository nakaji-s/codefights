package main

import "strings"

func properNounCorrection(noun string) string {
	return strings.Title(strings.ToLower(noun))
}
