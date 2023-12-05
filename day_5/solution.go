package day5

import (
	"fmt"
	"github.com/pivovarit/aoc/util"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
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

func almanacPart1(input []string) int {
	parser := Parser{Tokenizer{input: input}}
	state := parser.parse()

	var locations = make([]int, 0)
	for _, seed := range state.seeds {
		soil := lookup(seed, state.seedToSoil)
		fertilizer := lookup(soil, state.soilToFertilizer)
		water := lookup(fertilizer, state.fertilizerToWater)
		light := lookup(water, state.waterToLight)
		temperature := lookup(light, state.lightToTemperature)
		humidity := lookup(temperature, state.temperatureToHumidity)
		location := lookup(humidity, state.humidityToLocation)

		locations = append(locations, location)
	}

	sort.Ints(locations)

	return locations[0]
}

func almanacPart2(input []string) int {
	parser := Parser{Tokenizer{input: input}}
	state := parser.parse()
	seedRanges := expand(state.seeds)

	var minLocation = math.MaxInt

	var wg sync.WaitGroup
	var mu sync.Mutex

	for idx, seeds := range seedRanges {
		wg.Add(1)
		go func(idx int, seeds []int) {
			defer wg.Done()
			fmt.Printf("Seed range: %d, start: %d, count: %d\n", idx, seeds[0], len(seeds))

			var localMin = math.MaxInt

			for _, seed := range seeds {
				soil := lookup(seed, state.seedToSoil)
				fertilizer := lookup(soil, state.soilToFertilizer)
				water := lookup(fertilizer, state.fertilizerToWater)
				light := lookup(water, state.waterToLight)
				temperature := lookup(light, state.lightToTemperature)
				humidity := lookup(temperature, state.temperatureToHumidity)
				location := lookup(humidity, state.humidityToLocation)

				localMin = int(math.Min(float64(localMin), float64(location)))
			}
			fmt.Printf("Seed range: %d, minLocation: %d\n", idx, localMin)

			mu.Lock()
			minLocation = int(math.Min(float64(localMin), float64(minLocation)))
			mu.Unlock()
		}(idx, seeds)
	}

	wg.Wait()

	return minLocation
}

func lookup(key int, mappings []AlmanacMap) int {
	for _, mapping := range mappings {
		if key >= mapping.SourceRange.start && key < mapping.SourceRange.end {
			return mapping.DestinationRange.start + (key - mapping.SourceRange.start)
		}
	}

	return key
}

func unroll(start int, count int) []int {
	var seeds = make([]int, 0)
	for seed := start; seed < start+count; seed++ {
		seeds = append(seeds, seed)
	}
	return seeds
}

func expand(seedRanges []int) [][]int {
	if len(seedRanges)%2 != 0 {
		panic("expecting even seed ranges")
	}

	var seeds = make([][]int, 0)
	for i := 0; i < len(seedRanges); i += 2 {
		seeds = append(seeds, unroll(seedRanges[i], seedRanges[i+1]))
	}
	return seeds
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
