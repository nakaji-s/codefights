package main

import "strings"

func findEmailDomain(address string) string {
	str := strings.Split(address, "@")
	return str[len(str)-1]
}
