package day11

import (
	"github.com/pivovarit/aoc/util"
	"math"
	"slices"
)

func run() {
	input := util.ReadInput()

	util.Timed("cosmicExpansionPart1", func() int {
		return cosmicExpansionPart1(input)
	})
	util.Timed("cosmicExpansionPart2", func() int {
		return cosmicExpansionPart2(input)
	})
}

type (
	Tile     int32
	Universe [][]Tile
	Galaxy   struct{ row, col int }
)

const (
	emptySpace Tile = '.'
	galaxy     Tile = '#'
)

func cosmicExpansionPart1(input []string) int {
	universe := expand(parse(input))
	galaxies := findGalaxies(universe)
	var sum = 0
	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			if g1 != g2 {
				sum += int(math.Abs(float64(g1.row-g2.row)) + math.Abs(float64(g1.col-g2.col)))
			}
		}
	}
	return sum / 2
}

func cosmicExpansionPart2(input []string) int {
	return calculateExpandedPath(input, 1_000_000)
}

func calculateExpandedPath(input []string, n int) int {
	universe := parse(input)
	emptyRows, emptyCols := findEmpty(universe)
	galaxies := findGalaxies(universe)
	var sum = 0
	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			if g1 != g2 {
				var rowMax, rowMin = g2.row, g1.row
				var colMax, colMin = g2.col, g1.col
				if g1.row > g2.row {
					rowMax, rowMin = g1.row, g2.row
				}
				if g1.col > g2.col {
					colMax, colMin = g1.col, g2.col
				}

				emptyRowsCount := 0
				emptyColsCount := 0

				for _, r := range emptyRows {
					if r <= rowMax && r >= rowMin {
						emptyRowsCount++
					}
				}

				for _, c := range emptyCols {
					if c <= colMax && c >= colMin {
						emptyColsCount++
					}
				}

				sum += int(math.Abs(float64(g1.row-g2.row))+math.Abs(float64(g1.col-g2.col))) + ((n - 1) * emptyRowsCount) + ((n - 1) * emptyColsCount)
			}
		}
	}
	return sum / 2
}

func expand(universe Universe) Universe {
	rows, cols := findEmpty(universe)

	for idx, row := range rows {
		universe = slices.Insert(universe, row+idx+1, emptyRow(len(universe[0])))
	}

	for idx, col := range cols {
		for row := range universe {
			universe[row] = slices.Insert(universe[row], col+idx+1, emptySpace)
		}
	}

	return universe
}

func emptyRow(size int) []Tile {
	var row []Tile
	for i := 0; i < size; i++ {
		row = append(row, emptySpace)
	}
	return row
}

func findGalaxies(universe Universe) []Galaxy {
	var galaxies []Galaxy
	for r, row := range universe {
		for c, tile := range row {
			if tile == galaxy {
				galaxies = append(galaxies, Galaxy{r, c})
			}
		}
	}
	return galaxies
}

func findEmpty(universe Universe) ([]int, []int) {
	var cols []int
	var rows []int

	for i, row := range universe {
		if allEmpty(row) {
			rows = append(rows, i)
		}
	}

	for col := 0; col < len(universe[0]); col++ {
		empty := true
		for _, row := range universe {
			if row[col] != emptySpace {
				empty = false
			}
		}
		if empty {
			cols = append(cols, col)
		}
	}
	return rows, cols
}

func allEmpty(row []Tile) bool {
	for _, tile := range row {
		if tile != emptySpace {
			return false
		}
	}
	return true
}

func parse(input []string) Universe {
	var tiles Universe
	for _, line := range input {
		var tileRow []Tile
		for col := range line {
			tileRow = append(tileRow, Tile(line[col]))
		}
		tiles = append(tiles, tileRow)
	}
	return tiles
}
