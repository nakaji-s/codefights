package main

func isInfiniteProcess(a int, b int) bool {
	return !(b >= a && (b-a)%2 == 0)
}
