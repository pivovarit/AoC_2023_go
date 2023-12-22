package day22

import (
	"fmt"
	"github.com/pivovarit/aoc/util"
	"sort"
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
		id       int
		from, to Coordinates
	}
)

func sandSlabsPart1(input []string) int {
	cubes := getSettled(parse(input))
	count := 0
	for idx := 0; idx < len(cubes); idx++ {
		copiedSlice := make([]Cube, len(cubes))
		copy(copiedSlice, cubes)
		copiedSlice = append(copiedSlice[:idx], copiedSlice[idx+1:]...)
		newCubeList := make([]Cube, len(copiedSlice))
		copy(newCubeList, copiedSlice)
		newCubeList = getSettled(newCubeList)
		diff := false
		for i := 0; i < len(copiedSlice); i++ {
			for j := 0; j < len(newCubeList); j++ {
				if copiedSlice[i].id == newCubeList[j].id {
					if copiedSlice[i].from != newCubeList[j].from {
						diff = true
					}
					break
				}
			}
		}
		if !diff {
			count++
		}
	}
	return count
}

func sandSlabsPart2(input []string) int {
	cubes := getSettled(parse(input))
	diff := 0
	for idx := 0; idx < len(cubes); idx++ {
		copiedSlice := make([]Cube, len(cubes))
		copy(copiedSlice, cubes)
		copiedSlice = append(copiedSlice[:idx], copiedSlice[idx+1:]...)
		newCubeList := make([]Cube, len(copiedSlice))
		copy(newCubeList, copiedSlice)
		newCubeList = getSettled(newCubeList)
		for i := 0; i < len(copiedSlice); i++ {
			for j := 0; j < len(newCubeList); j++ {
				if copiedSlice[i].id == newCubeList[j].id {
					if copiedSlice[i].from != newCubeList[j].from {
						diff++
					}
					break
				}
			}
		}
	}
	return diff
}

func getSettled(cubes []Cube) []Cube {
	for {
		finished := true
		sort.Slice(cubes, func(i, j int) bool {
			return cubes[i].from.z < cubes[j].from.z
		})
		for idx, brick := range cubes {
			if brick.from.z > 0 {
				var touch bool
				for j := 0; j < idx; j++ {
					touch = false
					if cubes[j].to.z == brick.from.z-1 {
						if brick.from.x <= cubes[j].to.x && brick.to.x >= cubes[j].from.x &&
							brick.from.y <= cubes[j].to.y && brick.to.y >= cubes[j].from.y {
							touch = true
							break
						}
					}
				}
				if !touch {
					finished = false
					cubes[idx].from.z--
					cubes[idx].to.z--
				}
			}
		}
		if finished {
			break
		}
	}
	return cubes
}

func parse(input []string) (cubes []Cube) {
	for idx, line := range input {
		split := strings.Split(line, "~")
		from := strings.Split(split[0], ",")
		to := strings.Split(split[1], ",")

		cubes = append(cubes, Cube{
			from: Coordinates{asInt(from[0]), asInt(from[1]), asInt(from[2])},
			to:   Coordinates{asInt(to[0]), asInt(to[1]), asInt(to[2])},
			id:   idx,
		})
	}
	return cubes
}

func asInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("str: %+v should be an int\n", str))
	}
	return result
}
