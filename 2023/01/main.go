package main

import (
	_ "embed"
	"math"
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

func part1(input string) int {
	result := 0

	// run through all lines of input
	for _, line := range strings.Split(input, "\n") {
		first := ""
		last := ""

		// find the first number
		for _, char := range line {
			if unicode.IsDigit(char) {
				first = string(char)
				break
			}
		}

		// find the last number
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = string(line[i])
				break
			}
		}

		// concat first and last number and convert to int
		number, _ := strconv.Atoi(first + last)

		// add to result
		result += number
	}

	return result
}

var spelledNumbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func part2(input string) int {
	result := 0

	// run through all lines of input
	for _, line := range strings.Split(input, "\n") {
		first := ""
		firstNumberIndex := math.MaxInt64
		last := ""
		lastNumberIndex := math.MinInt64

		// find the first number
		for i, char := range line {
			if unicode.IsDigit(char) {
				first = string(char)
				firstNumberIndex = i
				break
			}
		}

		// find the last number
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = string(line[i])
				lastNumberIndex = i
				break
			}
		}

		// check if the first number is spelled out, don't break because we run through all numbers and check by index
		for word, number := range spelledNumbers {
			firstNumberIndexTemp := strings.Index(line, word)

			if firstNumberIndexTemp >= 0 && firstNumberIndexTemp < firstNumberIndex {
				first = number
				firstNumberIndex = firstNumberIndexTemp
			}
		}

		// check if the last number is spelled out, don't break because we run through all numbers and check by index
		for word, number := range spelledNumbers {
			lastNumberIndexTemp := strings.LastIndex(line, word)

			if lastNumberIndexTemp >= 0 && lastNumberIndexTemp > lastNumberIndex {
				last = number
				lastNumberIndex = lastNumberIndexTemp
			}
		}

		// concat first and last number and convert to int
		number, _ := strconv.Atoi(first + last)

		// add to result
		result += number
	}

	return result
}
