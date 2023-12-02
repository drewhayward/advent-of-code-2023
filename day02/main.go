package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CubeDraw struct {
	color string
	num   int
}

type Set struct {
	draws []CubeDraw
}

type Game struct {
	id   int
	sets []Set
}

const DRAWPATTERN = `(\d+) (blue|green|red)`

func readInput(path string) []Game {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	drawReg := regexp.MustCompile(DRAWPATTERN)
	games := make([]Game, 0)
	for scanner.Scan() {
		line := scanner.Text()
		game := Game{}

		lineChunks := strings.Split(line, ":")
		fmt.Sscanf(lineChunks[0], "Game %d", &game.id)
		for _, setString := range strings.Split(lineChunks[1], ";") {
			set := Set{}

			for _, draw := range drawReg.FindAllStringSubmatch(setString, -1) {
				num, _ := strconv.Atoi(draw[1])
				set.draws = append(set.draws, CubeDraw{
					num:   num,
					color: draw[2],
				})
			}

			game.sets = append(game.sets, set)
		}

		games = append(games, game)
	}
	return games
}

func part1(games []Game) {
	idSum := 0
	for _, game := range games {
		isGamePossible := true
		for _, set := range game.sets {
			for _, draw := range set.draws {
				if draw.color == "red" && draw.num > 12 {
					isGamePossible = false
				} else if draw.color == "green" && draw.num > 13 {
					isGamePossible = false
				} else if draw.color == "blue" && draw.num > 14 {
					isGamePossible = false
				}
			}
		}

		if isGamePossible {
			idSum += game.id
		}
	}

	println(idSum)
}

func part2(games []Game) {
	powerSum := 0
	for _, game := range games {
		rm := 0
		gm := 0
		bm := 0
		for _, set := range game.sets {
			for _, draw := range set.draws {
				if draw.color == "red" {
					rm = max(rm, draw.num)
				} else if draw.color == "green" {
					gm = max(gm, draw.num)
				} else if draw.color == "blue" {
					bm = max(bm, draw.num)
				}
			}
		}
		powerSum += rm * gm * bm
	}

	println(powerSum)
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
