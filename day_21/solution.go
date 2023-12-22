package day21

import (
	"container/list"
	"github.com/pivovarit/aoc/util"
)

func run() {
	input := util.ReadInput()

	util.Timed("stepCounterPart1", func() int {
		return stepCounterPart1(input)
	})
	util.Timed("stepCounterPart2", func() int {
		return stepCounterPart2(input)
	})
}

const (
	rocksTile = '#'
	startTile = 'S'
)

type (
	Direction struct{ row, col int }
	Location  struct{ row, col int }
)

var directions = [4]Direction{
	{0, -1}, // left
	{0, 1},  // right
	{1, 0},  // down
	{-1, 0}, // up
}

type QueueItem struct {
	location Location
	step     int
}

func stepCounterPart1(input []string) int {
	return traverse(input, findStart(input), 64, false)
}

func traverse(input []string, start Location, steps int, boundless bool) int {
	queue := list.New()
	queue.PushBack(QueueItem{start, 0})
	visited := make(map[QueueItem]bool)
	plots := 0
	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(QueueItem)
		queue.Remove(element)
		if node.step == steps {
			plots++
			continue
		}
		for _, direction := range directions {
			if boundless {
				newRow := node.location.row + direction.row
				newCol := node.location.col + direction.col

				if newRow >= len(input) {
					newRow = newRow % len(input)
				} else if newRow < 0 {
					if newRow%len(input) == 0 {
						newRow = len(input) - 1
					} else {
						newRow = len(input) + newRow%len(input)
					}
				}

				if newCol >= len(input[0]) {
					newCol = newCol % len(input[0])
				} else if newCol < 0 {
					if newCol%len(input[0]) == 0 {
						newCol = len(input[0]) - 1
					} else {
						newCol = len(input[0]) + newCol%len(input[0])
					}
				}

				newLocation := Location{node.location.row + direction.row, node.location.col + direction.col}
				if input[newRow][newCol] == rocksTile {
					continue
				}
				plot := QueueItem{newLocation, node.step + 1}
				if !visited[plot] {
					queue.PushBack(plot)
					visited[plot] = true
				}
			} else {
				newLocation := Location{node.location.row + direction.row, node.location.col + direction.col}
				if newLocation.row < 0 || newLocation.row >= len(input) || newLocation.col < 0 || newLocation.col >= len(input[0]) || input[newLocation.row][newLocation.col] == rocksTile {
					continue
				}
				plot := QueueItem{newLocation, node.step + 1}
				if !visited[plot] {
					queue.PushBack(plot)
					visited[plot] = true
				}
			}
		}
	}
	return plots
}

func findStart(input []string) Location {
	for row, line := range input {
		for col, char := range line {
			if char == startTile {
				return Location{row, col}
			}
		}
	}

	panic("could not find start")
}

func stepCounterPart2(input []string) int {
	steps := 202300
	start := findStart(input)
	size := len(input)
	half := size / 2
	p := make([]int, 3)
	p[0] = traverse(input, start, half, true)
	p[1] = traverse(input, start, half+size, true)
	p[2] = traverse(input, start, half+2*size, true)
	a := (p[2] + p[0] - 2*p[1]) / 2
	return a*steps*steps + (p[1]-p[0]-a)*steps + p[0]
}
