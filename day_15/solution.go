package day15

import (
	"github.com/pivovarit/aoc/util"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("lensLibraryPart1", func() int {
		return lensLibraryPart1(input)
	})
	util.Timed("lensLibraryPart2", func() int {
		return lensLibraryPart2(input)
	})
}

const (
	opRemove = '-'
	opAssign = '='
)

func lensLibraryPart1(input []string) int {
	var sum int32
	var current int32
	for _, char := range input[0] {
		if char == ',' {
			sum += current
			current = 0
		} else {
			current = ((current + char) * 17) % 256
		}
	}
	sum += current

	return int(sum)
}

func hash(input string) int {
	var current int32

	for _, char := range input {
		current = ((current + char) * 17) % 256
	}
	return int(current)
}

type Lens struct {
	id    string
	focal string
}

func lensLibraryPart2(input []string) int {
	var boxes = make([][]Lens, 256)
	for _, seq := range strings.Split(input[0], ",") {
		separatorIdx, separator := op(seq)
		lensId := seq[:separatorIdx]
		focal := seq[separatorIdx+1:]
		boxId := hash(lensId)
		switch separator {
		case opAssign:
			updated := false
			for i, s := range boxes[boxId] {
				if s.id == lensId {
					boxes[boxId][i] = Lens{lensId, focal}
					updated = true
					break
				}
			}
			if !updated {
				boxes[boxId] = append(boxes[boxId], Lens{lensId, focal})
			}
		case opRemove:
			for i, s := range boxes[boxId] {
				if s.id == lensId {
					boxes[boxId] = append(append(make([]Lens, 0), boxes[boxId][:i]...), boxes[boxId][i+1:]...)
					break
				}
			}
		}
	}

	var sum = 0

	for boxIdx, lenses := range boxes {
		if len(lenses) > 0 {
			for slotIdx, lens := range lenses {
				focalLength, _ := strconv.Atoi(lens.focal)
				sum += (boxIdx + 1) * (slotIdx + 1) * focalLength
			}
		}
	}
	return sum
}

func op(seq string) (int, int32) {
	for i, char := range seq {
		if char == opAssign || char == opRemove {
			return i, char
		}
	}
	return -1, 0
}
