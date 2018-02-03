package main

import (
	"fmt"
	"strings"
)

func addBorder(picture []string) []string {
	ret := []string{}
	maxLen := 0
	for _, str := range picture {
		ret = append(ret, fmt.Sprint("*", str, "*"))
		if maxLen < len(str) {
			maxLen = len(str)
		}
	}
	frame := strings.Repeat("*", maxLen+2)
	ret = append([]string{frame}, ret...)
	ret = append(ret, frame)
	return ret
}
