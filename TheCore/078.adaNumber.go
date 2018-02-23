package main

import (
	"strconv"
	"strings"
)

func adaNumber(line string) bool {
	tmp := strings.Replace(line, "_", "", -1)
	chunks := strings.Split(tmp, "#")

	if len(chunks)%2 == 0 {
		return false
	}

	if len(chunks) == 1 {
		for _, rune := range chunks[0] {
			if rune < '0' || rune > '9' {
				return false
			}
		}
		if len(chunks[0]) == 0 {
			return false
		}

		return true
	}

	if base, err := strconv.ParseInt(chunks[0], 10, 32); err != nil {
		return false
	} else {
		if base > 16 || base < 2 {
			return false
		}
		for _, rune := range chunks[1] {
			if rightNum, err := strconv.ParseInt(string(rune), int(base), 32); err != nil {
				return false
			} else {
				if rightNum >= base {
					return false
				}
			}
		}
		if len(chunks[1]) == 0 {
			return false
		}
	}

	return true
}
