package main

func createArray(size int) []int {
	ret := make([]int, size)
	for i := range ret {
		ret[i] = 1
	}

	return ret
}
