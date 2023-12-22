package day18

import (
	"github.com/pivovarit/aoc/util"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("lavaductLagoonPart1", func() int {
		return lavaductLagoonPart1(input)
	})
	util.Timed("lavaductLagoonPart2", func() int {
		return lavaductLagoonPart2(input)
	})
}

const (
	up    = 'U'
	down  = 'D'
	left  = 'L'
	right = 'R'
)

type Point struct{ row, col int }

func (p *Point) left(amount int) Point {
	return Point{p.row, p.col - amount}
}

func (p *Point) right(amount int) Point {
	return Point{p.row, p.col + amount}
}

func (p *Point) down(amount int) Point {
	return Point{p.row + amount, p.col}
}

func (p *Point) up(amount int) Point {
	return Point{p.row - amount, p.col}
}

func lavaductLagoonPart1(input []string) int {
	last := Point{0, 0}
	var points = []Point{last}
	var distance = 0

	for _, line := range input {
		split := strings.Fields(line)
		meters, _ := strconv.Atoi(split[1])
		// 0 means R, 1 means D, 2 means L, and 3 means U
		distance += meters
		switch int32(split[0][0]) {
		case up:
			last = last.up(meters)
		case down:
			last = last.down(meters)
		case left:
			last = last.left(meters)
		case right:
			last = last.right(meters)
		default:
			panic("unknown direction")
		}
		points = append(points, last)
	}

	return shoelace(points) + distance/2 + 1
}

func shoelace(points []Point) int {
	sum := 0
	last := points[len(points)-1]
	for _, point := range points {
		sum += point.row*last.col - point.col*last.row
		last = point
	}
	return int(math.Abs(float64(sum))) / 2
}

func lavaductLagoonPart2(input []string) int {
	last := Point{0, 0}
	var points = []Point{last}
	var distance = 0

	for _, line := range input {
		split := strings.Fields(line)
		colour := split[2][1 : len(split[2])-1]
		meters := parseMeters(colour)
		distance += meters

		switch colour[len(colour)-1] {
		case '0':
			last = last.right(meters)
		case '1':
			last = last.down(meters)
		case '2':
			last = last.left(meters)
		case '3':
			last = last.up(meters)
		default:
			panic("unknown direction")
		}
		points = append(points, last)
	}
	return shoelace(points) + distance/2 + 1
}

var bi = new(big.Int)

func parseMeters(colour string) int {
	bi.SetString(colour[1:len(colour)-1], 16)
	return int(bi.Uint64())
}
