package main

import (
	_ "embed"
	"fmt"
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
	println(part2(input))
}

var PART_1_SYMBOLS = []rune{'/', '@', '*', '%', '$', '+', '-', '=', '&', '#', '!'}
var PART_2_SYMBOLS = []rune{'/', '@', '%', '$', '+', '-', '=', '&', '#', '!'} // removed '*' because thats the gear

const SYMBOL_KEY = 'S'
const EMPTY_KEY = '.'
const GEAR_KEY = '*'

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

					isAllowed, _ = checkAdjacentForSymbol(matrix, char, SYMBOL_KEY, rowIdx, colIdx)
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

// part2
func part2(input string) int {
	result := 0

	matrix := convertInputToMatrixPart2(input)

	number := ""

	// Build up a map of gears with the position as key and both connections as value
	// Example: map["1/3":[467, 35], "8/5":[755, 598]]
	gears := make(map[string][2]int)

	for rowIdx, row := range matrix {
		currentGearPos := [2]int{-1, -1}
		for colIdx, char := range row {
			if unicode.IsDigit(char) {
				number += string(char)

				hasGear, gearPos := checkAdjacentForSymbol(matrix, char, GEAR_KEY, rowIdx, colIdx)
				if hasGear && currentGearPos[0] == -1 && currentGearPos[1] == -1 {
					currentGearPos[0] = gearPos[0]
					currentGearPos[1] = gearPos[1]
				}
			} else {
				if len(number) > 0 && currentGearPos[0] != -1 && currentGearPos[1] != -1 {
					tmp, _ := strconv.Atoi(string(number))

					gear, found := gears[getKey(currentGearPos)]
					if found {
						gear[1] = tmp
						gears[getKey(currentGearPos)] = gear
					} else {
						gears[getKey(currentGearPos)] = [2]int{tmp, -1}
					}
				}

				currentGearPos = [2]int{-1, -1}
				number = ""
			}
		}

		if len(number) > 0 && currentGearPos[0] != -1 && currentGearPos[1] != -1 {
			tmp, _ := strconv.Atoi(string(number))

			gear, found := gears[getKey(currentGearPos)]
			if found {
				gear[1] = tmp
				gears[getKey(currentGearPos)] = gear
			} else {
				gears[getKey(currentGearPos)] = [2]int{-1, tmp}
			}
		}

		currentGearPos = [2]int{-1, -1}
		number = ""
	}

	for _, key := range gears {
		if key[0] != -1 && key[1] != -1 {
			result += key[0] * key[1]
		}
	}

	// printGears(gears)

	return result
}

func printGears(gears map[string][2]int) {
	for key, value := range gears {
		println(key, value[0], value[1])
	}
}

func getKey(pos [2]int) string {
	return (fmt.Sprint(pos[0]) + "/" + fmt.Sprint(pos[1]))
}

func checkAdjacentForSymbol(matrix [][]rune, char rune, symbol rune, rowIdx int, colIdx int) (found bool, symbolPosition [2]int) {
	isAllowed := false

	symbolPos := [...]int{-1, -1}

	// check left
	if colIdx > 0 {
		if matrix[rowIdx][colIdx-1] == symbol {
			isAllowed = true
			symbolPos[0] = rowIdx
			symbolPos[1] = colIdx - 1
		}
	}

	// check right
	if colIdx < len(matrix[rowIdx])-1 {
		if matrix[rowIdx][colIdx+1] == symbol {
			isAllowed = true
			symbolPos[0] = rowIdx
			symbolPos[1] = colIdx + 1
		}
	}

	// check top
	if rowIdx > 0 {
		if matrix[rowIdx-1][colIdx] == symbol {
			isAllowed = true
			symbolPos[0] = rowIdx - 1
			symbolPos[1] = colIdx
		}
	}

	// check bottom
	if rowIdx < len(matrix)-1 {
		if matrix[rowIdx+1][colIdx] == symbol {
			isAllowed = true
			symbolPos[0] = rowIdx + 1
			symbolPos[1] = colIdx
		}
	}

	// check top left
	if rowIdx > 0 && colIdx > 0 {
		if matrix[rowIdx-1][colIdx-1] == symbol {
			isAllowed = true
			symbolPos[0] = rowIdx - 1
			symbolPos[1] = colIdx - 1
		}
	}

	// check top right
	if rowIdx > 0 && colIdx < len(matrix[rowIdx])-1 {
		if matrix[rowIdx-1][colIdx+1] == symbol {
			isAllowed = true
			symbolPos[0] = rowIdx - 1
			symbolPos[1] = colIdx + 1
		}
	}

	// check bottom left
	if rowIdx < len(matrix)-1 && colIdx > 0 {
		if matrix[rowIdx+1][colIdx-1] == symbol {
			isAllowed = true
			symbolPos[0] = rowIdx + 1
			symbolPos[1] = colIdx - 1
		}
	}

	// check bottom right
	if rowIdx < len(matrix)-1 && colIdx < len(matrix[rowIdx])-1 {
		if matrix[rowIdx+1][colIdx+1] == symbol {
			isAllowed = true
			symbolPos[0] = rowIdx + 1
			symbolPos[1] = colIdx + 1
		}
	}

	return isAllowed, symbolPos
}

func convertInputToMatrix(input string) [][]rune {
	matrix := make([][]rune, 0)

	for _, line := range strings.Split(input, "\n") {
		row := make([]rune, 0)
		for _, char := range line {
			if checkIfIsSymbol(char, PART_1_SYMBOLS) && char != EMPTY_KEY {
				char = SYMBOL_KEY
			}

			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func convertInputToMatrixPart2(input string) [][]rune {
	matrix := make([][]rune, 0)

	for _, line := range strings.Split(input, "\n") {
		row := make([]rune, 0)
		for _, char := range line {
			if checkIfIsSymbol(char, PART_2_SYMBOLS) && char != EMPTY_KEY && char != GEAR_KEY {
				char = SYMBOL_KEY
			}

			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func checkIfIsSymbol(char rune, symbols []rune) bool {
	for _, symbol := range symbols {
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
