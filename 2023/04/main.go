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

var BLOCK_LENGTH = 3

func part1(input string) int {
	result := 0
	winningNumberBlockCount := 0
	myNumberBlockCount := 0

	winCount := 0

	for idx, line := range strings.Split(input, "\n") {
		winCount = 0

		tmp := strings.Split(line, ":")

		numbers := strings.Split(tmp[1], "|")

		if idx == 0 {
			winningNumberBlockCount = len(numbers[0]) / BLOCK_LENGTH
			myNumberBlockCount = len(numbers[1]) / BLOCK_LENGTH
		}

		winningNumberBlocks := numbers[0]
		myNumberBlocks := numbers[1]

		winningNumbers := make([]int, winningNumberBlockCount)
		myNumbers := make([]int, myNumberBlockCount)

		for i := 0; i < winningNumberBlockCount; i++ {
			tmpNumber, _ := strconv.Atoi(strings.TrimSpace(winningNumberBlocks[i*BLOCK_LENGTH : (i+1)*BLOCK_LENGTH]))
			winningNumbers[i] = tmpNumber
		}

		for i := 0; i < myNumberBlockCount; i++ {
			tmpNumber, _ := strconv.Atoi(strings.TrimSpace(myNumberBlocks[i*BLOCK_LENGTH : (i+1)*BLOCK_LENGTH]))
			myNumbers[i] = tmpNumber
		}

		for _, myNumber := range myNumbers {
			for _, winningNumber := range winningNumbers {
				if myNumber == winningNumber {
					winCount++
				}
			}
		}

		cardWin := 0

		if winCount > 0 {
			cardWin = 1
			for i := 0; i < winCount-1; i++ {
				cardWin *= 2
			}

		}

		result += cardWin
		println("Card", idx+1, "won", winCount, "times\t, card win:", cardWin, "\t, total scratch cards:", result)
	}

	return result
}

func part2(input string) int {
	scratchCardCount := 0
	cardWinCount := 0

	WINNING_NUMBER_BLOCK_COUNT := 0
	PLAYER_NUMBER_BLOCK_COUNT := 0

	lines := strings.Split(input, "\n")
	lineCount := len(lines)

	// Keep track of many copies we have of each card
	// If we win once on a card, we add copy of the next card
	// If we win twice on a card, we add a copy for the next two cards and so on
	copyCounter := make([]int, lineCount)
	for i := range copyCounter {
		copyCounter[i] = 1
	}

	// Do some splitting in a one-liner to get the lengths of the winning and player numbers
	preCalcNumbers := strings.Split(strings.Split(lines[0], ":")[1], "|")

	WINNING_NUMBER_BLOCK_COUNT = len(preCalcNumbers[0]) / BLOCK_LENGTH
	PLAYER_NUMBER_BLOCK_COUNT = len(preCalcNumbers[1]) / BLOCK_LENGTH

	// Create the winning numbers and my numbers arrays once to avoid creating
	// them in each loop and just fully overwrite them each time
	winningNumbers := make([]int, WINNING_NUMBER_BLOCK_COUNT)
	playerNumbers := make([]int, PLAYER_NUMBER_BLOCK_COUNT)

	for idx, line := range lines {
		cardWinCount = 0

		// Card1: 12 23 | 67 89 -> [" 12 23", " 67 89"]
		// numbers[0] -> winning numbers
		// numbers[1] -> player numbers
		numbers := strings.Split(strings.Split(line, ":")[1], "|")

		// one number block is 3 characters long, so we can just split the string
		// into blocks and then convert them to ints (trimming whitespace)
		for i := 0; i < WINNING_NUMBER_BLOCK_COUNT; i++ {
			winningNumbers[i] = blockToInt(numbers[0], i)
		}

		for i := 0; i < PLAYER_NUMBER_BLOCK_COUNT; i++ {
			playerNumbers[i] = blockToInt(numbers[1], i)
		}

		// Count the number of winning numbers that are also in the player numbers
		for _, playerNumber := range playerNumbers {
			for _, winningNumber := range winningNumbers {
				if playerNumber == winningNumber {
					cardWinCount++
				}
			}
		}

		// If we have won on this card, add the number of copies we have to the next card
		// The moment we hit idx, we have won on all previous cards and have all the copies that are possible
		for i := 1; i <= cardWinCount; i++ {
			copyCounter[idx+i] += copyCounter[idx]
		}

		// Just add the number of scratch cards we have to the total
		scratchCardCount += copyCounter[idx]
		println("Card", idx+1, "won", cardWinCount, "times\t, copies:", copyCounter[idx], "\t, total scratch cards:", scratchCardCount)
	}

	return scratchCardCount
}

func blockToInt(str string, blockIdx int) int {
	number, _ := strconv.Atoi(strings.TrimSpace(str[blockIdx*BLOCK_LENGTH : (blockIdx+1)*BLOCK_LENGTH]))
	return number
}
