package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) [][]int {
	contents, _ := os.ReadFile(path)

	lines := strings.Split(string(contents), "\n")

	histories := make([][]int, 0)
	for i, line := range lines {
		if line == "" {
			continue
		}
		histories = append(histories, make([]int, 0))
		for _, numStr := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(numStr)
			histories[i] = append(histories[i], num)
		}
	}

	return histories
}

func continueSequence(history []int) int {
	levels := make([][]int, 1)
	levels[0] = history

	// Construct the levels
	for depth, nonzero := 1, true; nonzero; depth += 1 {
		// Add a new level
		nonzero = false
		prevLevel := levels[depth-1]
		levels = append(levels, make([]int, 0))
		for i := 0; i < len(prevLevel)-1; i += 1 {
			first, second := prevLevel[i], prevLevel[i+1]
			diff := second - first
			if diff != 0 {
				nonzero = true
			}
			levels[depth] = append(levels[depth], diff)
		}
	}

	// Build the next step
	for depth := len(levels) - 3; depth >= 0; depth-- {
		lower := levels[depth+1][len(levels[depth+1])-1]
		prior := levels[depth][len(levels[depth])-1]
		levels[depth] = append(levels[depth], lower+prior)
	}

	return levels[0][len(levels[0])-1]
}

func continueSequenceBack(history []int) int {
	levels := make([][]int, 1)
	levels[0] = history

	// Construct the levels
	for depth, nonzero := 1, true; nonzero; depth += 1 {
		// Add a new level
		nonzero = false
		prevLevel := levels[depth-1]
		levels = append(levels, make([]int, 0))
		for i := 0; i < len(prevLevel)-1; i += 1 {
			first, second := prevLevel[i], prevLevel[i+1]
			diff := second - first
			if diff != 0 {
				nonzero = true
			}
			levels[depth] = append(levels[depth], diff)
		}
	}

	// Build the next step
	for depth := len(levels) - 3; depth >= 0; depth-- {
		lower := levels[depth+1][0]
		next := levels[depth][0]
		levels[depth] = append([]int{next - lower}, levels[depth]...)
	}

	return levels[0][0]
}

func part1(path string) {
	histories := readInput(path)
	total := 0
	for _, hist := range histories {
		total += continueSequence(hist)
	}

	fmt.Println(total)
}

func part2(path string) {
	histories := readInput(path)
	total := 0
	for _, hist := range histories {
		total += continueSequenceBack(hist)
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
