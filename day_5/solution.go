package day5

import (
	"github.com/pivovarit/aoc/util"
	"math"
	"sort"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("almanacPart1", func() int {
		return almanacPart1(input)
	})
	util.Timed("almanacPart2", func() int {
		return almanacPart2(input)
	})
}

type AlmanacMap struct {
	DestinationRange Range
	SourceRange      Range
}

type Range struct {
	start int
	end   int
}

func (r *Range) offset(offset int) Range {
	return Range{
		start: r.start + offset,
		end:   r.end + offset,
	}
}

func almanacPart1(input []string) int {
	parser := Parser{Tokenizer{input: input}}
	state := parser.parse()

	var minLocation = math.MaxInt
	for _, seed := range state.seeds {
		soil := lookup(seed, state.seedToSoil)
		fertilizer := lookup(soil, state.soilToFertilizer)
		water := lookup(fertilizer, state.fertilizerToWater)
		light := lookup(water, state.waterToLight)
		temperature := lookup(light, state.lightToTemperature)
		humidity := lookup(temperature, state.temperatureToHumidity)
		location := lookup(humidity, state.humidityToLocation)

		minLocation = min(minLocation, location)
	}

	return minLocation
}

func almanacPart2(input []string) int {
	parser := Parser{Tokenizer{input: input}}
	state := parser.parse()
	seedRanges := asRanges(state.seeds)
	soilRanges := mapRanges(seedRanges, state.seedToSoil)
	fertilizerRanges := mapRanges(soilRanges, state.soilToFertilizer)
	waterRanges := mapRanges(fertilizerRanges, state.fertilizerToWater)
	lightRanges := mapRanges(waterRanges, state.waterToLight)
	temperatureRanges := mapRanges(lightRanges, state.lightToTemperature)
	humidityRanges := mapRanges(temperatureRanges, state.temperatureToHumidity)
	locationRanges := mapRanges(humidityRanges, state.humidityToLocation)

	var minLocation = math.MaxInt
	for _, location := range locationRanges {
		minLocation = min(minLocation, location.start)
	}

	return minLocation
}

func mapRanges(from []Range, to []AlmanacMap) []Range {
	var mapped = make([]Range, 0)
	for _, r1 := range from {
		var mappings []Range
		for _, r2 := range to {
			source, overlaps := intersection(r1, r2.SourceRange)
			if overlaps {
				diff := r2.DestinationRange.start - r2.SourceRange.start
				mapped = append(mapped, source.offset(diff))
				mappings = append(mappings, source)
			}
		}

		mapped = append(mapped, subtract(r1, mappings)...)
	}
	return mapped
}

func subtract(r1 Range, ranges []Range) []Range {
	remaining := []Range{r1}

	for _, r2 := range ranges {
		var newResult []Range

		for _, current := range remaining {
			if r2.end <= current.start || r2.start >= current.end {
				newResult = append(newResult, current)
			} else {
				if r2.start > current.start {
					newResult = append(newResult, Range{current.start, r2.start})
				}
				if r2.end < current.end {
					newResult = append(newResult, Range{r2.end, current.end})
				}
			}
		}

		remaining = newResult
	}

	return remaining
}

func intersection(r1 Range, r2 Range) (Range, bool) {
	if r1.start > r2.end || r2.start > r1.end {
		return Range{}, false
	}

	return Range{max(r1.start, r2.start), min(r1.end, r2.end)}, true
}

func lookup(key int, mappings []AlmanacMap) int {
	for _, mapping := range mappings {
		if key >= mapping.SourceRange.start && key < mapping.SourceRange.end {
			return mapping.DestinationRange.start + (key - mapping.SourceRange.start)
		}
	}

	return key
}

func asRanges(seeds []int) []Range {
	var ranges = make([]Range, 0)

	for i := 0; i < len(seeds); i += 2 {
		ranges = append(ranges, Range{
			start: seeds[i],
			end:   seeds[i] + seeds[i+1],
		})
	}
	return ranges
}

type Parser struct {
	tokenizer Tokenizer
}

func (p *Parser) parseSeeds(input string) []int {
	var result []int
	for _, str := range strings.Fields(input) {
		seed, _ := strconv.Atoi(str)
		result = append(result, seed)
	}
	return result
}

func (p *Parser) parseAlmanac(input []string) []AlmanacMap {
	var result []AlmanacMap
	for _, line := range input {
		fields := strings.Fields(line)
		destinationRangeStart, _ := strconv.Atoi(fields[0])
		sourceRangeStart, _ := strconv.Atoi(fields[1])
		rangeLength, _ := strconv.Atoi(fields[2])
		fertilizerMap := AlmanacMap{
			DestinationRange: Range{destinationRangeStart, destinationRangeStart + rangeLength},
			SourceRange:      Range{sourceRangeStart, sourceRangeStart + rangeLength},
		}
		result = append(result, fertilizerMap)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].SourceRange.start < result[j].SourceRange.start
	})

	return result
}

type State struct {
	seeds                 []int
	seedToSoil            []AlmanacMap
	soilToFertilizer      []AlmanacMap
	fertilizerToWater     []AlmanacMap
	waterToLight          []AlmanacMap
	lightToTemperature    []AlmanacMap
	temperatureToHumidity []AlmanacMap
	humidityToLocation    []AlmanacMap
}

func (p *Parser) parse() State {
	seeds := p.parseSeeds(strings.Split(p.tokenizer.next()[0], ":")[1])
	seedToSoil := p.parseAlmanac(p.tokenizer.next()[1:])
	seedToFertilizer := p.parseAlmanac(p.tokenizer.next()[1:])
	fertilizerToWater := p.parseAlmanac(p.tokenizer.next()[1:])
	waterToLight := p.parseAlmanac(p.tokenizer.next()[1:])
	lightToTemperature := p.parseAlmanac(p.tokenizer.next()[1:])
	temperatureToHumidity := p.parseAlmanac(p.tokenizer.next()[1:])
	humidityToLocation := p.parseAlmanac(p.tokenizer.next()[1:])

	if p.tokenizer.hasNext() {
		panic("illegal token encountered")
	}
	r := State{
		seeds:                 seeds,
		seedToSoil:            seedToSoil,
		soilToFertilizer:      seedToFertilizer,
		fertilizerToWater:     fertilizerToWater,
		waterToLight:          waterToLight,
		lightToTemperature:    lightToTemperature,
		temperatureToHumidity: temperatureToHumidity,
		humidityToLocation:    humidityToLocation,
	}
	return r
}

type Tokenizer struct {
	input   []string
	lastIdx int
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
