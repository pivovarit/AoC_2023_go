package day24

import (
	"fmt"
	"github.com/pivovarit/aoc/util"
	"sort"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("neverTellMeTheOddsPart1", func() int {
		return neverTellMeTheOddsPart1(input)
	})
	util.Timed("neverTellMeTheOddsPart2", func() int {
		return neverTellMeTheOddsPart2(input)
	})
}

func neverTellMeTheOddsPart1(input []string) (count int) {
	hailstones := parse(input)

	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			h1 := hailstones[i]
			h2 := hailstones[j]

			if intersection, exists := intersect(h1, h2); exists {
				if intersection.x < float64(minAreaBound) || intersection.x > float64(maxAreaBound) ||
					intersection.y < float64(minAreaBound) || intersection.y > float64(maxAreaBound) {
					continue
				}

				if getTime(h1, intersection.x) >= 0 && getTime(h2, intersection.x) >= 0 {
					count++
				}
			}
		}
	}

	return count
}

func neverTellMeTheOddsPart2(input []string) int {
	hailstones, velocitiesX, velocitiesY, velocitiesZ := parseWithVelocities(input)

	potentialVelocity := make([]int, 2001)
	for x := -1000; x <= 1000; x++ {
		potentialVelocity[x+1000] = x
	}

	rvx := getRockVelocity(velocitiesX)
	rvy := getRockVelocity(velocitiesY)
	rvz := getRockVelocity(velocitiesZ)

	results := make(map[int]int)
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			stoneA := hailstones[i]
			stoneB := hailstones[j]

			ma := float64(stoneA.velocity.y-rvy) / float64(stoneA.velocity.x-rvx)
			mb := float64(stoneB.velocity.y-rvy) / float64(stoneB.velocity.x-rvx)

			ca := float64(stoneA.location.y) - ma*float64(stoneA.location.x)
			cb := float64(stoneB.location.y) - mb*float64(stoneB.location.x)

			rpx := int((cb - ca) / (ma - mb))
			rpy := int(ma*float64(rpx) + ca)

			time := (rpx - stoneA.location.x) / int(float64(stoneA.velocity.x-rvx))
			rpz := stoneA.location.z + (stoneA.velocity.z-rvz)*time

			result := rpx + rpy + rpz
			if _, ok := results[result]; !ok {
				results[result] = 1
			} else {
				results[result]++
			}
		}
	}

	var keys []int
	for k := range results {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return results[keys[i]] > results[keys[j]]
	})
	return keys[0]
}

const (
	minAreaBound = 200000000000000
	maxAreaBound = 400000000000000
)

type (
	Location struct {
		x, y, z int
	}
	Velocity struct {
		x, y, z int
	}
	Intersection struct {
		x, y float64
	}
	Hailstone struct {
		location Location
		velocity Velocity
	}
)

func intersect(h1, h2 Hailstone) (Intersection, bool) {
	x1, x2, y1, y2 := float64(h1.location.x), float64(h1.location.x+100000000000000*h1.velocity.x), float64(h1.location.y), float64(h1.location.y+100000000000000*h1.velocity.y)
	x3, x4, y3, y4 := float64(h2.location.x), float64(h2.location.x+100000000000000*h2.velocity.x), float64(h2.location.y), float64(h2.location.y+100000000000000*h2.velocity.y)

	denominator := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if denominator == 0 {
		return Intersection{}, false
	}

	return Intersection{
		x: ((x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)) / denominator,
		y: ((x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)) / denominator,
	}, true
}

func parse(input []string) (stones []Hailstone) {
	for _, line := range input {
		parts := strings.Split(line, " @ ")
		position := strings.Split(parts[0], ", ")
		velocity := strings.Split(parts[1], ", ")

		stones = append(stones, Hailstone{
			Location{asInt(position[0]), asInt(position[1]), asInt(position[2])},
			Velocity{asInt(velocity[0]), asInt(velocity[1]), asInt(velocity[2])},
		})
	}
	return stones
}

func getTime(stone Hailstone, p float64) float64 {
	return (p - float64(stone.location.x)) / float64(stone.velocity.x)
}

func asInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("str: %+v should be an int\n", str))
	}
	return result
}

func getRockVelocity(velocities map[int][]int) int {
	possibleV := make([]int, 0)
	for x := -1000; x <= 1000; x++ {
		possibleV = append(possibleV, x)
	}

	for vel, values := range velocities {
		if len(values) < 2 {
			continue
		}

		newPossibleV := make([]int, 0)
		for _, possible := range possibleV {
			if possible-vel != 0 && (values[0]-values[1])%(possible-vel) == 0 {
				newPossibleV = append(newPossibleV, possible)
			}
		}

		possibleV = newPossibleV
	}

	return possibleV[0]
}

func parseWithVelocities(input []string) ([]Hailstone, map[int][]int, map[int][]int, map[int][]int) {
	var hailstones []Hailstone
	var velocitiesX = make(map[int][]int)
	var velocitiesY = make(map[int][]int)
	var velocitiesZ = make(map[int][]int)
	for _, line := range input {
		parts := strings.Split(line, " @ ")
		location := strings.Split(parts[0], ", ")
		x := asInt(location[0])
		y := asInt(location[1])
		z := asInt(location[2])

		velocity := strings.Split(parts[1], ", ")
		vx := asInt(velocity[0])
		vy := asInt(velocity[1])
		vz := asInt(velocity[2])

		if v, ok := velocitiesX[vx]; ok {
			velocitiesX[vx] = append(v, x)
		} else {
			velocitiesX[vx] = []int{x}
		}

		if v, ok := velocitiesY[vy]; ok {
			velocitiesY[vy] = append(v, y)
		} else {
			velocitiesY[vy] = []int{y}
		}

		if v, ok := velocitiesZ[vz]; ok {
			velocitiesZ[vz] = append(v, z)
		} else {
			velocitiesZ[vz] = []int{z}
		}

		hailstones = append(hailstones, Hailstone{Location{x, y, z}, Velocity{vx, vy, vz}})
	}
	return hailstones, velocitiesX, velocitiesY, velocitiesZ
}
