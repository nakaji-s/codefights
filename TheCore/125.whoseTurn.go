package main

import "strings"

func whoseTurn(p string) bool {
	config := strings.Split(p, ";")
	return (getSumXY(config[0])+getSumXY(config[1]))%2 == (getSumXY(config[2])+getSumXY(config[3]))%2
}

func getSumXY(s string) int {
	return int(s[0]-'a') + int(s[1]-'1')
}
