package main

import "math"

func bishopAndPawn(bishop string, pawn string) bool {
	x := int(pawn[0]) - int(bishop[0])
	y := int(pawn[1]) - int(bishop[1])
	return math.Abs(float64(x)) == math.Abs(float64(y))
}
