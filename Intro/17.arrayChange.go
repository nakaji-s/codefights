package main

func arrayChange(inputArray []int) int {
	ret := 0
	for i := 1; i < len(inputArray); i++ {
		if inputArray[i-1] >= inputArray[i] {
			num := 1 + inputArray[i-1] - inputArray[i]
			inputArray[i] += num
			ret += num
		}
	}
	return ret
}
