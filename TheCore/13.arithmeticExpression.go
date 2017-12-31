package main

func arithmeticExpression(a int, b int, c int) bool {
	if a+b == c {
		return true
	}
	if a-b == c {
		return true
	}
	if a*b == c {
		return true
	}
	if float64(a)/float64(b) == float64(c) {
		return true
	}
	return false
}
