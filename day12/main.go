package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(ints []int) int {
	total := 0
	for _, num := range ints {
		total += num
	}
	return total
}

// Generate the possible nonogram row options
// nums: group sizes left to place
// length: the number of spaces to place the remaining groups in
// prefix: the prefix to be added to the options
// c: the channel to send the results over
func generateOptions(nums []int, length int, prefix string, c chan string) {
	// Base case
	if len(nums) == 0 {
		c <- prefix + strings.Repeat(".", length)
		return
	}

	groupSize := nums[0]
	leftover := len(nums[1:]) + sum(nums[1:])
	// fmt.Println(length, groupSize, leftover)
	for start := 0; start <= length-leftover-groupSize; start++ {
		if prefix != "" && start == 0 {
			continue
		}
		newPrefix := strings.Repeat(".", start) + strings.Repeat("#", groupSize)
		// If there are more groups to place, we need to make sure we
		// Add a spacer on the last section
		// if start == length-leftover-groupSize && len(nums) > 1 {
		// 	newPrefix += "."
		// }
		generateOptions(nums[1:], length-len(newPrefix), prefix+newPrefix, c)
	}

	if prefix == "" {
		close(c)
	}
}

type Record struct {
	pattern string
	groups  []int
}

func countMatches(record Record) int {
	optChan := make(chan string)
	go generateOptions(record.groups, len(record.pattern), "", optChan)

	matches := 0
	nopts := 0
	for {
		option, ok := <-optChan
		if !ok {
			break
		}

		nopts += 1
		if nopts%1000 == 0 {
			fmt.Println(nopts, option)
		}

		match := true
		for i := 0; i < len(record.pattern); i++ {
			if record.pattern[i] != '?' && record.pattern[i] != option[i] {
				match = false
			}
		}

		if match {
			matches += 1
		}
	}
	return matches
}

func readInput(path string) []Record {
	contents, _ := os.ReadFile(path)
	lines := strings.Split(strings.Trim(string(contents), "\n"), "\n")
	records := make([]Record, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")

		nums := make([]int, 0)
		for _, numStr := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		records = append(records, Record{
			pattern: parts[0],
			groups:  nums,
		})
	}

	return records
}

func part1(path string) {
	records := readInput(path)

	total := 0
	for _, record := range records {
		total += countMatches(record)
	}

	fmt.Println(total)
}

func part2(path string) {
	records := readInput(path)
	scale := 5

	total := 0
	for _, record := range records {
		fmt.Println("solving")
		record.pattern = strings.Repeat(record.pattern, scale)

		repeated := make([]int, 0, len(record.groups)*scale)
		for i := 0; i < scale; i++ {
			repeated = append(repeated, record.groups...)
		}
		record.groups = repeated

		fmt.Println(record)

		total += countMatches(record)
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
