package day8

import (
	"github.com/pivovarit/aoc/util"
	"strings"
)

const (
	left  = 'L'
	right = 'R'
)

func run() {
	input := util.ReadInput()

	util.Timed("hauntedWasteLandPart1", func() int {
		return hauntedWasteLandPart1(input)
	})
	util.Timed("hauntedWasteLandPart2", func() int {
		return hauntedWasteLandPart2(input)
	})
}

func hauntedWasteLandPart1(input []string) int {
	instructions, directions := parse(input)
	return pathLength("AAA", instructions, directions, func(s string) bool { return s == ("ZZZ") })
}

func hauntedWasteLandPart2(input []string) int {
	instructions, directions := parse(input)
	startNodes := getStartNodes(directions)
	var paths = make([]int, len(startNodes))
	for i, node := range startNodes {
		paths[i] = pathLength(node, instructions, directions, func(s string) bool { return strings.HasSuffix(s, "Z") })
	}
	return lcm(paths)
}

func pathLength(start, instructions string, directions map[string]Directions, isEnd func(string) bool) int {
	steps := 0
	now := start

	for i := 0; !isEnd(now); i++ {
		steps++
		if instructions[i%len(instructions)] == left {
			now = directions[now].left
		} else if instructions[i%len(instructions)] == right {
			now = directions[now].right
		} else {
			panic("illegal direction value")
		}
	}
	return steps
}

func getStartNodes(directions map[string]Directions) []string {
	var startNodes []string

	for node := range directions {
		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
		}
	}
	return startNodes
}

func parse(input []string) (string, map[string]Directions) {
	instructions := input[0]
	var roadMap = make(map[string]Directions)

	for _, line := range input[2:] {
		split := strings.Split(line, "=")
		from := strings.TrimSpace(split[0])
		directionsLine := strings.ReplaceAll(strings.TrimSpace(split[1])[1:], ")", "")
		directions := strings.Split(directionsLine, ",")
		roadMap[from] = Directions{strings.TrimSpace(directions[0]), strings.TrimSpace(directions[1])}
	}

	return instructions, roadMap
}

type Directions struct {
	left, right string
}

func lcm(numbers []int) int {
	result := 1
	for _, x := range numbers {
		gcd := result
		b := x
		for b != 0 {
			t := b
			b = gcd % b
			gcd = t
		}
		result = result / gcd * x
	}
	return result
}
