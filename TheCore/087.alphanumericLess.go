package main

import (
	"fmt"
	"strconv"
)

func alphanumericLess(s1 string, s2 string) bool {
	offset_s1 := 0
	offset_s2 := 0
	trailingZero_s1 := 0
	trailingZero_s2 := 0

	for {
		isLetter_s1, num_s1, tmpTrailingZero, raw_s1, tmpOffset := getNext(s1, offset_s1)
		if trailingZero_s1 == 0 {
			trailingZero_s1 = tmpTrailingZero
		}
		offset_s1 = tmpOffset
		isLetter_s2, num_s2, tmpTrailingZero, raw_s2, tmpOffset := getNext(s2, offset_s2)
		if trailingZero_s2 == 0 {
			trailingZero_s2 = tmpTrailingZero
		}
		offset_s2 = tmpOffset

		if isLetter_s1 && !isLetter_s2 {
			return false
		}
		if !isLetter_s1 && isLetter_s2 {
			return true
		}

		fmt.Println(num_s1, num_s2, offset_s1, offset_s2)
		if num_s1 == num_s2 {
			if offset_s1 >= len(s1) && offset_s2 >= len(s2) {
				if raw_s1 == raw_s2 {
					break
				}

				if trailingZero_s1 != trailingZero_s2 {
					return trailingZero_s1 > trailingZero_s2
				}

				if len(raw_s1) != len(raw_s2) {
					return len(raw_s1) < len(raw_s2)
				}
				for i := 0; i < len(raw_s1); i++ {
					if raw_s1[i] != raw_s2[i] {
						return raw_s1[i] < raw_s2[i]
					}
				}
			}
			continue
		}

		return num_s1 < num_s2
	}

	return false
}

func getNext(s string, offset int) (bool, int64, int, string, int) {
	num := int64(0)
	isLetter := false
	trailingZero := 0
	zeroStreaking := false

	tmp := []byte{}
	for offset < len(s) {
		if s[offset] >= 'a' && s[offset] <= 'z' {
			if len(tmp) > 0 {
				break
			}
			num = int64(s[offset])
			isLetter = true
			offset++
			break
		} else {
			if s[offset] == '0' && (zeroStreaking || len(tmp) == 0) {
				trailingZero++
				zeroStreaking = true
			} else {
				zeroStreaking = false
			}
			tmp = append(tmp, s[offset])
			offset++
		}
	}

	if len(tmp) > 0 {
		num, _ = strconv.ParseInt(string(tmp), 10, 64)
	}

	return isLetter, num, trailingZero, string(tmp), offset
}
