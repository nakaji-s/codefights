package main

func minimalNumberOfCoins(coins []int, price int) int {
	sum := 0
	count := 0
	for i := len(coins) - 1; i >= 0; i-- {
		for {
			if sum+coins[i] <= price {
				sum += coins[i]
				count++
			} else {
				break
			}
		}
	}

	return count
}
