package main

import (
	"reflect"
	"strconv"
)

func correctNonogram(size int, nonogramField [][]string) bool {
	definitionSize := (size + 1) / 2

	// check row
	for i := definitionSize; i < len(nonogramField); i++ {
		squareSizes := []int{}
		for j := 0; j < definitionSize; j++ {
			if nonogramField[i][j] != "-" {
				num, _ := strconv.Atoi(nonogramField[i][j])
				squareSizes = append(squareSizes, num)
			}
		}

		realSquareSizes := []int{}
		streak := 0
		for j := definitionSize; j < len(nonogramField); j++ {
			if nonogramField[i][j] == "#" {
				streak++
				if j == len(nonogramField)-1 && streak > 0 {
					realSquareSizes = append(realSquareSizes, streak)
				}
			} else {
				if streak > 0 {
					realSquareSizes = append(realSquareSizes, streak)
					streak = 0
				}
			}
		}

		if !reflect.DeepEqual(squareSizes, realSquareSizes) {
			return false
		}
	}

	// check column
	for i := definitionSize; i < len(nonogramField); i++ {
		squareSizes := []int{}
		for j := 0; j < definitionSize; j++ {
			if nonogramField[j][i] != "-" {
				num, _ := strconv.Atoi(nonogramField[j][i])
				squareSizes = append(squareSizes, num)
			}
		}

		realSquareSizes := []int{}
		streak := 0
		for j := definitionSize; j < len(nonogramField); j++ {
			if nonogramField[j][i] == "#" {
				streak++
				if j == len(nonogramField)-1 && streak > 0 {
					realSquareSizes = append(realSquareSizes, streak)
				}
			} else {
				if streak > 0 {
					realSquareSizes = append(realSquareSizes, streak)
					streak = 0
				}
			}
		}

		if !reflect.DeepEqual(squareSizes, realSquareSizes) {
			return false
		}
	}

	return true
}
