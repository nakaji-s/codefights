package main

func threeSplit(a []int) int {
	sum := 0
	for _, num := range a {
		sum += num
	}
	average := sum / 3

	return createChunks(a, average, 0, 3)
}

func createChunks(a []int, average, offset, remaining int) int {
	count := 0
	sum := 0
	for i := offset; i < len(a); i++ {
		sum += a[i]
		if sum == average {
			if remaining == 1 {
				if i == len(a)-1 {
					return 1
				}
			} else {
				count += createChunks(a, average, i+1, remaining-1)
			}
		}
	}

	return count
}
