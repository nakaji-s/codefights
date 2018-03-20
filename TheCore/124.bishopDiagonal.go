package main

import (
	"fmt"
	"sort"
)

func bishopDiagonal(bishop1 string, bishop2 string) []string {
	diffX := int(bishop1[0]) - int(bishop2[0])
	diffY := int(bishop1[1]) - int(bishop2[1])

	slope := float64(diffY) / float64(diffX)
	ret := []string{}
	if slope == -1 || slope == 1 {

		for i := 0; ; i++ {
			x := int(bishop1[0]) - 'a' + 1 + i
			y := int(bishop1[1]) - '1' + 1 + i*int(slope)

			if x == 1 || x == 8 || y == 1 || y == 8 {
				ret = append(ret, fmt.Sprintf("%c%d", x+'a'-1, y))
				break
			}
		}

		for i := 0; ; i++ {
			x := int(bishop2[0]) - 'a' + 1 - i
			y := int(bishop2[1]) - '1' + 1 - i*int(slope)

			if x == 1 || x == 8 || y == 1 || y == 8 {
				ret = append(ret, fmt.Sprintf("%c%d", x+'a'-1, y))
				break
			}
		}
	} else {
		ret = []string{bishop1, bishop2}
	}

	sort.Strings(ret)

	return ret
}
