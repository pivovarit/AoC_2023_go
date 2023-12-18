package day17

import (
	"github.com/pivovarit/aoc/util"
	"math"
)

func run() {
	input := util.ReadInput()

	util.Timed("clumsyCruciblePart1", func() int {
		return clumsyCruciblePart1(input)
	})
	util.Timed("clumsyCruciblePart2", func() int {
		return clumsyCruciblePart2(input)
	})
}

func clumsyCruciblePart1(input []string) int {
	start := Point{0, 0}
	end := Point{len(input[0]) - 1, len(input) - 1}

	return pathWeight(input, start, end, 1, 3)
}

func clumsyCruciblePart2(input []string) int {
	start := Point{0, 0}
	end := Point{len(input[0]) - 1, len(input) - 1}

	return pathWeight(input, start, end, 4, 10)
}

type (
	Point struct {
		row, col int
	}
	Direction struct {
		x, y int
	}
	State struct {
		point     Point
		direction Direction
		streak    int
	}
)

func (p Direction) left() Direction {
	return Direction{p.y, -p.x}
}

func (p Direction) right() Direction {
	return Direction{-p.y, p.x}
}

func pathWeight(grid []string, start, end Point, minSameDirection, maxSameDirection int) int {
	remaining := []State{{start, Direction{1, 0}, 0}, {start, Direction{0, 1}, 0}}
	visited := map[State]int{{start, Direction{0, 0}, 0}: 0}
	minWeight := math.MaxInt

	for len(remaining) > 0 {
		current := remaining[0]
		remaining = remaining[1:]

		if current.point == end && current.streak >= minSameDirection {
			minWeight = min(minWeight, visited[current])
		}

		for _, direction := range directionsFor(current) {
			next := Point{current.point.row + direction.x, current.point.col + direction.y}
			if !next.inside(grid) {
				continue
			}

			total := visited[current] + int(grid[next.col][next.row]-'0')
			if (direction == current.direction && current.streak < maxSameDirection) || (direction != current.direction && current.streak >= minSameDirection) {
				state := State{next, direction, getDirectionCount(direction, current)}
				if value, found := visited[state]; !found || value > total {
					visited[state] = total
					remaining = append(remaining, state)
				}
			}
		}
	}

	return minWeight
}

func getDirectionCount(direction Direction, current State) int {
	if direction == current.direction {
		return current.streak + 1
	} else {
		return 1
	}
}

func directionsFor(state State) [3]Direction {
	return [3]Direction{state.direction, state.direction.left(), state.direction.right()}
}

func (p Point) inside(board []string) bool {
	return p.row >= 0 && p.row < len(board[0]) && p.col >= 0 && p.col < len(board)
}
