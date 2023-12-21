package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	key   string
	left  string
	right string
}

const LINE_PATTERN = `(\w+) = \((\w+), (\w+)\)`

func stringToRunes(s string) []rune {
	runes := make([]rune, 0)
	for _, r := range s {
		runes = append(runes, r)
	}
	return runes
}

func readInput(path string) (string, map[string]Node) {
	contents, _ := os.ReadFile(path)
	parts := strings.Split(string(contents), "\n\n")
	instructions := parts[0]
	re := regexp.MustCompile(LINE_PATTERN)

	lines := strings.Split(parts[1], "\n")
	nodes := make(map[string]Node)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			continue
		}

		n := Node{
			key:   matches[1],
			left:  matches[2],
			right: matches[3],
		}
		nodes[n.key] = n
	}

	return instructions, nodes
}

func part1(path string) {
	instrs, nodes := readInput(path)
	ins_runes := stringToRunes(instrs)

	i := 0
	steps := 0
	curr := "AAA"
	for curr != "ZZZ" {
		println(curr)
		fmt.Println(nodes[curr])
		if ins_runes[i] == 'R' {
			curr = nodes[curr].right
		} else {
			curr = nodes[curr].left
		}

		i = (i + 1) % len(instrs)
		steps += 1
	}
	println(steps)
}

func part2(path string) {
	instrs, nodes := readInput(path)
	ins_runes := stringToRunes(instrs)

	positions := make([]string, 0)
	for key := range nodes {
		if strings.HasSuffix(key, "A") {
			positions = append(positions, key)
		}
	}

	stillSearching := func(nodes []string) bool {
		for _, node := range nodes {
			if !strings.HasSuffix(node, "Z") {
				return true
			}
		}
		return false
	}

	i := 0
	steps := 0
	for stillSearching(positions) {
		// fmt.Println(positions)
		for j := range positions {
			if ins_runes[i] == 'R' {
				positions[j] = nodes[positions[j]].right
			} else {
				positions[j] = nodes[positions[j]].left
			}
		}

		i = (i + 1) % len(instrs)
		steps += 1
	}
	println(steps)
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
