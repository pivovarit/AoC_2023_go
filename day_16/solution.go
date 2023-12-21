package day16

import (
	"github.com/pivovarit/aoc/util"
)

func run() {
	input := util.ReadInput()

	util.Timed("theFloorWillBeLavaPart1", func() int {
		return theFloorWillBeLavaPart1(input)
	})
	util.Timed("theFloorWillBeLavaPart2", func() int {
		return theFloorWillBeLavaPart2(input)
	})
}

func theFloorWillBeLavaPart1(input []string) int {
	return energize(Location{row: 0, col: 0}, west, input, initVisited(input))
}

func energize(location Location, direction Direction, input []string, visited [][]bool) int {
	beam(location, direction, input, visited, map[LocationWithDirection]bool{})

	var sum = 0
	for _, bools := range visited {
		for _, energized := range bools {
			if energized {
				sum++
			}
		}
	}
	return sum
}

func theFloorWillBeLavaPart2(input []string) int {
	var highest = -1

	for i := 0; i < len(input); i++ {
		highest = max(
			highest,
			energize(Location{row: i, col: 0}, west, input, initVisited(input)),
			energize(Location{row: i, col: len(input[0]) - 1}, east, input, initVisited(input)))
	}

	for i := 0; i < len(input[0]); i++ {
		highest = max(
			highest,
			energize(Location{row: 0, col: i}, north, input, initVisited(input)),
			energize(Location{row: len(input) - 1, col: i}, south, input, initVisited(input)))
	}
	return highest
}

func initVisited(input []string) [][]bool {
	var visited = make([][]bool, len(input))
	for i := 0; i < len(input); i++ {
		visited[i] = make([]bool, len(input[i]))
	}
	return visited
}

func beam(location Location, direction Direction, data []string, visited [][]bool, seen map[LocationWithDirection]bool) {
	if location.row < 0 || location.col < 0 ||
		seen[location.withDirection(direction)] ||
		location.row > len(data)-1 || location.col > len(data[location.row])-1 {
		return
	}
	visited[location.row][location.col] = true
	seen[location.withDirection(direction)] = true

	tile := rune(data[location.row][location.col])

	switch tile {
	case '.':
		switch direction {
		case west:
			beam(location.right(), west, data, visited, seen)
		case east:
			beam(location.left(), east, data, visited, seen)
		case north:
			beam(location.down(), north, data, visited, seen)
		case south:
			beam(location.up(), south, data, visited, seen)
		}
	case '\\':
		switch direction {
		case west:
			beam(location.down(), north, data, visited, seen)
		case east:
			beam(location.up(), south, data, visited, seen)
		case north:
			beam(location.right(), west, data, visited, seen)
		case south:
			beam(location.left(), east, data, visited, seen)
		}
	case '/':
		switch direction {
		case east:
			beam(location.down(), north, data, visited, seen)
		case west:
			beam(location.up(), south, data, visited, seen)
		case north:
			beam(location.left(), east, data, visited, seen)
		case south:
			beam(location.right(), west, data, visited, seen)
		}
	case '|':
		switch direction {
		case east, west:
			beam(location.up(), south, data, visited, seen)
			beam(location.down(), north, data, visited, seen)
		case north:
			beam(location.down(), north, data, visited, seen)
		case south:
			beam(location.up(), south, data, visited, seen)
		}
	case '-':
		switch direction {
		case east:
			beam(location.left(), east, data, visited, seen)
		case west:
			beam(location.right(), west, data, visited, seen)
		case north, south:
			beam(location.left(), east, data, visited, seen)
			beam(location.right(), west, data, visited, seen)
		}
	}
}

type (
	Direction             int
	Location              struct{ row, col int }
	LocationWithDirection struct {
		location  Location
		direction Direction
	}
)

const (
	east Direction = iota
	west
	north
	south
)

func (l *Location) right() Location {
	return Location{l.row, l.col + 1}
}

func (l *Location) left() Location {
	return Location{l.row, l.col - 1}
}

func (l *Location) up() Location {
	return Location{l.row - 1, l.col}
}
func (l *Location) down() Location {
	return Location{l.row + 1, l.col}
}

func (l *Location) withDirection(direction Direction) LocationWithDirection {
	return LocationWithDirection{location: *l, direction: direction}
}
