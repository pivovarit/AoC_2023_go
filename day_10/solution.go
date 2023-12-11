package day10

import (
	"fmt"
	"github.com/pivovarit/aoc/util"
)

func run() {
	input := util.ReadInput()

	util.Timed("pipeMazePart1", func() int {
		return pipeMazePart1(input) // 6860
	})
	util.Timed("pipeMazePart2", func() int {
		return pipeMazePart2(input)
	})

	visualize(input)
}

func visualize(input []string) {
	board := parse(input)
	starting := findStart(board)
	pipe, direction := findPipe(board, starting)
	_, pathMap := followPipe(pipe, board, direction, []Location{pipe}, initPathMap(board))
	sanitize(board, pathMap)
	replaceStart(board, starting)
	prettyPrint(board)
}

func pipeMazePart1(input []string) int {
	board := parse(input)
	pipe, direction := findPipe(board, findStart(board))
	path, _ := followPipe(pipe, board, direction, []Location{pipe}, initPathMap(board))

	return len(path) / 2
}

func initPathMap(board Board) [][]bool {
	var pathMap = make([][]bool, len(board))
	for r := range board {
		pathMap[r] = make([]bool, len(board[r]))
	}
	return pathMap
}

func pipeMazePart2(input []string) int {
	board := parse(input)
	str := findStart(board)
	pipe, direction := findPipe(board, str)
	_, pathMap := followPipe(pipe, board, direction, []Location{pipe}, initPathMap(board))
	sanitize(board, pathMap)
	replaceStart(board, str)

	var sum = 0
	for _, row := range board {
		inside := false
		for _, sign := range row {
			if sign == verticalPipe || sign == southWestPipe || sign == southEastPipe {
				inside = !inside
			} else {
				if inside {
					if !isPipe(sign) {
						sum++
					}
				}
			}
		}
	}

	return sum
}

func replaceStart(board Board, str Location) {
	if isPipeAt(board, str.right()) && getPipe(board, str.right()).canBeEnteredFrom(west) {
		if isPipeAt(board, str.left()) && getPipe(board, str.left()).canBeEnteredFrom(east) {
			board[str.row][str.col] = horizontalPipe
		}
	}

	if isPipeAt(board, str.left()) && getPipe(board, str.left()).canBeEnteredFrom(east) {
		if isPipeAt(board, str.up()) && getPipe(board, str.up()).canBeEnteredFrom(south) {
			board[str.row][str.col] = northWestPipe
		}

		if isPipeAt(board, str.down()) && getPipe(board, str.down()).canBeEnteredFrom(north) {
			board[str.row][str.col] = southWestPipe
		}
	}

	if isPipeAt(board, str.right()) && getPipe(board, str.right()).canBeEnteredFrom(west) {
		if isPipeAt(board, str.up()) && getPipe(board, str.up()).canBeEnteredFrom(south) {
			board[str.row][str.col] = northEastPipe
		}

		if isPipeAt(board, str.down()) && getPipe(board, str.down()).canBeEnteredFrom(north) {
			board[str.row][str.col] = southEastPipe
		}
	}
}

const start = 'S'

const (
	verticalPipe   = '|'
	horizontalPipe = '-'
	northEastPipe  = 'L'
	northWestPipe  = 'J'
	southWestPipe  = '7'
	southEastPipe  = 'F'
)

func sanitize(board Board, pathMap [][]bool) Board {
	for r, row := range board {
		for c := range row {
			if !pathMap[r][c] {
				if board[r][c] != 'I' {
					board[r][c] = '.'
				}
			}
		}
	}
	return board
}

func findPipe(board Board, str Location) (Location, Direction) {
	if isPipeAt(board, str.up()) && getPipe(board, str.up()).canBeEnteredFrom(south) {
		return str.up(), south
	} else if isPipeAt(board, str.right()) && getPipe(board, str.right()).canBeEnteredFrom(west) {
		return str.right(), west
	} else if isPipeAt(board, str.down()) && getPipe(board, str.down()).canBeEnteredFrom(north) {
		return str.down(), north
	} else if isPipeAt(board, str.left()) && getPipe(board, str.left()).canBeEnteredFrom(east) {
		return str.left(), east
	}

	panic("could not find any pipes, can't proceed!")
}

