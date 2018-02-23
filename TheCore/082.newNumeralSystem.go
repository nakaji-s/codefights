package main

import "fmt"

func newNumeralSystem(number string) []string {
	ret := []string{}

	for i := 0; ; i++ {
		firstChar := byte('A' + i)
		secondChar := number[0] - byte(i)

		str := fmt.Sprintf("%s + %s", string(firstChar), string(secondChar))
		ret = append(ret, str)

		if firstChar+1 >= secondChar {
			break
		}
	}

	return ret
}
