package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func readInput(inputFilePath string) []string {
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		digits := make([]rune, 0)
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}

		pair := []rune{digits[0], digits[len(digits)-1]}

		num, err := strconv.Atoi(string(pair))
		if err != nil {
			panic("Ahhh")
		}

		sum += num
	}
	println(sum)
}

func sanitizeString(weirdNum string) string {
	switch weirdNum {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return weirdNum
	}
}

func reverse(s string) string {
	size := len(s)
	flipped := make([]rune, size)
	for i, char := range s {
		flipped[size-1-i] = char
	}
	return string(flipped)
}

func part2(lines []string) {
	sum := 0
	numRe := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	flippedRe := regexp.MustCompile(`(\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)
	for _, line := range lines {
		// Do the fancy matching
		firstDigit := string(numRe.Find([]byte(line)))
		secondDigit := string(flippedRe.Find([]byte(reverse(line))))
		pair := sanitizeString(firstDigit) + sanitizeString(reverse(secondDigit))
		num, _ := strconv.Atoi(pair)
		sum += num
	}
	println(sum)
}

func main() {
	// Get input
	if len(os.Args) != 3 {
		println(len(os.Args))
		println("Usage: [INPUT FILE] [PART 1 or 2]")
	} else {
		inputFilePath := os.Args[1]
		part := os.Args[2]
		input := readInput(inputFilePath)
		if part == "1" {
			part1(input)
		} else {
			part2(input)
		}
	}
}
