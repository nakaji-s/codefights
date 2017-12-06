package main

func evenDigitsOnly(n int) bool {
	for n >= 10 {
		if n%2 == 1 {
			return false
		}
		n = n / 10
	}
	return (n%2 == 0)
}
