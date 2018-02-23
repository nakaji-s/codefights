package main

import (
	"fmt"
	"strings"
)

func integerToStringOfFixedWidth(number int, width int) string {
	ret := fmt.Sprint(number)
	if len(ret) > width {
		return ret[len(ret)-width:]
	} else {
		return strings.Repeat("0", width-len(ret)) + ret
	}
}
