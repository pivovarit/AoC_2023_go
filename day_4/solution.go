package day4

import (
	"github.com/pivovarit/aoc/util"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("scratchCardsPart1", func() {
		println(scratchCardsPart1(input))
	})
	util.Timed("scratchCardsPart2", func() {
		println(scratchCardsPart2(input))
	})
	util.Timed("scratchCardsPart2_optimized", func() {
		println(scratchCardsPart2_optimized(input))
	})
}

func scratchCardsPart1(input []string) int {
	var sum = 0
	for _, line := range input {
		p := parser{line}
		card := p.card()
		matches := card.matches()
		if matches != 0 {
			sum += int(math.Pow(2, float64(matches-1)))
		}
	}
	return sum
}

func scratchCardsPart2(input []string) int {
	var sum = 0
	var cards []card
	for _, line := range input {
		p := parser{line}
		cards = append(cards, p.card())
	}
	var results = make(map[int]int)
	for _, c := range cards {
		sum += count(cards, c, results)
	}
	return sum
}

func count(deck []card, crd card, results map[int]int) int {
	var sum = 1
	for i := crd.id; i < crd.id+crd.matches(); i++ {
		if i < len(deck) {
			if val, ok := results[crd.id]; ok {
				return val
			}
			sum += count(deck, deck[i], results)
		}
	}
	results[crd.id] = sum
	return sum
}

type parser struct {
	input string
}

type card struct {
	id              int
	winning, chosen []int
}

func (c *card) matches() int {
	var matching = 0
	for _, number := range c.chosen {
		if slices.Contains(c.winning, number) {
			matching++
		}
	}
	return matching
}

func (p *parser) card() card {
	return card{
		id:      p.roundId(),
		winning: p.winningNumbers(),
		chosen:  p.chosenNumbers(),
	}
}

func (p *parser) roundId() int {
	split := strings.Split(p.input, ":")
	id, err := strconv.Atoi(strings.Fields(split[0])[1])
	if err != nil {
		log.Panicf("illegal input format")
	}

	return id
}

func parseCards(input string) []int {
	cards := strings.Fields(input)
	var ids = make([]int, 0)

	for _, card := range cards {
		id, err := strconv.Atoi(card)
		if err != nil {
			log.Panicf("illegal input format")
		}
		ids = append(ids, id)
	}

	return ids
}

func (p *parser) winningNumbers() []int {
	winningCards := strings.Split(strings.Split(p.input, ":")[1], "|")[0]
	return parseCards(strings.TrimSpace(winningCards))
}

func (p *parser) chosenNumbers() []int {
	chosenCards := strings.Split(strings.Split(p.input, ":")[1], "|")[1]
	return parseCards(strings.TrimSpace(chosenCards))
}
