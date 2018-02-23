package main

import "fmt"

func fileNaming(names []string) []string {
	ret := []string{}
	m := map[string]int{}
	for _, name := range names {
		fileName := name

		for i := 0; ; i++ {
			n, exist := m[fileName]
			if exist == true {
				fileName = name + fmt.Sprint("(", n+i, ")")
			} else {
				break
			}
		}
		m[fileName]++
		ret = append(ret, fileName)
	}
	return ret
}
