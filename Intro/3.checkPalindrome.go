package main

func checkPalindrome(inputString string) bool {
	for i := 0; i < len(inputString)/2; i++ {
		if inputString[i] != inputString[len(inputString)-1-i] {
			return false
		}
	}
	return true
}
