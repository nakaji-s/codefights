package main

func alphabeticShift(inputString string) string {
	var replacedString string
	for _, rune := range inputString {
		replacedString += string((rune+1-'a')%('z'+1-'a') + 'a')
	}
	return replacedString
}
