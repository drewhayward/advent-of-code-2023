package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	sets "github.com/deckarep/golang-set/v2"
)

type Point struct {
	x int
	y int
}

func readInput(path string) ([]Point, int, int) {
	contents, _ := os.ReadFile(path)
	lines := strings.Split(strings.Trim(string(contents), "\n "), "\n")

	width, height := len(lines[0]), len(lines)
	points := make([]Point, 0)
	for y, line := range lines {
		for x, elem := range line {
			if elem == '#' {
				points = append(points, Point{x, y})
			}
		}
	}

	return points, width, height
}

func scaledDist(path string, scale int) {
	points, width, height := readInput(path)

	// Build collections of points
	occupied_x := sets.NewSet[int]()
	occupied_y := sets.NewSet[int]()
	for _, point := range points {
		occupied_x.Add(point.x)
		occupied_y.Add(point.y)
	}

	// Invert the sets to get the list of rows/cols to expand
	empty_x := sets.NewSet[int]()
	for i := 0; i < width; i++ {
		if !occupied_x.Contains(i) {
			empty_x.Add(i)
		}
	}
	empty_y := sets.NewSet[int]()
	for i := 0; i < height; i++ {
		if !occupied_y.Contains(i) {
			empty_y.Add(i)
		}
	}

	// Rescale the points based off how many expansions are before them
	expand_x := sets.Sorted[int](empty_x)
	expand_y := sets.Sorted[int](empty_y)
	for i, point := range points {
		xFactor := 0
		yFactor := 0
		for curr := 0; curr < len(expand_x) && expand_x[curr] < point.x; curr++ {
			xFactor += 1
		}
		for curr := 0; curr < len(expand_y) && expand_y[curr] < point.y; curr++ {
			yFactor += 1
		}

		points[i] = Point{
			x: point.x + xFactor*(scale-1),
			y: point.y + yFactor*(scale-1),
		}
	}

	// Calc the pairwise distances
	total := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a, b := points[i], points[j]
			total += int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
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
			scaledDist(inputFilePath, 2)
		} else {
			scaledDist(inputFilePath, 1000000)
		}
	}
}
