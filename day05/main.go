package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)



type Remap struct {
    dest int
    source int
    size int
}

type NumRange struct {
    start int
    end int
}

const DIGIT_REGEX = `\d+`

func readInput(path string) ([]int, [][]Remap)  {
    digitPattern := regexp.MustCompile(DIGIT_REGEX)
    contents, _ := os.ReadFile(path)

    sections := strings.Split(string(contents), "\n\n")
    
    // Make seedNums
    seedNums := make([]int, 0)
    for _, seed := range digitPattern.FindAllString(sections[0], -1) {
        num, _ := strconv.Atoi(seed)
        seedNums = append(seedNums, num)
    }

    
    // Save remap structure
    gardenMaps := make([][]Remap, len(sections) - 1)
    for i, section := range sections[1:] {
        gardenMaps[i] = make([]Remap, 0)
        var dst, src, sz int
        for _, line := range strings.Split(section, "\n")[1:] {
            fmt.Sscanf(line, "%d %d %d", &dst, &src, &sz)
            gardenMaps[i] = append(gardenMaps[i], Remap{source: src, dest: dst, size: sz})
        }
        sort.Slice(gardenMaps[i], func (j, k int) bool {
            return gardenMaps[i][j].source < gardenMaps[i][k].source
        })
    }
    

    return seedNums, gardenMaps
}

func mapSeed(seed int, remaps [][]Remap) int {
    curr := seed
    for _, m := range remaps {
        for _, remap := range m {
            // Check if the current value applies to a remap
            if curr <= remap.source + remap.size && remap.source <= curr {
                curr = remap.dest + (curr - remap.source)
                break
            }
        }
    }
    return curr
}

func part1(path string) {
    seedNums, maps := readInput(path)
     
    smallestLoc := math.MaxInt32
    for _, seed := range seedNums {
        smallestLoc = min(smallestLoc, mapSeed(seed, maps))
    }

    fmt.Printf("Nearest location needing a seed: %d\n", smallestLoc)
}

func part2(path string) {
    seedNums, maps := readInput(path)

    smallestLoc := math.MaxInt32

    // Keep track of a list of ranges from input and 
    // transform them through the remaps
    rngs := make([]NumRange, 0)
    for i := 0; i < len(seedNums); i += 2 {
        startSeed := seedNums[i]
        numSeeds := seedNums[i + 1]
        rngs = append(rngs, NumRange{start: startSeed, end: startSeed + numSeeds})
    }
    sort.Slice(rngs, func(i,j int) bool {
        return rngs[i].start < rngs[j].start
    })

    for _, m := range maps {
        newRanges := make([]NumRange, len(rngs))

        irng := 0
        imap := 0
        for true {

            // Skip 
            if rngs[irng].end < m[imap].source {
                irng++
                continue
            } else if m[imap].source + m[imap].size {
                imap++
                continue
            }

        }    
    }

    fmt.Printf("Nearest location needing a seed: %d\n", smallestLoc)
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
