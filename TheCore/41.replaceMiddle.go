package main

func replaceMiddle(arr []int) []int {
	if len(arr)%2 == 0 {
		tmp := arr[len(arr)/2-1] + arr[len(arr)/2]
		ret := append(arr[:len(arr)/2-1], arr[len(arr)/2:]...)
		ret[len(arr)/2-1] = tmp
		return ret
	}
	return arr
}
