package main

func concatenateArrays(a []int, b []int) []int {
	return append(a, b...)
}
