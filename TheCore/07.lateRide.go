package main

func lateRide(n int) int {
	return n/60/10 + n/60%10 + n%60/10 + n%60%10
}
