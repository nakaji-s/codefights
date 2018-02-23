package main

import (
	"fmt"
	"strconv"
)

func additionWithoutCarrying(param1 int, param2 int) int {
	param1StrRev := Reverse(fmt.Sprint(param1))
	param2StrRev := Reverse(fmt.Sprint(param2))
	sumRev := ""
	for i := 0; i < len(param1StrRev) && i < len(param2StrRev); i++ {
		tmp := ([]rune(param1StrRev)[i] - '0' +
			[]rune(param2StrRev)[i] - '0') % 10
		sumRev += fmt.Sprint(tmp)
	}
	if len(param1StrRev) > len(param2StrRev) {
		sumRev += param1StrRev[len(param2StrRev):]
	} else if len(param1StrRev) < len(param2StrRev) {
		sumRev += param2StrRev[len(param1StrRev):]
	}

	ret, _ := strconv.Atoi(Reverse(sumRev))
	return ret
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
