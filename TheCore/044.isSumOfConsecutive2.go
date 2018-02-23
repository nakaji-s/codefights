package main

func isSumOfConsecutive2(n int) int {
	m := map[int]int{}

	for i := 1; i < 14; i++ {
		sum := i
		for j := i + 1; j < 13; j++ {
			sum += j
			if sum > 25 {
				break
			}

			m[sum]++
		}
	}

	return m[n]
}
