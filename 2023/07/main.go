package main

import (
	_ "embed"
	"fmt"
	"sort"
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
	println("PART 1:", part1(input))
	println("PART 2:", part2(input))
}

type Card int

func ParseCard(input rune, part2 bool) int {
	switch input {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if part2 {
			return 1
		}
		return 11
	case 'T':
		return 10
	default:
		return int(input) - 48
	}
}

type Hand struct {
	Cards [5]int
	Bid   int
	Type  int
	Orig  string
}

func (h Hand) String() string {
	cardStrings := ""
	for _, card := range h.Cards {
		cardStrings += fmt.Sprintf("%v ", card)
	}
	return fmt.Sprintf("%v %v %v %v", h.Orig, cardStrings, h.Bid, h.Type)
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (h *Hand) SetType(part2 bool) {
	mapping := map[int]int{}
	highestCount := 0

	for _, card := range h.Cards {
		mapping[card]++
		highestCount = max(highestCount, mapping[card])
	}
	jokerCount := mapping[1]

	switch len(mapping) {
	case 1:
		h.Type = FiveOfAKind
	case 2:
		if part2 && jokerCount > 0 {
			h.Type = FiveOfAKind
			return
		}
		if highestCount == 4 {
			h.Type = FourOfAKind
			return
		}
		h.Type = FullHouse
	case 3:
		if highestCount == 3 {
			if part2 && jokerCount > 0 {
				h.Type = FourOfAKind
				return
			}
			h.Type = ThreeOfAKind
			return
		}
		if part2 && jokerCount > 1 {
			h.Type = FourOfAKind
			return
		} else if part2 && jokerCount > 0 {
			h.Type = FullHouse
			return
		}
		h.Type = TwoPair
	case 4:
		if part2 && jokerCount > 0 {
			h.Type = ThreeOfAKind
			return
		}
		h.Type = OnePair
	case 5:
		if part2 && jokerCount > 0 {
			h.Type = OnePair
			return
		}
		h.Type = HighCard
	default:
		h.Type = HighCard
	}
}

func part1(input string) int {
	return part(input, false)
}

func part2(input string) int {

	return part(input, true)
}

func part(input string, part2 bool) int {
	result := 0

	hands := parseInput(input, part2)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			for idx := 0; idx < 5; idx++ {
				if hands[i].Cards[idx] == hands[j].Cards[idx] {
					continue
				}
				return hands[i].Cards[idx] < hands[j].Cards[idx]
			}
		}
		return hands[i].Type < hands[j].Type
	})

	for rank, hand := range hands {
		result += hand.Bid * (rank + 1)
	}

	return result
}

func parseInput(input string, part2 bool) []Hand {
	lines := strings.Split(input, "\n")
	hands := make([]Hand, len(lines))

	for idx, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		hand := Hand{}

		// first 5 runes of each line are the cards
		for i, str := range split[0] {
			if i > 4 {
				break
			}
			hand.Cards[i] = ParseCard(str, part2)
		}
		hand.SetType(part2)
		hand.Bid, _ = strconv.Atoi(string(split[1]))
		hands[idx] = hand
	}

	return hands
}
