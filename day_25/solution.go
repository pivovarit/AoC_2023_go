package day25

import (
	"github.com/pivovarit/aoc/util"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("snowverloadPart1", func() int {
		return snowverloadPart1(input)
	})
	util.Timed("snowverloadPart2", func() int {
		return snowverloadPart2(input)
	})
}

func snowverloadPart1(input []string) int {
	return 0
}

func snowverloadPart2(input []string) int {
	return 0
}

type Component struct {
	name      string
	connected []string
}

func parse(input []string) (components []Component) {
	for _, line := range input {
		split := strings.Split(line, ":")
		name := split[0]
		connected := strings.Fields(split[1])
		components = append(components, Component{name, connected})
	}
	return components
}
