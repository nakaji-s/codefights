package main

import "strings"

func isUnstablePair(filename1 string, filename2 string) bool {
	ret1 := isFilename1Long(filename1, filename2)
	ret2 := isFilename1Long(strings.ToUpper(filename1), strings.ToUpper(filename2))
	return ret1 != ret2
}

func isFilename1Long(filename1 string, filename2 string) bool {
	for i, _ := range filename1 {
		if len(filename2) <= i {
			return true
		}

		if filename1[i] > filename2[i] {
			return true
		} else if filename1[i] < filename2[i] {
			return false
		}
	}

	return false
}
