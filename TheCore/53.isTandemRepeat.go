package main

func isTandemRepeat(inputString string) bool {
	s1 := inputString[0 : len(inputString)/2]
	s2 := inputString[len(inputString)/2 : len(inputString)]
	return s1 == s2
}
