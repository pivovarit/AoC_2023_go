package day13

import (
	"github.com/pivovarit/aoc/util"
)

func run() {
	input := util.ReadInput()

	util.Timed("pointOfIncidencePart1", func() int {
		return pointOfIncidencePart1(input)
	})
	util.Timed("pointOfIncidencePart2", func() int {
		return pointOfIncidencePart2(input)
	})
}

func pointOfIncidencePart1(input []string) int {
	return pointOfIncidence(input,
		func(pattern Pattern) (bool, int) { return findHorizontalReflection(pattern) },
		func(pattern Pattern) (bool, int) { return findVerticalReflection(pattern) })
}

func pointOfIncidencePart2(input []string) int {
	return pointOfIncidence(input,
		func(pattern Pattern) (bool, int) { return findHorizontalReflectionWithSmudge(pattern) },
		func(pattern Pattern) (bool, int) { return findVerticalReflectionWithSmudge(pattern) })
}

func pointOfIncidence(input []string, horizontalReflection func(Pattern) (bool, int), verticalReflection func(Pattern) (bool, int)) int {
	parser := Parser{Tokenizer{input: input}}
	patterns := parser.patterns()

	var sum = 0

	for _, pattern := range patterns {
		found, count := horizontalReflection(pattern)
		if found {
			sum += count * 100
		}
		found, count = verticalReflection(pattern)
		if found {
			sum += count
		}
	}

	return sum
}

func rowDifferences(r1 int, r2 int, pattern Pattern) int {
	diffs := 0
	for i := 0; i < len(pattern[r1]); i++ {
		if pattern[r1][i] != pattern[r2][i] {
			diffs++
		}
	}
	return diffs
}

func colDifferences(c1 int, c2 int, pattern Pattern) int {
	diffs := 0
	for _, row := range pattern {
		if row[c1] != row[c2] {
			diffs++
		}
	}
	return diffs
}

func colEquals(c1 int, c2 int, pattern Pattern) bool {
	for _, row := range pattern {
		if row[c1] != row[c2] {
			return false
		}
	}
	return true
}

func findVerticalReflection(pattern Pattern) (bool, int) {
	for colId := 0; colId < len(pattern[0])-1; colId++ {
		isRef := true
		for j := colId; j >= 0; j-- {
			if colId+colId-j+1 < len(pattern[0]) {
				if !colEquals(colId+colId-j+1, j, pattern) {
					isRef = false
					break
				}
			}
		}
		if !isRef {
			continue
		}
		result := colId + 1
		return true, result
	}
	return false, -1
}

func findHorizontalReflection(pattern Pattern) (bool, int) {
	for rowId := 0; rowId < len(pattern)-1; rowId++ {
		isRef := true
		for j := rowId; j >= 0; j-- {
			if rowId+rowId-j+1 < len(pattern) {
				if pattern[rowId+rowId-j+1] != pattern[j] {
					isRef = false
					break
				}
			}
		}
		if !isRef {
			continue
		}
		result := rowId + 1
		return true, result
	}
	return false, -1
}

func findVerticalReflectionWithSmudge(pattern Pattern) (bool, int) {
	for colId := 0; colId < len(pattern[0])-1; colId++ {
		smudges := 0
		isRef := true
		for j := colId; j >= 0; j-- {
			if colId+colId-j+1 < len(pattern[0]) {
				diffs := colDifferences(colId+colId-j+1, j, pattern)
				if diffs > 1 || (diffs > 0 && smudges == 1) {
					isRef = false
					break
				} else if diffs == 1 {
					smudges++
				}
			}
		}
		if !isRef {
			continue
		}
		if smudges == 1 {
			result := colId + 1
			return true, result
		}
	}
	return false, -1
}

func findHorizontalReflectionWithSmudge(pattern Pattern) (bool, int) {
	for rowId := 0; rowId < len(pattern)-1; rowId++ {
		smudges := 0
		isRef := true
		for j := rowId; j >= 0; j-- {
			if rowId+rowId-j+1 < len(pattern) {
				diffs := rowDifferences(rowId+rowId-j+1, j, pattern)
				if diffs > 1 || (diffs > 0 && smudges == 1) {
					isRef = false
					break
				} else if diffs == 1 {
					smudges++
				}
			}
		}
		if !isRef {
			continue
		}
		if smudges == 1 {
			result := rowId + 1
			return true, result
		}
	}
	return false, -1
}

type (
	Pattern []string
	Parser  struct {
		tokenizer Tokenizer
	}
	Tokenizer struct {
		input   []string
		lastIdx int
	}
)

func (p *Parser) patterns() []Pattern {
	var patterns []Pattern
	for p.tokenizer.hasNext() {
		patterns = append(patterns, p.tokenizer.next())
	}
	return patterns
}

func (t *Tokenizer) hasNext() bool {
	return t.lastIdx < len(t.input)-1
}

func (t *Tokenizer) next() []string {
	var next []string
	for i := t.lastIdx; i < len(t.input)-1; i++ {
		if len(t.input[i]) == 0 {
			next = t.input[t.lastIdx:i]
			t.lastIdx = i + 1
			return next
		}
	}

	result := t.input[t.lastIdx:]
	t.lastIdx = len(t.input) - 1
	return result
}
