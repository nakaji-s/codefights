package main

import (
	"fmt"
	"regexp"
)

func nthNumber(s string, n int) string {
	re := regexp.MustCompile(`(?:\d+\D+){` + fmt.Sprint(n-1) + `}[0]*([1-9]\d*)`)
	return re.FindStringSubmatch(s)[1]
}
