package main

func arrayConversion(inputArray []int) int {
	for i := 0; len(inputArray) > 1; i++ {
		if i%2 == 0 {
			inputArray = mergeArray(inputArray)
		} else {
			inputArray = mergeArrayProduct(inputArray)
		}
	}
	return inputArray[0]
}

func mergeArray(inputArray []int) []int {
	for i := len(inputArray) - 1; i >= 1; i -= 2 {
		inputArray[i-1] = inputArray[i-1] + inputArray[i]
		inputArray = append(inputArray[:i], inputArray[i+1:]...)
	}

	return inputArray
}

func mergeArrayProduct(inputArray []int) []int {
	for i := len(inputArray) - 1; i >= 1; i -= 2 {
		inputArray[i-1] = inputArray[i-1] * inputArray[i]
		inputArray = append(inputArray[:i], inputArray[i+1:]...)
	}

	return inputArray
}
