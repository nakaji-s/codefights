package main

func differentRightmostBit(n int, m int) int {
	return -^((^(n ^ m)) ^ ((^(n ^ m)) + 1)) / 2
}
