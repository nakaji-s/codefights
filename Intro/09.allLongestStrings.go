package main

func allLongestStrings(inputArray []string) []string {
	maxLen := 0
	longest := []string{}
	for _, str := range inputArray {
		if maxLen == len(str) {
			longest = append(longest, str)
		} else if maxLen < len(str) {
			longest = []string{str}
			maxLen = len(str)
		}
	}
	return longest
}
