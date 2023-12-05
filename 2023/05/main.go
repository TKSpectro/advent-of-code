package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed input.txt
var input string

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func init() {
	// nice idea from https://github.com/alexchao26/advent-of-code-go
	// do this in init (not main) so test file has same input
	input = cleanInput(input)
	if len(input) == 0 {
		panic("empty input.txt file")
	}

	// set GOMAXPROCS to 3/4 the number of CPUs
	runtime.GOMAXPROCS(runtime.NumCPU() * 3 / 4)
	println("\n\nUsing:", runtime.NumCPU(), "cores", "\n\n")
}

func cleanInput(input string) string {
	return strings.TrimRight(input, "\n")
}

func main() {
	println(part2(input))
}

type Mapping struct {
	destStart int
	srcStart  int
	length    int
}

func (m Mapping) String() string {
	return fmt.Sprintf("destStart: %v, srcStart: %v, length: %v", m.destStart, m.srcStart, m.length)
}

var MAPPINGS = map[string]int{
	"seed-to-soil":            0,
	"soil-to-fertilizer":      1,
	"fertilizer-to-water":     2,
	"water-to-light":          3,
	"light-to-temperature":    4,
	"temperature-to-humidity": 5,
	"humidity-to-location":    6,
}

var R_MAPPINGS = map[int]string{
	0: "seed-to-soil",
	1: "soil-to-fertilizer",
	2: "fertilizer-to-water",
	3: "water-to-light",
	4: "light-to-temperature",
	5: "temperature-to-humidity",
	6: "humidity-to-location",
}

func NewMappingList() [][]Mapping {
	mappings := make([][]Mapping, 0)
	for i := 0; i < len(MAPPINGS); i++ {
		mappings = append(mappings, []Mapping{})
	}
	return mappings
}

func part1(input string) int {
	lowestLocationNumber := math.MaxInt

	seeds := make([]int, 0)
	mappings := NewMappingList()

	parseMappings(input, &seeds, &mappings)

	for _, seed := range seeds {
		num := seed
		for i := 0; i < len(MAPPINGS); i++ {
			num = getNumberForNextStage(num, &mappings[i])
		}

		if num < lowestLocationNumber {
			lowestLocationNumber = num
		}
	}

	return lowestLocationNumber
}

// part2 uses goroutines to speed up the calculation as we have some pretty big ranges
// go from ~2m30s to ~40s on 8 cores on my machine (XPS 15 9500, Intel i7-10750H (12) @ 5.000GHz)
// also i just wanted to try out goroutines
func part2(input string) int {
	lowestLocationNumber := math.MaxInt

	seeds := make([]int, 0)
	mappings := NewMappingList()

	parseMappings(input, &seeds, &mappings)

	coroutineCount := len(seeds) / 2
	startTime := time.Now()

	wg.Add(coroutineCount)

	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		// rng = range
		rng := seeds[i+1]

		println("seed: ", seed, "started", "rng: ", rng)

		go getLowestLocationNumber(seed, rng, &mappings, &lowestLocationNumber)
	}

	wg.Wait()
	log.Printf("time: %s", time.Since(startTime))

	return lowestLocationNumber
}

func getLowestLocationNumber(seed int, rng int, mappings *[][]Mapping, lowestLocationNumber *int) {
	defer wg.Done()

	for i := seed; i < seed+rng; i++ {
		num := i

		for i := 0; i < len(MAPPINGS); i++ {
			num = getNumberForNextStage(num, &(*mappings)[i])
		}

		if num < *lowestLocationNumber {
			*lowestLocationNumber = num
		}

		// print the progress in the current seed every 10%
		if i%int(rng/10) == 0 {
			println("\tseed: ", seed, "\tprogress: ", int(float64(i-seed)/float64(rng)*100), "%")
		}
	}

	println("seed: ", seed, " done")
}

// printMappings prints the seeds and mappings
func printMappings(seeds *[]int, mappings *[][]Mapping) {
	print("seeds: ")
	for _, seed := range *seeds {
		print(seed, ", ")
	}

	println("")
	for i, mapping := range *mappings {
		println("")
		fmt.Printf("%s\n", R_MAPPINGS[i])
		for _, m := range mapping {
			println(m.String())
		}
	}
}

// parseMappings parses the input string and fills the seeds and mappings
func parseMappings(input string, seeds *[]int, mappings *[][]Mapping) {
	lines := strings.Split(input, "\n")
	currentMappingIdx := -1

	for idx, line := range lines {
		// skip empty lines
		if len(line) == 0 {
			continue
		}

		// first line contains seeds like "seeds: 81 14 16"
		if idx == 0 {
			numbers := strings.Split(strings.Split(line, ": ")[1], " ")
			for _, part := range numbers {
				num, _ := strconv.Atoi(part)

				*seeds = append(*seeds, num)
			}
			continue
		}

		// line contains mapping like "seed-to-soil" so the next lines will contain mappings
		if len(R_MAPPINGS)-1 > currentMappingIdx && strings.HasPrefix(line, R_MAPPINGS[currentMappingIdx+1]) {
			currentMappingIdx++
			continue
		}

		// line contains mapping like "50 100 3"
		numbers := strings.Split(line, " ")
		destStart, _ := strconv.Atoi(numbers[0])
		srcStart, _ := strconv.Atoi(numbers[1])
		length, _ := strconv.Atoi(numbers[2])

		(*mappings)[currentMappingIdx] = append((*mappings)[currentMappingIdx], Mapping{
			destStart: destStart,
			srcStart:  srcStart,
			length:    length,
		})
	}
}

// getNumberForNextStage returns the number corresponding number in the next stage
func getNumberForNextStage(number int, mapping *[]Mapping) int {
	for _, m := range *mapping {
		if number >= m.srcStart && number < m.srcStart+m.length {
			return m.destStart + number - m.srcStart
		}
	}

	return number
}
