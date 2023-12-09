package day9

import (
	"github.com/pivovarit/aoc/util"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("mirageMaintenancePart1", func() int {
		return mirageMaintenancePart1(input)
	})
	util.Timed("mirageMaintenancePart2", func() int {
		return mirageMaintenancePart2(input)
	})
}

type History []int

func (h *History) values() []int {
	return *h
}

func mirageMaintenancePart1(input []string) int {
	histories := parse(input)
	var sum = 0
	for _, history := range histories {
		_, right := extrapolate(history)
		sum += right
	}

	return sum
}

func mirageMaintenancePart2(input []string) int {
	histories := parse(input)
	var sum = 0
	for _, history := range histories {
		left, _ := extrapolate(history)
		sum += left
	}

	return sum
}

func parse(input []string) []History {
	var histories []History
	for _, line := range input {
		var history History
		fields := strings.Fields(line)
		for _, field := range fields {
			entry, _ := strconv.Atoi(field)
			history = append(history, entry)
		}
		histories = append(histories, history)
	}
	return histories
}

func derivatives(history History) [][]int {
	var differentiates [][]int
	differentiates = append(differentiates, history.values())
	source := history.values()
	for !allZeroes(source) {
		var diffs []int
		for i := 0; i < len(source)-1; i++ {
			diffs = append(diffs, source[i+1]-source[i])
		}
		differentiates = append(differentiates, diffs)
		source = diffs
	}

	return differentiates
}

func extrapolate(history History) (int, int) {
	diffs := derivatives(history)
	diffs[len(diffs)-1] = append(diffs[len(diffs)-1], 0)
	diffs[len(diffs)-1] = prepend(diffs[len(diffs)-1], 0)

	for i := len(diffs) - 2; i >= 0; i-- {
		diffs[i] = append(diffs[i], diffs[i][len(diffs[i])-1]+diffs[i+1][len(diffs[i+1])-1])
		diffs[i] = prepend(diffs[i], diffs[i][0]-diffs[i+1][0])
	}

	return diffs[0][0], diffs[0][len(diffs[0])-1]
}

func allZeroes(slice []int) bool {
	for _, i := range slice {
		if i != 0 {
			return false
		}
	}
	return true
}

func prepend[T any](slice []T, elems ...T) []T {
	return append(elems, slice...)
}
