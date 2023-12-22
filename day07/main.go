package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

const (
	HIGH_CARD int = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

var CARD_TO_RANK = map[rune]int{
	'1': 0, '2': 1, '3': 2,
	'4': 3, '5': 4, '6': 5,
	'7': 6, '8': 7, '9': 8,
	'T': 9, 'J': 10, 'Q': 11,
	'K': 12, 'A': 13,
}

var CARD_TO_RANK_WILD = map[rune]int{
	'J': -1, '1': 0, '2': 1, '3': 2,
	'4': 3, '5': 4, '6': 5, '7': 6,
	'8': 7, '9': 8, 'T': 9, 'Q': 11,
	'K': 12, 'A': 13,
}

func countsToType(counts []int) int {
	if counts[0] == 5 {
		return FIVE_OF_A_KIND
	}
	first, second := counts[0], counts[1]

	if first == 4 {
		return FOUR_OF_A_KIND
	} else if first == 3 && second == 2 {
		return FULL_HOUSE
	} else if first == 3 {
		return THREE_OF_A_KIND
	} else if first == 2 && second == 2 {
		return TWO_PAIR
	} else if first == 2 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

type Hand struct {
	cards  [5]rune
	counts []int
	bid    int
}

func NewHand(cardStr string, bid int) Hand {
	cards := [5]rune{}
	counter := make(map[rune]int)
	for i, card := range cardStr {
		counter[card] += 1
		cards[i] = card
	}

	var counts []int
	for _, count := range counter {
		counts = append(counts, count)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	return Hand{
		cards:  cards,
		counts: counts,
		bid:    bid,
	}
}

func (hand *Hand) Type() int {
	return countsToType(hand.counts)
}

func (hand *Hand) TypeWithWild() int {
	// Count the wilds
	numWilds := 0
	for _, card := range hand.cards {
		if card == 'J' {
			numWilds += 1
		}
	}

	if numWilds == 0 {
		return hand.Type()
	}

	// Remove that count from the counts
	newCounts := make([]int, len(hand.counts))
	copy(newCounts, hand.counts)
	for i, cnt := range newCounts {
		if cnt == numWilds {
			newCounts = remove[int](newCounts, i)
			break
		}
	}

	// In case they're all wild
	if numWilds == 5 {
		return FIVE_OF_A_KIND
	}

	// Add it to the highest available count
	newCounts[0] += numWilds

	return countsToType(newCounts)
}

func (hand *Hand) Beats(other *Hand) bool {
	if hand.Type() != other.Type() {
		return hand.Type() > other.Type()
	}

	for i := range hand.cards {
		rank, otherRank := CARD_TO_RANK[hand.cards[i]], CARD_TO_RANK[other.cards[i]]

		if rank != otherRank {
			return rank > otherRank
		}
	}

	return true
}

func (hand *Hand) BeatsWithWild(other *Hand) bool {
	if hand.TypeWithWild() != other.TypeWithWild() {
		return hand.TypeWithWild() > other.TypeWithWild()
	}

	for i := range hand.cards {
		rank, otherRank := CARD_TO_RANK_WILD[hand.cards[i]], CARD_TO_RANK_WILD[other.cards[i]]

		if rank != otherRank {
			return rank > otherRank
		}
	}

	return true
}

func readInput(path string) []Hand {
	contents, _ := os.ReadFile(path)
	lines := strings.Split(strings.Trim(string(contents), "\n "), "\n")

	hands := make([]Hand, 0)
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		bid, _ := strconv.Atoi(pieces[1])
		hands = append(hands, NewHand(pieces[0], bid))
	}
	return hands
}

func part1(path string) {
	hands := readInput(path)

	sort.Slice(hands, func(i, j int) bool {
		return hands[j].Beats(&hands[i])
	})

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}

	fmt.Println(total)
}

func part2(path string) {
	hands := readInput(path)

	sort.Slice(hands, func(i, j int) bool {
		return hands[j].BeatsWithWild(&hands[i])
	})

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
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
