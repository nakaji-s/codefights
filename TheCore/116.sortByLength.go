package main

import "sort"

func sortByLength(inputArray []string) []string {
	sort.Slice(inputArray, func(i, j int) bool {
		return len(inputArray[i]) < len(inputArray[j])
	})

	return inputArray
}
