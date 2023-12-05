package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


type Card struct {
    WinningNums map[int]struct{}
    MyNums map[int]struct{}
}

func makeNumberSet(s string) map[int]struct{} {
    nums := strings.Split(s, " ")
    numMap := make(map[int]struct{})
    for _, num := range nums {
        if num == "" {
            continue
        }
        
        num, _ := strconv.Atoi(strings.Trim(num, " "))
        numMap[num] = struct{}{}
    }

    return numMap
}

func readInput(path string) []Card  {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)

    cards := make([]Card, 0)
	for scanner.Scan() {
		line := scanner.Text()
        
        numbers := strings.Split(line, ":")[1]
        groups := strings.Split(numbers, " | ")
        
        cards = append(cards, Card{
            WinningNums: makeNumberSet(groups[0]),
            MyNums: makeNumberSet(groups[1]),
        })
    }
	return cards
}

func part1(cards []Card) {
    total := 0.0
    for _, card := range cards {
        count := 0
        for num := range card.MyNums {
            if _, ok := card.WinningNums[num]; ok {
                count++
            }
        }

        total += math.Pow(2,float64(max(0, count - 1)))  
    }	

    println(int(total))
}

func part2(cards []Card) {
    cardValues := make([]int, len(cards))
    for i := range cardValues {
        cardValues[i] = 1
    }

    for i, card := range cards {
        count := 0
        for num := range card.MyNums {
            if _, ok := card.WinningNums[num]; ok {
                count++
            }
        }

        for k := i + 1; k <= i + count; k++ {
            cardValues[k] += cardValues[i]
        }
    }	

    // Sum the card values
    total := 0
    for _, v := range cardValues {
        total += v
    }

    println(int(total))
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
        fmt.Println(input)
		if part == "1" {
			part1(input)
		} else {
			part2(input)
		}
	}
}
