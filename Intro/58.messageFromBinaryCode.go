package main

import "strconv"

func messageFromBinaryCode(code string) string {
	ret := ""
	for i := 0; i < len(code); i += 8 {
		n, _ := strconv.ParseInt(code[i:i+8], 2, 0)
		ret += string(n)
	}
	return ret
}