func followPipe(location Location, board Board, direction Direction, acc []Location, pathMap [][]bool) ([]Location, [][]bool) {
	pathMap[location.row][location.col] = true
	if isPipeAt(board, location) {
		pipe := getPipe(board, location)
		if pipe == start {
			return acc, pathMap
		}
		next, nextDirection := pipe.follow(location, direction)
		return followPipe(next, board, nextDirection, append(acc, next), pathMap)
	} else {
		return acc, pathMap
	}
}

func findStart(board Board) Location {
	for rowIdx, row := range board {
		for colIdx, sign := range row {
			if sign == start {
				return Location{row: rowIdx, col: colIdx}
			}
		}
	}
	panic("no 'S' on the board")
}

func parse(input []string) Board {
	var result [][]int32

	for _, line := range input {
		var next []int32
		for _, sign := range line {
			next = append(next, sign)
		}
		result = append(result, next)
	}

	return result
}

type (
	Direction int
	Board     [][]int32
	Location  struct {
		row, col int
	}
)

func (l Location) left() Location {
	return Location{l.row, l.col - 1}
}

func (l Location) right() Location {
	return Location{l.row, l.col + 1}
}

func (l Location) up() Location {
	return Location{l.row - 1, l.col}
}

func (l Location) down() Location {
	return Location{l.row + 1, l.col}
}
func getPipe(board Board, loc Location) Pipe {
	return Pipe(board[loc.row][loc.col])
}

func isPipe(sign int32) bool {
	switch sign {
	case verticalPipe:
		return true
	case horizontalPipe:
		return true
	case northEastPipe:
		return true
	case northWestPipe:
		return true
	case southWestPipe:
		return true
	case southEastPipe:
		return true
	default:
		return false
	}
}
func isPipeAt(board Board, loc Location) bool {
	rows := len(board)
	cols := len(board[0])

	if loc.row >= 0 && loc.row < rows {
		if loc.col >= 0 && loc.col < cols {
			return isPipe(board[loc.row][loc.col])
		}
	}
	return false
}

const (
	east Direction = iota
	west
	south
	north
)

type Pipe byte

func (p Pipe) canBeEnteredFrom(dir Direction) bool {
	switch p {
	case verticalPipe:
		return dir == north || dir == south
	case horizontalPipe:
		return dir == west || dir == east
	case northEastPipe:
		return dir == north || dir == east
	case northWestPipe:
		return dir == north || dir == west
	case southWestPipe:
		return dir == south || dir == west
	case southEastPipe:
		return dir == south || dir == east
	default:
		panic("illegal pipe value")
	}
}

func (p Pipe) follow(location Location, entry Direction) (Location, Direction) {
	switch p {
	case verticalPipe:
		switch entry {
		case south:
			return location.up(), south
		case north:
			return location.down(), north
		default:
			panic("illegal pipe entry")
		}

	case horizontalPipe:
		switch entry {
		case east:
			return location.left(), east
		case west:
			return location.right(), west
		default:
			panic("illegal pipe entry")
		}
	case northEastPipe:
		switch entry {
		case north:
			return location.right(), west
		case east:
			return location.up(), south
		default:
			panic("illegal pipe entry")
		}
	case northWestPipe:
		switch entry {
		case north:
			return location.left(), east
		case west:
			return location.up(), south
		default:
			panic("illegal pipe entry")
		}
	case southWestPipe:
		switch entry {
		case south:
			return location.left(), east
		case west:
			return location.down(), north
		default:
			panic("illegal pipe entry")
		}
	case southEastPipe:
		switch entry {
		case south:
			return location.right(), west
		case east:
			return location.down(), north
		default:
			panic("illegal pipe entry")
		}
	default:
		panic("this is not a pipe!")
	}
}

func prettyPrint(board Board) {
	var newBoard = make(Board, len(board))

	for i, row := range board {
		newBoard[i] = make([]int32, len(board[i]))
		copy(newBoard[i], row)
	}
	for i, row := range newBoard {
		for j, sign := range row {
			switch sign {
			case '.':
				newBoard[i][j] = ' '
			case '|':
				newBoard[i][j] = '║'
			case '-':
				newBoard[i][j] = '═'
			case 'L':
				newBoard[i][j] = '╚'
			case 'J':
				newBoard[i][j] = '╝'
			case '7':
				newBoard[i][j] = '╗'
			case 'F':
				newBoard[i][j] = '╔'
			}
		}
	}

	for _, row := range newBoard {
		for _, sign := range row {
			fmt.Printf("%c", sign)
		}
		println()

	}
}
