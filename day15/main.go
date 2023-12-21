package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

type HashMap[T any] struct {
	table  [][]T
	hash   func(T) int
	itemEq func(T, T) bool
	size   int
}

func NewHashMap[T any](size int, hash func(T) int, eq func(T, T) bool) HashMap[T] {
	table := make([][]T, size)
	for i := range table {
		table[i] = make([]T, 0)
	}

	return HashMap[T]{
		table:  table,
		hash:   hash,
		itemEq: eq,
		size:   size,
	}
}

func (hashmap *HashMap[T]) Add(item T) {
	boxIdx := hashmap.hash(item)

	// Swap item if present
	for i, other := range hashmap.table[boxIdx] {
		if hashmap.itemEq(item, other) {
			hashmap.table[boxIdx][i] = item
			return
		}
	}

	// Otherwise append
	hashmap.table[boxIdx] = append(hashmap.table[boxIdx], item)
}

func (hashmap *HashMap[T]) Print() {
	for boxNum, box := range hashmap.table {
		if len(box) > 0 {
			fmt.Printf("Box %d: %v\n", boxNum, box)
		}
	}
}

func (hashmap *HashMap[T]) Remove(item T) {
	boxIdx := hashmap.hash(item)

	// Remove item if present
	for i, other := range hashmap.table[boxIdx] {
		if hashmap.itemEq(item, other) {
			hashmap.table[boxIdx] = remove[T](hashmap.table[boxIdx], i)
		}
	}
}

type Entry struct {
	label string
	lens  int
}

func readInput(path string) []string {
	contents, _ := os.ReadFile(path)

	initLine := strings.Trim(string(contents), "\n ")

	return strings.Split(initLine, ",")
}

func hash(value string) int {
	current := 0
	for _, char := range value {
		current += int(char)
		current *= 17
		current = current % 256
	}

	return current
}

func hashEntry(e Entry) int {
	return hash(e.label)
}

func entryEq(e1 Entry, e2 Entry) bool {
	return e1.label == e2.label
}

func part1(path string) {
	total := 0
	for _, init := range readInput(path) {
		total += hash(init)
	}
	fmt.Println(total)
}

func part2(path string) {
	hm := NewHashMap[Entry](
		256,
		hashEntry,
		entryEq,
	)

	pat := regexp.MustCompile(`(\w+)(=|-)(\d?)`)
	for _, instruction := range readInput(path) {
		results := pat.FindStringSubmatch(instruction)
		label, command := results[1], results[2]
		lens, _ := strconv.Atoi(results[3])

		entry := Entry{
			label: label,
			lens:  lens,
		}

		if command == "=" {
			hm.Add(entry)
		} else if command == "-" {
			hm.Remove(entry)
		}
	}

	// Get the focusing power
	total := 0
	for boxNum, box := range hm.table {
		for slotNum, entry := range box {
			power := (boxNum + 1) * (slotNum + 1) * entry.lens
			total += power
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
