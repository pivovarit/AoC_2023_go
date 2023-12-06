package day6

import (
	"github.com/pivovarit/aoc/util"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("waitForItPart1", func() int {
		return waitForItPart1(input)
	})
	util.Timed("waitForItPart2", func() int {
		return waitForItPart2(input)
	})
}

func waitForItPart1(input []string) int {
	parser := Parser{input: input}
	return waitForIt(parser.parse())
}

func waitForItPart2(input []string) int {
	parser := Parser{input: input}
	return waitForIt([]Race{parser.parseSingleRace()})
}

func waitForIt(races []Race) int {
	var result = 0

	for _, race := range races {
		var raceResults = make([]int, race.time+1)
		for buttonTime := 0; buttonTime <= race.time; buttonTime++ {
			if buttonTime == 0 || buttonTime == race.time {
				raceResults[buttonTime] = 0
			} else {
				raceResults[buttonTime] = (race.time - buttonTime) * buttonTime
			}
		}

		wins := 0
		for _, raceResult := range raceResults {
			if raceResult > race.record {
				wins++
			}
		}

		if result == 0 {
			result = wins
		} else {
			result *= wins
		}
	}

	return result
}

type Parser struct {
	input []string
}

func (p *Parser) parse() []Race {
	times := strings.Fields(strings.Split(p.input[0], ":")[1])
	records := strings.Fields(strings.Split(p.input[1], ":")[1])

	var races []Race

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		record, _ := strconv.Atoi(records[i])
		races = append(races, Race{
			time:   time,
			record: record,
		})
	}
	return races
}

func (p *Parser) parseSingleRace() Race {
	times := strings.Fields(strings.Split(p.input[0], ":")[1])
	records := strings.Fields(strings.Split(p.input[1], ":")[1])

	time, _ := strconv.Atoi(strings.Join(times, ""))
	record, _ := strconv.Atoi(strings.Join(records, ""))

	return Race{
		time:   time,
		record: record,
	}
}

type Race struct {
	time, record int
}
