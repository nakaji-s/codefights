package main

func firstReverseTry(arr []int) []int {
	if len(arr) > 0 {
		arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]
	}
	return arr
}
