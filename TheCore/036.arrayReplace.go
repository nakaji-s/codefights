package main

func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for i, _ := range inputArray {
		if inputArray[i] == elemToReplace {
			inputArray[i] = substitutionElem
		}
	}

	return inputArray
}
