package main

import (
	_ "embed"
	"strconv"
	"strings"
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
	const ALLOWED_RED = 12
	const ALLOWED_GREEN = 13
	const ALLOWED_BLUE = 14

	result := 0

	for _, line := range strings.Split(input, "\n") {
		line, _ = strings.CutPrefix(line, "Game ")

		// split the game id from the rest of the line
		temp := strings.Split(line, ":")

		gameId, _ := strconv.Atoi(temp[0])
		line = temp[1]

		// split the rest of the line into the individual pulls
		pulls := strings.Split(line, ";")
		allowedPull := true

		for _, pull := range pulls {
			red := 0
			green := 0
			blue := 0

			// split the pull into the individual cubes (3 blue, 2 red, etc)
			cubes := strings.Split(pull, ",")

			for _, cube := range cubes {
				// split the cube into the count and color (and remove whitespace)
				temp := strings.Split(strings.TrimSpace(cube), " ")

				count, _ := strconv.Atoi(temp[0])
				color := temp[1]

				if color == "red" {
					red += count
				} else if color == "green" {
					green += count
				} else if color == "blue" {
					blue += count
				}
			}

			// if any of the colors are over the allowed amount, this pull is not allowed (and we can stop checking the line)
			if red > ALLOWED_RED || green > ALLOWED_GREEN || blue > ALLOWED_BLUE {
				allowedPull = false
				break
			}
		}

		if allowedPull {
			result += gameId
		}
	}

	return result
}

func part2(input string) int {
	result := 0

	for _, line := range strings.Split(input, "\n") {
		// count of the minimum number of cubes of each color needed to make this line/pull possible
		minRed := 0
		minGreen := 0
		minBlue := 0

		// split the game id from the rest of the line
		line, _ = strings.CutPrefix(line, "Game ")
		temp := strings.Split(line, ":")

		// gameId, _ := strconv.Atoi(temp[0])
		line = temp[1]

		// split the rest of the line into the individual pulls
		pulls := strings.Split(line, ";")

		for _, pull := range pulls {
			// split the pull into the individual cubes (3 blue, 2 red, etc)
			temp := strings.Split(pull, ",")

			for _, cube := range temp {
				temp := strings.Split(strings.TrimSpace(cube), " ")

				count, _ := strconv.Atoi(temp[0])
				color := temp[1]

				// if the count is greater than the current min, update the min
				if color == "red" && count > minRed {
					minRed = count
				} else if color == "green" && count > minGreen {
					minGreen = count
				} else if color == "blue" && count > minBlue {
					minBlue = count
				}
			}
		}

		result += minRed * minGreen * minBlue
	}

	return result
}
