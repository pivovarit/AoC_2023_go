package day4

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func scratchCardsPart2_optimized(input []string) int {
	cardsCount := make(map[int]int)
	for _, line := range input {
		cardNum, winning, chosen := parse(line)
		cardsCount[cardNum]++
		for i := 1; i <= matches(winning, chosen); i++ {
			cardsCount[cardNum+i] += cardsCount[cardNum]
		}
	}

	var sum int
	for _, count := range cardsCount {
		sum += count
	}

	return sum
}

func matches(winning, chosen []int) int {
	var matching = 0
	for _, number := range chosen {
		if slices.Contains(winning, number) {
			matching++
		}
	}
	return matching
}

func parse(input string) (int, []int, []int) {
	return roundId(input), winningNumbers(input), chosenNumbers(input)
}

func roundId(input string) int {
	split := strings.Split(input, ":")
	id, err := strconv.Atoi(strings.Fields(split[0])[1])
	if err != nil {
		log.Panicf("illegal input format")
	}

	return id
}

func winningNumbers(input string) []int {
	winningCards := strings.Split(strings.Split(input, ":")[1], "|")[0]
	return parseCards(strings.TrimSpace(winningCards))
}

func chosenNumbers(input string) []int {
	chosenCards := strings.Split(strings.Split(input, ":")[1], "|")[1]
	return parseCards(strings.TrimSpace(chosenCards))
}
