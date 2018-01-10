package main

func candles(candlesNumber int, makeNew int) int {
	return candlesNumber + makeNewCandles(candlesNumber/makeNew, candlesNumber%makeNew, makeNew)
}

func makeNewCandles(candlesNumber int, remaining int, makeNew int) int {
	if (candlesNumber+remaining)/makeNew == 0 {
		return candlesNumber
	}
	return candlesNumber + makeNewCandles((candlesNumber+remaining)/makeNew, (candlesNumber+remaining)%makeNew, makeNew)
}
