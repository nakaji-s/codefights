package main

import "sort"

func maximumSum(a []int, q [][]int) int {
	queryCountMap := map[int]int{}

	for _, numRange := range q {
		for i := numRange[0]; i <= numRange[1]; i++ {
			queryCountMap[i]++
		}
	}

	targetArray := [][]int{}
	for i, num := range queryCountMap {
		targetArray = append(targetArray, []int{i, num})
	}
	sort.Slice(targetArray, func(i, j int) bool {
		return targetArray[i][1] > targetArray[j][1]
	})

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	rearrangedA := make([]int, len(a))
	for i, num := range a {
		if len(targetArray) > i {
			rearrangedA[targetArray[i][0]] = num
		}
	}

	sum := 0
	for _, numRange := range q {
		for i := numRange[0]; i <= numRange[1]; i++ {
			sum += rearrangedA[i]
		}
	}

	return sum
}
