package main

func alternatingSums(a []int) []int {
	retA := 0
	retB := 0
	for i, n := range a {
		if i%2 == 0 {
			retA += n
		} else {
			retB += n
		}
	}
	return []int{retA, retB}
}
