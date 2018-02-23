package main

func isSmooth(arr []int) bool {
	if arr[0] != arr[len(arr)-1] {
		return false
	}
	if len(arr)%2 == 0 {
		return arr[0] == arr[len(arr)/2-1]+arr[len(arr)/2]
	} else {
		return arr[0] == arr[len(arr)/2]
	}
}
