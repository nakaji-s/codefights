package main

import (
	"strconv"
	"strings"
)

func higherVersion(ver1 string, ver2 string) bool {
	ver1s := strings.Split(ver1, ".")
	ver2s := strings.Split(ver2, ".")

	for i, ver := range ver1s {
		ver1, _ := strconv.Atoi(ver)
		ver2, _ := strconv.Atoi(ver2s[i])
		if ver1 > ver2 {
			return true
		} else if ver1 < ver2 {
			return false
		}
	}

	return false
}
