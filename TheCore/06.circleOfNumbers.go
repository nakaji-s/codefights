package main

func circleOfNumbers(n int, firstNumber int) int {
	return (firstNumber + n/2) % n
}
