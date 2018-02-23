package main

func equalPairOfBits(n int, m int) int {
	return -^((n ^ m) ^ ((n ^ m) + 1)) / 2
}
