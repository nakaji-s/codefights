package main

import "strings"
import (
	"fmt"
	"unicode"
)

func chessNotation(notation string) string {
	field := make([][]string, 8)
	for i := 0; i < 8; i++ {
		field[i] = make([]string, 8)
	}

	rotateField := make([][]string, 8)
	for i := 0; i < 8; i++ {
		rotateField[i] = make([]string, 8)
	}

	lines := strings.Split(notation, "/")
	for i, line := range lines {
		offset := 0
		for j, character := range line {
			if unicode.IsDigit(character) {
				digit := int(character - '0')
				offset += digit - 1
			} else {
				field[i][j+offset] = string(character)
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			rotateField[i][j] = field[7-j][i]
		}
	}

	rotateLines := []string{}
	for _, line := range rotateField {
		lineString := ""
		emptyCount := 0
		for _, character := range line {
			if character == "" {
				emptyCount++
			} else {
				if emptyCount > 0 {
					lineString = lineString + fmt.Sprint(emptyCount)
					emptyCount = 0
				}
				lineString = lineString + character
			}
		}

		if emptyCount > 0 {
			lineString = lineString + fmt.Sprint(emptyCount)
		}
		rotateLines = append(rotateLines, lineString)
	}

	return strings.Join(rotateLines, "/")
}
