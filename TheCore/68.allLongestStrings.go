package main

func allLongestStrings(inputArray []string) []string {
	maxLen := 0
	ret := []string{}
	for _, str := range inputArray {
		if maxLen < len(str) {
			ret = []string{str}
			maxLen = len(str)
		} else if maxLen == len(str) {
			ret = append(ret, str)
		}
	}

	return ret
}
