package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Location struct {
	x      int
	y      int
	symbol rune
}

type PartNumber struct {
	start     int
	end       int
	number    int
	confirmed bool
}

func markNumbers(pns []PartNumber, x int) {
	for i, num := range pns {
		if !((x+1 < num.start) || (num.end <= x-1)) {
			pns[i].confirmed = true
		}
	}
}

func matchNumbers(pns []PartNumber, x int) []PartNumber {
	matched := make([]PartNumber, 0)
	for i, num := range pns {
		if !((x+1 < num.start) || (num.end <= x-1)) {
			matched = append(matched, pns[i])
		}
	}

	return matched
}

func readInput(inputFilePath string) ([][]PartNumber, []Location, int) {
	// For each symbol location, check the surroundings for a number which touches it
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	defer file.Close()

	// Build array of rows, where each entry is a number
	// Numbers contain their start/end
	// Also track the symbol locations
	partNumbers := make([][]PartNumber, 0)
	locations := make([]Location, 0)
	scanner := bufio.NewScanner(file)
	y := 0
	lineLength := 0
	for scanner.Scan() {
		partNumbers = append(partNumbers, make([]PartNumber, 0))

		wasDigit := false
		pnum := PartNumber{confirmed: false}
		line := scanner.Text()
		for x, ch := range line {
			// Handle digits
			if unicode.IsDigit(ch) {
				val, _ := strconv.Atoi(string(ch))
				if wasDigit {
					pnum.number = pnum.number*10 + val
				} else {
					pnum.number = val
					pnum.start = x
				}
				wasDigit = true
			} else {
				// We had a number and we finished it
				if wasDigit {
					pnum.end = x
					partNumbers[y] = append(partNumbers[y], pnum)
					pnum = PartNumber{confirmed: false}
				}
				if ch != '.' {
					// Save symbol locations
					locations = append(locations, Location{x: x, y: y, symbol: ch})
				}

				wasDigit = false
			}
		}

		// Handle number ending the line
		if wasDigit {
			pnum.end = len(line)
			partNumbers[y] = append(partNumbers[y], pnum)
		}
		lineLength = len(line)
		y++
	}
	return partNumbers, locations, lineLength
}

func prettyPrint(pnums [][]PartNumber, locs []Location, lineLength int) {
	currLoc := 0
	for y := 0; y < len(pnums); y++ {
		currNum := 0
		for x := 0; x < lineLength; {
			// fmt.Printf("\n%d %d\n", x, y)
			if currLoc < len(locs) && locs[currLoc].x == x && locs[currLoc].y == y {
				fmt.Print("*")
				currLoc++
				x++
			} else if len(pnums[y]) > 0 && currNum < len(pnums[y]) && pnums[y][currNum].start == x {
				if pnums[y][currNum].confirmed {
					fmt.Printf("\033[32m%d\033[0m", pnums[y][currNum].number)
					x += pnums[y][currNum].end - pnums[y][currNum].start
					// print(" ")
					// x++
				} else {
					fmt.Printf("\033[31m%d\033[0m", pnums[y][currNum].number)
					x += pnums[y][currNum].end - pnums[y][currNum].start
				}
				currNum++
			} else {
				fmt.Print(".")
				x++
			}
		}
		fmt.Printf("\n")
	}
}

func part1(inputFilePath string) {
	partNumbers, locations, _ := readInput(inputFilePath)

	for _, location := range locations {
		markNumbers(partNumbers[location.y-1], location.x)
		markNumbers(partNumbers[location.y], location.x)
		markNumbers(partNumbers[location.y+1], location.x)
	}

	total := 0
	for _, row := range partNumbers {
		for _, pnum := range row {
			if pnum.confirmed {
				total += pnum.number
			}
		}
	}
	println(total)
}

func part2(path string) {
	partNumbers, locations, _ := readInput(path)

	total := 0
	for _, location := range locations {
		if location.symbol != '*' {
			continue
		}

		matched := make([]PartNumber, 0)
		matched = append(matched, matchNumbers(partNumbers[location.y-1], location.x)...)
		matched = append(matched, matchNumbers(partNumbers[location.y], location.x)...)
		matched = append(matched, matchNumbers(partNumbers[location.y+1], location.x)...)

		if len(matched) == 2 {
			total += matched[0].number * matched[1].number
		}
	}
	fmt.Println(total)
}

func main() {
	// Get input
	if len(os.Args) != 3 {
		println(len(os.Args))
		println("Usage: [INPUT FILE] [PART 1 or 2]")
	} else {
		inputFilePath := os.Args[1]
		part := os.Args[2]
		if part == "1" {
			part1(inputFilePath)
		} else {
			part2(inputFilePath)
		}
	}
}
