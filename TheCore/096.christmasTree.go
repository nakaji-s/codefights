package main

import "strings"

func christmasTree(levelNum int, levelHeight int) []string {
	ret := []string{}

	// crown
	ret = append(ret, strings.Repeat(" ", levelHeight+levelNum-1)+" *")
	ret = append(ret, strings.Repeat(" ", levelHeight+levelNum-1)+" *")
	ret = append(ret, strings.Repeat(" ", levelHeight+levelNum-1)+"***")

	// level
	for i := 0; i < levelNum; i++ {
		for j := 0; j < levelHeight; j++ {
			ret = append(ret, strings.Repeat(" ", levelHeight-j+(levelNum-i-1)-1)+"*****"+strings.Repeat("*", (j+i)*2))
		}
	}

	// foot
	longest := len(ret[len(ret)-1])
	for i := 0; i < levelNum; i++ {
		ret = append(ret, strings.Repeat(" ", (longest-levelHeight)/2)+strings.Repeat("*", levelHeight+(levelHeight+1)%2))
	}

	return ret
}
