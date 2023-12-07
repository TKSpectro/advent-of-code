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
	result := 1

	times := []int{}
	distances := []int{}

	parseInput(input, &times, &distances)

	for i := 0; i < len(times); i++ {
		result *= calculateWinning(times[i], distances[i])
	}

	return result
}

func part2(input string) int {
	time, distance := parseInputPart2(input)

	return calculateWinning(time, distance)
}

// calculateWinning
func calculateWinning(raceDuration int, distanceRecord int) int {
	result := 0

	// ignore the first millisecond and the last millisecond as the boat would move 0 distance
	for holdDuration := 1; holdDuration < raceDuration; holdDuration++ {
		timeToDrive := raceDuration - holdDuration
		distanceTraveled := holdDuration * timeToDrive

		if distanceTraveled > distanceRecord {
			result++
		}
	}

	return result
}

func parseInput(input string, times *[]int, distances *[]int) {
	lines := strings.Split(input, "\n")

	timeStrings := strings.Split(lines[0], " ")

	curNum := 0
	for _, timeString := range timeStrings {
		if len(timeString) == 0 || timeString == "Time:" {
			continue
		}

		curNum, _ = strconv.Atoi(timeString)
		*times = append(*times, curNum)
	}

	distanceStrings := strings.Split(lines[1], " ")
	for _, distanceString := range distanceStrings {
		if len(distanceString) == 0 || distanceString == "Distance:" {
			continue
		}

		curNum, _ = strconv.Atoi(distanceString)
		*distances = append(*distances, curNum)
	}
}

func parseInputPart2(input string) (time int, distance int) {
	lines := strings.Split(input, "\n")

	timeString := strings.Split(lines[0], ":")[1]
	timeString = strings.Replace(timeString, " ", "", -1)

	durationString := strings.Split(lines[1], ":")[1]
	durationString = strings.Replace(durationString, " ", "", -1)

	time, _ = strconv.Atoi(timeString)
	distance, _ = strconv.Atoi(durationString)

	return time, distance
}
