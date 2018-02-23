package main

func killKthBit(n int, k int) int {
	return n &^ (1 << uint(k-1))
}
