package day14

import (
	"github.com/pivovarit/aoc/util"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("parabolicReflectorDishPart1", func() int {
		return parabolicReflectorDishPart1(input)
	})
	util.Timed("parabolicReflectorDishPart2", func() int {
		return parabolicReflectorDishPart2(input)
	})
}

const (
	cubeRock   = '#'
	roundRock  = 'O'
	emptySpace = '.'
)

const cycles = 1_000_000_000

type Direction int

const (
	north Direction = iota
	west
	south
	east
)

func parabolicReflectorDishPart1(input []string) int {
	sum := 0

	for colId := 0; colId < len(input[0]); colId++ {
		count := 0
		lastCube := -1
		for rowId, row := range input {
			switch row[colId] {
			case cubeRock:
				sum += calculateLoad(count, len(input)-lastCube-1)
				lastCube = rowId
				count = 0
			case roundRock:
				count++
			}
		}
		sum += calculateLoad(count, len(input)-lastCube-1)
	}
	return sum
}

func getLoad(board [][]rune) int {
	weight := 0

	for rowId := range board {
		for colId := range board[0] {
			if board[rowId][colId] != roundRock {
				continue
			}

			weight += len(board) - rowId
		}
	}

	return weight
}

func calculateLoad(rocks int, startingRow int) int {
	if rocks == 1 {
		return startingRow
	} else {
		return arithmeticSeriesSum(startingRow-rocks+1, startingRow, rocks)
	}
}

func arithmeticSeriesSum(first, last, count int) int {
	return ((first + last) * count) / 2
}

func parabolicReflectorDishPart2(input []string) int {
	var charInput = make([][]rune, len(input))
	for i, s := range input {
		charInput[i] = make([]rune, len(s))
	}
	for rowId, s := range input {
		for colId, val := range s {
			charInput[rowId][colId] = val
		}
	}

	var seen = make(map[string]int)
	seen[hash(charInput)] = 0
	last := charInput
	for iteration := 1; iteration < cycles; iteration++ {
		last = tiltCycle(last)
		hashed := hash(last)
		cycleId, exists := seen[hashed]
		if exists {
			cycleLength := iteration - cycleId
			for i := 0; i < (cycles-iteration)%cycleLength; i++ {
				last = tilt(east, tilt(south, tilt(west, tilt(north, last))))
			}

			return getLoad(last)
		} else {
			seen[hashed] = iteration
		}
	}

	return -1
}

func tiltCycle(last [][]rune) [][]rune {
	last = tilt(north, last)
	last = tilt(west, last)
	last = tilt(south, last)
	last = tilt(east, last)
	return last
}

func hash(input [][]rune) string {
	var result strings.Builder
	result.Grow(len(input) * len(input[0]))
	for _, runes := range input {
		for _, r := range runes {
			result.WriteRune(r)
		}
	}

	return result.String()
}

func tilt(direction Direction, input [][]rune) [][]rune {
	switch direction {
	case north:
		for colId := 0; colId < len(input[0]); colId++ {
			count := 0
			lastCube := -1
			for rowId, row := range input {
				switch row[colId] {
				case cubeRock:
					if count > 0 {
						for i := lastCube + 1; i < rowId; i++ {
							if i < lastCube+1+count {
								input[i][colId] = roundRock
							} else {
								input[i][colId] = emptySpace
							}
						}
					}
					lastCube = rowId
					count = 0
				case roundRock:
					count++
				}
			}
			if count > 0 {
				for i := lastCube + 1; i < len(input); i++ {
					if i < lastCube+1+count {
						input[i][colId] = roundRock
					} else {
						input[i][colId] = emptySpace
					}
				}
			}
		}
		return input

	case west:
		for rowId, row := range input {
			count := 0
			lastCube := -1
			for colId, char := range row {
				switch char {
				case cubeRock:
					if count > 0 {
						for i := lastCube + 1; i < colId; i++ {
							if i < lastCube+1+count {
								input[rowId][i] = roundRock
							} else {
								input[rowId][i] = emptySpace
							}
						}
					}
					lastCube = colId
					count = 0
				case roundRock:
					count++
				}
			}

			if count > 0 {
				for i := lastCube + 1; i < len(input[0]); i++ {
					if i < lastCube+1+count {
						input[rowId][i] = roundRock
					} else {
						input[rowId][i] = emptySpace
					}
				}
			}
		}
		return input
	case south:
		for colId := 0; colId < len(input[0]); colId++ {
			count := 0
			lastCube := len(input)
			for rowId := len(input) - 1; rowId >= 0; rowId-- {
				row := input[rowId]
				switch row[colId] {
				case cubeRock:
					if count > 0 {
						for i := lastCube - 1; i > rowId; i-- {
							if i > lastCube-1-count {
								input[i][colId] = roundRock
							} else {
								input[i][colId] = emptySpace
							}
						}
					}
					lastCube = rowId
					count = 0
				case roundRock:
					count++
				}
			}
			if count > 0 {
				for i := lastCube - 1; i >= 0; i-- {
					if i > lastCube-1-count {
						input[i][colId] = roundRock
					} else {
						input[i][colId] = emptySpace
					}
				}
			}

		}
		return input
	case east:
		for rowId, row := range input {
			count := 0
			lastCube := len(input[0])
			for colId := len(row) - 1; colId >= 0; colId-- {
				char := row[colId]
				switch char {
				case cubeRock:
					if count > 0 {
						for i := lastCube - 1; i > colId; i-- {
							if i > lastCube-1-count {
								input[rowId][i] = roundRock
							} else {
								input[rowId][i] = emptySpace
							}
						}
					}
					lastCube = colId
					count = 0
				case roundRock:
					count++
				}
			}

			if count > 0 {
				for i := lastCube - 1; i >= 0; i-- {
					if i > lastCube-1-count {
						input[rowId][i] = roundRock
					} else {
						input[rowId][i] = emptySpace
					}
				}
			}

		}
		return input

	default:
		panic("unknown direction")
	}
}
