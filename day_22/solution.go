package day22

import (
	"fmt"
	"github.com/pivovarit/aoc/util"
	"slices"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("sandSlabsPart1", func() int {
		return sandSlabsPart1(input)
	})
	util.Timed("sandSlabsPart2", func() int {
		return sandSlabsPart2(input)
	})
}

type (
	Coordinates struct {
		x, y, z int
	}
	Cube struct {
		from, to Coordinates
	}
)

func sandSlabsPart1(input []string) int {
	cubes := parse(input)
	settleAll(cubes)
	return disintegrateCounting(cubes)
}

func sandSlabsPart2(input []string) int {
	cubes := parse(input)
	settleAll(cubes)
	return disintegrateSumming(cubes)
}

func parse(input []string) (cubes []Cube) {
	for _, line := range input {
		split := strings.Split(line, "~")
		from := strings.Split(split[0], ",")
		to := strings.Split(split[1], ",")

		cubes = append(cubes, Cube{
			from: Coordinates{asInt(from[0]), asInt(from[1]), asInt(from[2])},
			to:   Coordinates{asInt(to[0]), asInt(to[1]), asInt(to[2])},
		})
	}
	slices.SortFunc(cubes, func(a, b Cube) int {
		return a.from.z - b.from.z
	})
	return cubes
}

func asInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("str: %+v should be an int\n", str))
	}
	return result
}

func settle(xy [][]int, from, to Coordinates) (int, int) {
	z1, z2 := 0, 0

	if from.x == to.x && from.z == to.z {
		for i := from.y; i <= to.y; i++ {
			if xy[from.x][i] > z1 {
				z1 = xy[from.x][i]
			}
		}

		z1, z2 = z1+1, z1+1

		for i := from.y; i <= to.y; i++ {
			xy[from.x][i] = z1
		}

	} else if from.y == to.y && from.z == to.z {
		for i := from.x; i <= to.x; i++ {
			if xy[i][from.y] > z1 {
				z1 = xy[i][from.y]
			}
		}

		z1, z2 = z1+1, z1+1

		for i := from.x; i <= to.x; i++ {
			xy[i][from.y] = z1
		}

	} else if from.x == to.x && from.y == to.y {
		z1 = xy[from.x][from.y] + 1
		z2 = z1 + to.z - from.z
		xy[from.x][from.y] = z2
	}

	return z1, z2
}

func disintegrateCounting(blocks []Cube) (count int) {
	for i := range blocks {
		newBlocks := make([]Cube, len(blocks))
		copy(newBlocks, blocks)
		if i == 0 {
			newBlocks = newBlocks[1:]
		} else if i == len(newBlocks)-1 {
			newBlocks = blocks[:len(newBlocks)-1]
		} else {
			newBlocks = append(newBlocks[:i], newBlocks[i+1:]...)
		}

		if settleAll(newBlocks) == 0 {
			count++
		}
	}
	return count
}

func disintegrateSumming(blocks []Cube) (sum int) {
	for i := range blocks {
		newBlocks := make([]Cube, len(blocks))
		copy(newBlocks, blocks)
		if i == 0 {
			newBlocks = newBlocks[1:]
		} else if i == len(newBlocks)-1 {
			newBlocks = blocks[:len(newBlocks)-1]
		} else {
			newBlocks = append(newBlocks[:i], newBlocks[i+1:]...)
		}

		sum += settleAll(newBlocks)
	}
	return sum
}

func settleAll(blocks []Cube) (count int) {
	xy := make([][]int, 10)
	for i := range xy {
		xy[i] = make([]int, 10)
	}

	for i, b := range blocks {
		b.from.z, b.to.z = settle(xy, b.from, b.to)
		if b.from != blocks[i].from || b.to != blocks[i].to {
			blocks[i] = Cube{b.from, b.to}
			count++
		}
	}
	return count
}
