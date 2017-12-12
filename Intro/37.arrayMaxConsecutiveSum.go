package main

func arrayMaxConsecutiveSum(inputArray []int, k int) int {
	max := 0
	for i := 0; i < len(inputArray)+1-k; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += inputArray[j]
		}
		if max < sum {
			max = sum
		}
	}

	return max
}
