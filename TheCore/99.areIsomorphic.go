package main

func areIsomorphic(array1 [][]int, array2 [][]int) bool {
	if len(array1) == len(array2) {
		for i, _ := range array1 {
			if len(array1[i]) != len(array2[i]) {
				return false
			}
		}
		return true
	}
	return false
}
