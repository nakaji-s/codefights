package main

func houseNumbersSum(inputArray []int) int {
	sum := 0
	for _, num := range inputArray {
		sum += num
		if num == 0 {
			break
		}
	}

	return sum
}
