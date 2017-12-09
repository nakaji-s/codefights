package main

func depositProfit(deposit int, rate int, threshold int) int {
	n := float64(deposit)
	for i := 0; ; i++ {
		if n >= float64(threshold) {
			return i
		}
		n = float64(n) * float64(100+rate) / float64(100)
	}
}
