package main

import "strconv"

func treeBottom(tree string) []int {
	depth := 0
	maxDepth := 0
	ret := []int{}

	str := []byte{}
	for _, character := range tree {
		switch byte(character) {
		case '(':
			depth++
		case ')':
			depth--
		case ' ':
			if len(str) > 0 {
				num, _ := strconv.ParseInt(string(str), 10, 32)
				str = []byte{}

				if maxDepth < depth {
					ret = []int{int(num)}
					maxDepth = depth
				} else if maxDepth == depth {
					ret = append(ret, int(num))
				}
			}
		default:
			str = append(str, byte(character))
		}
	}

	return ret
}
