package main

import (
	"fmt"
	"os"
	"strings"

	sets "github.com/deckarep/golang-set/v2"
)

type Direction struct {
	dx int
	dy int
}

var DIRS = [4]Direction{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type Point struct {
	x int
	y int
}

func (p Point) step(d Direction) Point {
	return Point{
		p.x + d.dx,
		p.y + d.dy,
	}
}

type State struct {
	p   Point
	dir Direction
}

type Board struct {
	elements [][]rune
	max_x    int
	max_y    int
}

func (b *Board) neighbors(s State) []State {
	neighbors := make([]State, 0, 4)

	advance := func(s State, dir Direction) {
		s.p = s.p.step(dir)
		s.dir = dir
		if s.p.x >= 0 && s.p.x < b.max_x && s.p.y >= 0 && s.p.y < b.max_y {
			neighbors = append(neighbors, s)
		}
	}
	switch element := b.elements[s.p.y][s.p.x]; element {
	case '/':
		advance(s, Direction{
			dx: -s.dir.dy,
			dy: -s.dir.dx,
		})
	case '\\':
		advance(s, Direction{
			dx: s.dir.dy,
			dy: s.dir.dx,
		})
	case '|': // Vertical splitter
		if s.dir.dy == 0 {
			advance(s, Direction{0, 1})
			advance(s, Direction{0, -1})
		} else {
			advance(s, s.dir)
		}
	case '-': // Horizontal splitter
		if s.dir.dx == 0 {
			// Split into the horizontal dirs
			advance(s, Direction{1, 0})
			advance(s, Direction{-1, 0})
		} else {
			advance(s, s.dir)
		}
	case '.':
		advance(s, s.dir)
	}

	return neighbors
}

func readInput(filePath string) Board {
	contents, _ := os.ReadFile(filePath)

	board := make([][]rune, 0)
	var max_x, max_y int
	for y, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			break
		}
		board = append(board, make([]rune, 0))

		for _, elem := range line {
			board[y] = append(board[y], elem)
		}

		max_x = max(len(line), max_x)
		max_y++
	}

	return Board{
		elements: board,
		max_x:    max_x,
		max_y:    max_y,
	}
}

func getNumEnergized(start State, board Board) int {
	visited := sets.NewSet[State]()
	visitedPositions := sets.NewSet[Point]()
	frontier := make([]State, 0)
	frontier = append(frontier, start)

	var current State
	for len(frontier) > 0 {
		current, frontier = frontier[0], frontier[1:]
		// fmt.Printf("At %v heading in %v\n", current.p, current.dir)

		visitedPositions.Add(current.p)
		if visited.Contains(current) {
			continue
		}
		visited.Add(current)

		for _, neighbor := range board.neighbors(current) {
			frontier = append(frontier, neighbor)
		}
	}
	return visitedPositions.Cardinality()
}

func part1(path string) {
	board := readInput(path)

	fmt.Println(getNumEnergized(State{
		p:   Point{0, 0},
		dir: Direction{1, 0},
	}, board))
}

func part2(path string) {
	board := readInput(path)

	mostEnergized := 0
	// top
	for x := 0; x < board.max_x; x++ {
		mostEnergized = max(mostEnergized, getNumEnergized(
			State{
				dir: Direction{0, 1},
				p:   Point{x, 0},
			}, board))
	}
	// left
	for y := 0; y < board.max_y; y++ {
		mostEnergized = max(mostEnergized, getNumEnergized(
			State{
				dir: Direction{1, 0},
				p:   Point{0, y},
			}, board))
	}
	// right
	for y := 0; y < board.max_y; y++ {
		mostEnergized = max(mostEnergized, getNumEnergized(
			State{
				dir: Direction{-1, 0},
				p:   Point{board.max_x - 1, y},
			}, board))
	}
	// bottom
	for x := 0; x < board.max_x; x++ {
		mostEnergized = max(mostEnergized, getNumEnergized(
			State{
				dir: Direction{0, -1},
				p:   Point{x, board.max_y - 1},
			}, board))
	}

	fmt.Println(mostEnergized)
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
