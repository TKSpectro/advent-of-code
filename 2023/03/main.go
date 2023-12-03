package main

import (
	_ "embed"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func init() {
	// nice idea from https://github.com/alexchao26/advent-of-code-go
	// do this in init (not main) so test file has same input
	input = cleanInput(input)
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func cleanInput(input string) string {
	return strings.TrimRight(input, "\n")
}

func main() {
	println(part1(input))
}

var SYMBOLS_RUNE = [...]rune{'/', '@', '*', '%', '$', '+', '-', '=', '&', '#', '!'}

const SYMBOL_KEY = 'S'
const EMPTY_KEY = '.'

func part1(input string) int {
	result := 0

	matrix := convertInputToMatrix(input)

	number := ""
	isAllowed := false

	for rowIdx, row := range matrix {
		for colIdx, char := range row {
			if unicode.IsDigit(char) {
				number += string(char)

				if !isAllowed {
					isAllowed = checkAdjacentForSymbol(matrix, char, rowIdx, colIdx)
				}
			} else {
				if len(number) > 0 && isAllowed {
					tmp, _ := strconv.Atoi(string(number))
					result += tmp
				}

				number = ""
				isAllowed = false

			}
		}

		if len(number) > 0 && isAllowed {
			tmp, _ := strconv.Atoi(string(number))
			result += tmp
		}

		number = ""
		isAllowed = false
	}

	// printMatrix(matrix)

	return result
}

func part2(input string) int {
	result := 0

	return result
}

func checkAdjacentForSymbol(matrix [][]rune, char rune, rowIdx int, colIdx int) bool {
	isAllowed := false

	// check left
	if colIdx > 0 {
		if matrix[rowIdx][colIdx-1] == SYMBOL_KEY {
			isAllowed = true
		}
	}

	// check right
	if colIdx < len(matrix[rowIdx])-1 {
		if matrix[rowIdx][colIdx+1] == SYMBOL_KEY {
			isAllowed = true
		}
	}

	// check top
	if rowIdx > 0 {
		if matrix[rowIdx-1][colIdx] == SYMBOL_KEY {
			isAllowed = true
		}
	}

	// check bottom
	if rowIdx < len(matrix)-1 {
		if matrix[rowIdx+1][colIdx] == SYMBOL_KEY {
			isAllowed = true
		}
	}

	// check top left
	if rowIdx > 0 && colIdx > 0 {
		if matrix[rowIdx-1][colIdx-1] == SYMBOL_KEY {
			isAllowed = true
		}
	}

	// check top right
	if rowIdx > 0 && colIdx < len(matrix[rowIdx])-1 {
		if matrix[rowIdx-1][colIdx+1] == SYMBOL_KEY {
			isAllowed = true
		}
	}

	// check bottom left
	if rowIdx < len(matrix)-1 && colIdx > 0 {
		if matrix[rowIdx+1][colIdx-1] == SYMBOL_KEY {
			isAllowed = true
		}
	}

	// check bottom right
	if rowIdx < len(matrix)-1 && colIdx < len(matrix[rowIdx])-1 {
		if matrix[rowIdx+1][colIdx+1] == SYMBOL_KEY {
			isAllowed = true
		}
	}

	return isAllowed
}

func convertInputToMatrix(input string) [][]rune {
	matrix := make([][]rune, 0)

	for _, line := range strings.Split(input, "\n") {
		row := make([]rune, 0)
		for _, char := range line {
			if checkIfIsSymbol(char) && char != EMPTY_KEY {
				char = SYMBOL_KEY
			}

			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func checkIfIsSymbol(char rune) bool {
	for _, symbol := range SYMBOLS_RUNE {
		if char == symbol {
			return true
		}
	}

	return false
}

func printMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, char := range row {
			print(string(char))
		}
		println()
	}
}
