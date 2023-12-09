package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func readInput(inputFilePath string) ([]int, []int) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digitRe := regexp.MustCompile(`\d+`)

	times := make([]int, 0)
	scanner.Scan()
	for _, time := range digitRe.FindAllString(scanner.Text(), -1) {
		num, _ := strconv.Atoi(time)
		times = append(times, num)
	}
	dists := make([]int, 0)
	scanner.Scan()
	for _, time := range digitRe.FindAllString(scanner.Text(), -1) {
		num, _ := strconv.Atoi(time)
		dists = append(dists, num)
	}

	return times, dists
}

func waysToWin(tint int, dint int) int {
	t, d := float64(tint), float64(dint)
	r1 := (t + math.Sqrt(math.Pow(t, 2)-4*d)) / 2.0
	r2 := (t - math.Sqrt(math.Pow(t, 2)-4*d)) / 2.0
	kmin := int(math.Ceil(min(r1, r2)))
	kmax := int(math.Floor(max(r1, r2)))

	n_wins := kmax - kmin + 1

	if kmin*(tint-kmin) <= dint {
		n_wins--
	}

	if kmax*(tint-kmax) <= dint {
		n_wins--
	}

	return n_wins
}

func part1(inputFilePath string) {
	ts, ds := readInput(inputFilePath)

	total := 1
	for i := range ts {
		total *= waysToWin(ts[i], ds[i])
	}
	println(total)
}

func part2(path string) {
	times, dists := readInput(path)

	bigtime := 0
	for _, t := range times {
		p := math.Ceil(math.Log10(float64(t)))
		bigtime = bigtime*int(math.Pow(10, p)) + t
	}

	bigdist := 0
	for _, d := range dists {
		p := math.Ceil(math.Log10(float64(d)))
		bigdist = bigdist*int(math.Pow(10, p)) + d
	}

	println(waysToWin(bigtime, bigdist))
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
