package main

import "fmt"

func combs(comb1 string, comb2 string) int {
	n1 := getMinLen(comb1, comb2)
	n2 := getMinLen(comb2, comb1)
	if n1 > n2 {
		return n2
	}
	return n1
}

func getMinLen(comb1 string, comb2 string) int {
	minLen := len(comb1) + len(comb2)

OUTER:
	for i := 0; i < len(comb1); i++ {
		length := -1
		for j := 0; j < len(comb2); j++ {
			sum := comb2[j]
			if i+j < len(comb1) {
				sum += comb1[i+j]
			} else {
				length = i + len(comb2)
				break
			}
			if sum != '*'+'.' && sum != '.'+'.' {
				fmt.Println(i, j)
				continue OUTER
			}
			length = i + j
		}
		if minLen > length {
			minLen = length
		}
	}

	if minLen < len(comb1) {
		minLen = len(comb1)
	}
	if minLen < len(comb2) {
		minLen = len(comb2)
	}

	return minLen
}
