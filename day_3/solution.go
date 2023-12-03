package day3

import (
	"github.com/pivovarit/aoc/util"
	"math"
	"regexp"
	"strconv"
	"unicode"
)

var partNumberPattern = regexp.MustCompile(`\d+`)

func run() {
	input := util.ReadInput()

	println(gearRatiosPart1(input))
	println(gearRatiosPart2(input))
}

func gearRatiosPart1(input []string) int {
	var sum = 0

	for i, parts := range extractPartsArray(input) {
		for _, part := range parts {
			if isLegit(part, input, i-1) ||
				isLegit(part, input, i+1) ||
				isLegitAtIndex(input[i], part.start-1) ||
				isLegitAtIndex(input[i], part.end) {
				sum += part.number
			}
		}
	}

	return sum
}

func findGears(part GearPart, input []string) []GearLocation {
	var gears []GearLocation
	found, row, col := findGearInRow(part, input, part.row-1)
	if found {
		gears = append(gears, GearLocation{
			row: row,
			col: col,
		})
	}

	found, row, col = findGearInRow(part, input, part.row+1)
	if found {
		gears = append(gears, GearLocation{
			row: row,
			col: col,
		})
	}

	found, col = findGearAtIndex(input[part.row], part.start-1)
	if found {
		gears = append(gears, GearLocation{
			row: part.row,
			col: col,
		})
	}

	found, col = findGearAtIndex(input[part.row], part.end)
	if found {
		gears = append(gears, GearLocation{
			row: part.row,
			col: col,
		})
	}

	return gears
}

func gearRatiosPart2(input []string) int {
	var sum = 0

	var partsByGear = make(map[GearLocation][]GearPart)

	for _, parts := range extractPartsArray(input) {
		for _, part := range parts {
			gears := findGears(part, input)
			for _, gear := range gears {
				addGear(partsByGear[gear], partsByGear, gear, part)
			}
		}
	}

	for _, parts := range partsByGear {
		if len(parts) == 2 {
			sum += parts[0].number * parts[1].number
		}
	}

	return sum
}

func addGear(gearParts []GearPart, partsByGear map[GearLocation][]GearPart, gear GearLocation, part GearPart) {
	if gearParts == nil {
		partsByGear[gear] = []GearPart{part}
	} else {
		partsByGear[gear] = append(partsByGear[gear], part)
	}
}

func isSymbol(char uint8) bool {
	return char != '.' && !unicode.IsDigit(rune(char))
}

func extractParts(entry string, row int) []GearPart {
	matches := partNumberPattern.FindAllStringIndex(entry, -1)
	var parts []GearPart
	for _, match := range matches {
		part, _ := strconv.Atoi(entry[match[0]:match[1]])
		parts = append(parts, GearPart{
			row:    row,
			number: part,
			start:  match[0],
			end:    match[1],
		})
	}

	return parts
}

func extractPartsArray(entries []string) [][]GearPart {
	var parts [][]GearPart
	for row, entry := range entries {
		parts = append(parts, extractParts(entry, row))
	}
	return parts
}

func findGearInRow(part GearPart, input []string, index int) (bool, int, int) {
	if index >= 0 && index < len(input) {
		start := int(math.Max(float64(part.start-1), 0))
		end := int(math.Min(float64(part.end), float64(len(input[index])-1)))

		for charIdx, char := range input[index][start : end+1] {
			if char == '*' {
				return true, index, charIdx + start
			}
		}
	}

	return false, 0, 0
}

func isLegit(part GearPart, input []string, index int) bool {
	if index >= 0 && index < len(input) {
		start := int(math.Max(float64(part.start-1), 0))
		end := int(math.Min(float64(part.end), float64(len(input[index])-1)))

		for _, char := range input[index][start : end+1] {
			if isSymbol(uint8(char)) {
				return true
			}
		}
	}

	return false
}

func findGearAtIndex(input string, index int) (bool, int) {
	if index >= 0 && index < len(input) {
		if isSymbol(input[index]) {
			return true, index
		}
	}

	return false, 0
}

func isLegitAtIndex(input string, index int) bool {
	if index >= 0 && index < len(input) {
		if isSymbol(input[index]) {
			return true
		}
	}

	return false
}

type GearLocation struct {
	row int
	col int
}

type GearPart struct {
	number int
	row    int
	start  int
	end    int // exclusive
}
