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

type Template string

// Returns true if the provided prefix matches the template
func (t Template) matchPrefix(prefix string) bool {
	if len(prefix) > len(t) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if (t)[i] != '?' && (t)[i] != (prefix)[i] {
			return false
		}
	}

	return true
}

type Record struct {
	template Template
	groups   []int
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
			template: Template(parts[0]),
			groups:   nums,
		})
	}

	return records
}

type Entry struct {
	group_index    int
	template_index int
}

func countMatches(r Record) int {
	// DP Memoization table
	mem := make(map[Entry]int)
	return counter(mem, r, 0, 0)
}

func counter(mem map[Entry]int, r Record, group_index int, template_index int) int {
	// Return memoized value if present
	key := Entry{group_index, template_index}
	if val, ok := mem[key]; ok {
		return val
	}

	// There are no more groups to place
	if group_index >= len(r.groups) {
		prefix := strings.Repeat(".", len(r.template)-template_index)

		if r.template[template_index:].matchPrefix(prefix) {
			mem[key] = 1
		} else {
			mem[key] = 0
		}
	} else {
		total := 0

		// Need to count the number of ways to place the rest of the items
		// |---------------------length----------------------|
		// |----------room to play-------|-----leftover------|
		//    |--size--| <- the group that needs to be placed
		//    ^ start
		length := len(r.template) - template_index
		size := r.groups[group_index]
		leftover := len(r.groups[group_index+1:]) + sum(r.groups[group_index+1:])
		for start := 0; start <= length-leftover-size; start++ {
			// Every group after the first needs to lead with a "."
			if group_index != 0 && start == 0 {
				continue
			}

			prefix := strings.Repeat(".", start) + strings.Repeat("#", size)
			if r.template[template_index:].matchPrefix(prefix) {
				total += counter(mem, r, group_index+1, template_index+len(prefix))
			}
		}

		mem[key] = total
	}

	return mem[key]
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

		repeated := make([]int, 0, len(record.groups)*scale)
		repPat := Template("")
		for i := 0; i < scale; i++ {
			if i != 0 {
				repPat += "?"
			}
			repPat += record.template
			repeated = append(repeated, record.groups...)
		}
		record.groups = repeated
		record.template = repPat

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
