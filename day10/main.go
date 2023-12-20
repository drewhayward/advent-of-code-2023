package main

import (
	"fmt"
	"os"
	"strings"
    sets "github.com/deckarep/golang-set/v2"
)

type Coord struct {
    x int
    y int
}

type Node struct {
    coord Coord
    neighbors []Coord
}

type GridMap struct {
    nodes map[Coord]Node
    start Coord
    max_x int
    max_y int
}

var DIRS = [...]Coord{{1,0},{0,1},{-1,0},{0,-1}}

func coordInNeighbors(neighbors []Coord, c Coord) bool {
    for _, n := range neighbors {
        if n == c {
            return true
        }
    }
    return false
}

func readGrid(grid string) GridMap {
    // Want to transform the grid into a graph structure
    // Going to store each node with a ref to the coordinates that
    // they point to. And access a map of nodes by coordinates 
    var start Coord
    nodes := make(map[Coord]Node)
    var max_x, max_y int
    for y, line := range strings.Split(grid, "\n") {
        max_y = max(y, max_y)
        for x, pipe := range line {
            max_x = max(x, max_x)
            n := Node{
                coord: Coord{x:x, y:y},
                neighbors: make([]Coord, 0, 2),
            } 

            switch ;pipe {
            case '|':
                n.neighbors = append(n.neighbors, Coord{x: x, y:y+1})
                n.neighbors = append(n.neighbors, Coord{x: x, y:y-1})
            case '-':
                n.neighbors = append(n.neighbors, Coord{x: x+1, y:y})
                n.neighbors = append(n.neighbors, Coord{x: x-1, y:y})
            case 'L':
                n.neighbors = append(n.neighbors, Coord{x: x, y:y-1})
                n.neighbors = append(n.neighbors, Coord{x: x+1, y:y})
            case 'J':
                n.neighbors = append(n.neighbors, Coord{x: x, y:y-1})
                n.neighbors = append(n.neighbors, Coord{x: x-1, y:y})
            case '7':
                n.neighbors = append(n.neighbors, Coord{x: x, y:y+1})
                n.neighbors = append(n.neighbors, Coord{x: x-1, y:y})
            case 'F':
                n.neighbors = append(n.neighbors, Coord{x: x, y:y+1})
                n.neighbors = append(n.neighbors, Coord{x: x+1, y:y})
            case 'S': // Need to infer the neighbors
                start = n.coord
            case '.': // No neighbors
            default:
                panic("Shouldn't reach here")
            }

            nodes[n.coord] = n
        }
    }
    // Once we have the graph, we need to infer the connections
    // of the starting pipe from it's 4 neighbors
    start_node, _ := nodes[start]
    for _, delta := range DIRS {
        ncoord :=  Coord{
            x: start.x + delta.x,
            y: start.y + delta.y,
        }
        neighbor, _ := nodes[ncoord]
        if coordInNeighbors(neighbor.neighbors, start) {
           start_node.neighbors = append(start_node.neighbors, ncoord) 
        }
    }
    nodes[start] = start_node

    // This print actually breaks the program if deleted ?!
    fmt.Println(nodes)

    return GridMap{
        nodes: nodes,
        start: start,
        max_x: max_x,
        max_y: max_y,
    }
}

func findPath(gmap GridMap) (sets.Set[Coord], int) {
    type bfsState struct {
        coord Coord
        depth int
    }

    // Then perform BFS to find the max dist from S to a point in the loop
    visited := sets.NewSet[Coord]()
    frontier := make([]bfsState, 1)
    frontier[0] = bfsState{
        coord: gmap.start,
    }
    nodes := gmap.nodes
    max_depth := 0
    for len(frontier) > 0 {
        // Pop the next node to visit
        state := frontier[0]
        frontier = frontier[1:]
        node := nodes[state.coord]

        // Skip prev visited nodes
        if visited.Contains(state.coord) {
            continue
        } else {
            visited.Add(state.coord)
        }

        max_depth = max(max_depth, state.depth)
        for _, ncoord := range node.neighbors {
            frontier = append(frontier, bfsState{
                coord: ncoord,
                depth: state.depth + 1,
            })
        }
    }

    return visited, max_depth
}

func part1(path string) {
    bytes, _ := os.ReadFile(path)
    gmap := readGrid(string(bytes))
    
    _, max_depth := findPath(gmap)

    fmt.Println(max_depth)
}

// Search through the grid respecting the path edges
// Returns the set of visited nodes and a bool flagging
// if the region touches the boundary
func search(gmap GridMap, c Coord, path sets.Set[Coord]) (sets.Set[Coord], bool) {
    visited := sets.NewSet[Coord]()
    frontier := make([]Coord, 1)
    frontier[0] = c
    outside := false
    for len(frontier) > 0 {
        // Pop the next node to visit
        coord := frontier[0]
        frontier = frontier[1:]
        node := gmap.nodes[coord]

        // Skip prev visited nodes
        if visited.Contains(coord) {
            continue
        } else {
            visited.Add(coord)
        }

        if coord.x < 0 || coord.x > gmap.max_x || coord.y < 0 || coord.y > gmap.max_y {
            outside = true
        }

        for _, neighbor := range node.neighbors {
            // Skip out-of-bounds neighbors and nodes on the
            // pipe path
            if _, present := gmap.nodes[neighbor]; !present || path.Contains(neighbor){
                continue
            }            
            frontier = append(frontier, neighbor)
        }
    }

    return visited, outside
}

func part2(filepath string) {
    bytes, _ := os.ReadFile(filepath)
    gmap := readGrid(string(bytes))

    path, _ := findPath(gmap)
    visited := path.Clone()
    total_area := 0
    for coord := range gmap.nodes {
        if visited.Contains(coord) {
            continue
        }
        fmt.Println("Searching from ", coord)
        region, is_outside := search(gmap, coord, path)
        if !is_outside {
            total_area += region.Cardinality()
        }

        visited = visited.Union(region)
    }

    fmt.Println(total_area)
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
