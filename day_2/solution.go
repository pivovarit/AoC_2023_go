package day2

import (
	"github.com/pivovarit/aoc/util"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	BLUE  = "blue"
	GREEN = "green"
	RED   = "red"
)

func run() {
	input := util.ReadInput()

	util.Timed("cubeConundrumPart1", func() int {
		return cubeConundrumPart1(input)
	})
	util.Timed("cubeConundrumPart1", func() int {
		return cubeConundrumPart2(input)
	})
}

func cubeConundrumPart1(input []string) int {
	colourMap := map[string]int{
		RED:   12,
		GREEN: 13,
		BLUE:  14,
	}

	idSum := 0
	for _, str := range input {
		id, rounds := parseEntry(str)
		invalid := false
		for _, round := range rounds {
			for _, choice := range round.choices {
				if choice.count > colourMap[choice.colour] {
					invalid = true
					break
				}
			}
		}
		if !invalid {
			idSum += id
		}
	}
	return idSum
}

type GameRound struct {
	choices []GameEntry
}

type GameEntry struct {
	count  int
	colour string
}

func Entry(count int, colour string) GameEntry {
	return GameEntry{
		count:  count,
		colour: colour,
	}
}

func Round(entries ...GameEntry) GameRound {
	return GameRound{choices: entries}
}

func parseEntry(input string) (int, []GameRound) {
	split := strings.Split(input, ":")
	id, err := strconv.Atoi(split[0][5:])
	if err != nil {
		log.Panicf("illegal input format")
	}

	var rounds []GameRound

	rawRounds := strings.Split(strings.TrimSpace(split[1]), ";")

	for _, round := range rawRounds {
		var entries []GameEntry
		rawEntries := strings.Split(round, ",")
		for _, entry := range rawEntries {
			parsedEntry := strings.Split(strings.TrimSpace(entry), " ")
			count, err := strconv.Atoi(strings.TrimSpace(parsedEntry[0]))
			if err != nil {
				panic("illegal game round format")
			}

			colour := strings.TrimSpace(parsedEntry[1])

			entries = append(entries, Entry(count, colour))
		}

		rounds = append(rounds, Round(entries...))
	}

	return id, rounds
}

func cubeConundrumPart2(input []string) int {
	sum := 0
	for _, str := range input {
		maxCubes := map[string]int{}
		_, rounds := parseEntry(str)
		for _, round := range rounds {
			for _, entry := range round.choices {
				maxCubes[entry.colour] = int(math.Max(float64(maxCubes[entry.colour]), float64(entry.count)))
			}
		}
		sum += maxCubes[RED] * maxCubes[GREEN] * maxCubes[BLUE]
	}
	return sum
}
