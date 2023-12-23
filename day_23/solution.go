package day23

import (
	"github.com/pivovarit/aoc/util"
	"math"
	"slices"
)

func run() {
	input := util.ReadInput()

	util.Timed("aLongWalkPart1", func() int {
		return aLongWalkPart1(input)
	})
	util.Timed("aLongWalkPart2", func() int {
		return aLongWalkPart2(input)
	})
}

func aLongWalkPart1(input []string) int {
	island := parse(input)

	minX, maxX, minY, maxY := island.bounds()
	start := Location{row: minX + 1, col: minY}
	end := Location{row: maxX - 1, col: maxY}

	neighbors := buildGraph(island, start, true)

	var goal = end
	var path = 0

	return traverse(neighbors, start, goal, make(map[Location]bool), path, 0)
}

func aLongWalkPart2(input []string) int {
	island := parse(input)

	minX, maxX, minY, maxY := island.bounds()
	start := Location{row: minX + 1, col: minY}
	end := Location{row: maxX - 1, col: maxY}

	neighbours := buildGraph(island, start, false)

	var goal = end
	var path = 0

	if len(neighbours[end]) > 0 {
		goal = neighbours[end][0].location
		path = neighbours[end][0].weight
	}

	return traverse(neighbours, start, goal, make(map[Location]bool), path, 0)
}

const (
	pathTile   = '.'
	forest     = '#'
	slopeUp    = '^'
	slopeRight = '>'
	slopeLeft  = '<'
	slopeDown  = 'v'
)

type (
	Island   map[Location]rune
	Graph    map[Location][]Weight
	Location struct{ row, col int }
	Weight   struct {
		location Location
		weight   int
	}
)

func (i Island) bounds() (minX, maxX, minY, maxY int) {
	minX, maxX = math.MaxInt, math.MinInt
	minY, maxY = math.MaxInt, math.MinInt
	for p := range i {
		minX = min(p.row, minX)
		maxX = max(p.row, maxX)
		minY = min(p.col, minY)
		maxY = max(p.col, maxY)
	}
	return minX, maxX, minY, maxY
}

func traverseSingle(grid Island, previous Location, current Location, cost int, cutting bool) (Weight, bool) {
	if currentTile, exists := grid[current]; exists && currentTile != forest {
		var cpt int
		for _, neighbour := range current.adjacent() {
			if c, ok := grid[neighbour]; ok && c != forest {
				cpt++
			}
		}
		if cpt > 2 {
			return Weight{location: current, weight: cost}, true
		}
	}

	if cutting {
		if c, ok := grid[current]; ok && c != pathTile {
			if current.row > previous.row && c != slopeRight ||
				current.row < previous.row && c != slopeLeft ||
				current.col > previous.col && c != slopeDown ||
				current.col < previous.col && c != slopeUp {
				return Weight{}, false
			}
		}
	}

	for _, n := range current.adjacent() {
		if c, ok := grid[n]; ok && c != forest && n != previous {
			return traverseSingle(grid, current, n, cost+1, cutting)
		}
	}

	return Weight{location: current, weight: cost}, true
}

func (p Location) adjacent() []Location {
	return []Location{
		{p.row, p.col - 1},
		{p.row, p.col + 1},
		{p.row - 1, p.col},
		{p.row + 1, p.col}}
}

func traverse(neighbours Graph, p, goal Location, visited map[Location]bool, cost int, maxCost int) int {
	if p == goal {
		return max(cost, maxCost)
	}

	visited[p] = true
	for _, pc := range neighbours[p] {
		if !visited[pc.location] {
			maxCost = traverse(neighbours, pc.location, goal, visited, cost+pc.weight, maxCost)
		}
	}
	visited[p] = false
	return maxCost
}

func buildGraph(grid Island, start Location, cutting bool) Graph {
	var result = make(map[Location][]Weight)

	var remaining = []Location{}
	remaining = append(remaining, start)

	for len(remaining) > 0 {
		p := remaining[0]
		remaining = remaining[1:]
		if c, ok := grid[p]; !ok || c == forest {
			continue
		}
		for _, n := range p.adjacent() {
			if c, ok := grid[n]; !ok || c == forest {
				continue
			}
			pc, ok := traverseSingle(grid, p, n, 1, cutting)
			if ok && !slices.Contains(result[p], pc) {
				result[p] = append(result[p], pc)
				remaining = append(remaining, pc.location)
			}
		}
	}

	return result
}

func parse(lines []string) (result Island) {
	result = make(Island)
	for j, l := range lines {
		for i, c := range l {
			result[Location{row: i, col: j}] = c
		}
	}
	return result
}
