package day7

import (
	"github.com/pivovarit/aoc/util"
	"sort"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()
	util.Timed("camelCardsPart1", func() int {
		return camelCardsPart1(input)
	})
	util.Timed("camelCardsPart2", func() int {
		return camelCardsPart2(input)
	})
}

const (
	highCard HandType = iota
	onePair
	twoPair
	threeOfKind
	fullHouse
	fourOfKind
	fiveOfKind
	joker = CardType('J')
)

type (
	CardType int32
	HandType int
	Hand     struct {
		cards [5]CardType
		bid   int
	}
)

func strength(card CardType, jokers bool) int {
	var values = map[CardType]int{
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'J': 10,
		'Q': 11,
		'K': 12,
		'A': 13,
	}
	if jokers {
		values[joker] = 0
	}
	return values[card]
}

func camelCardsPart1(input []string) int {
	return camelCards(input, false)
}

func camelCardsPart2(input []string) int {
	return camelCards(input, true)
}

func camelCards(input []string, jokers bool) int {
	hands := parse(input)

	sort.Slice(hands, func(i, j int) bool { return hands[i].compare(hands[j], jokers) })

	var result = 0

	for idx, hand := range hands {
		result += (idx + 1) * hand.bid
	}

	return result
}

func (h *Hand) compare(h2 Hand, jokers bool) bool {
	if h.getType(jokers) == h2.getType(jokers) {
		for i := 0; i < len(h.cards); i++ {
			s1 := strength(h.cards[i], jokers)
			s2 := strength(h2.cards[i], jokers)
			if s1 < s2 {
				return true
			}
			if s1 > s2 {
				return false
			}
		}
		return false
	} else {
		return h.getType(jokers) < h2.getType(jokers)
	}
}

func (h *Hand) getType(withJokers bool) HandType {
	var result = make(map[CardType]int)
	for _, card := range h.cards {
		val, found := result[card]
		if found {
			result[card] = val + 1
		} else {
			result[card] = 1
		}
	}

	size := len(result)
	jokers := result[joker]
	switch size {
	case 1:
		return fiveOfKind
	case 2:
		if withJokers && (jokers == 3 || jokers == 2) {
			return fiveOfKind
		}
		for _, count := range result {
			if count == 4 {
				if withJokers && jokers > 0 {
					return fiveOfKind
				}
				return fourOfKind
			}
		}
		return fullHouse
	case 3:
		for _, count := range result {
			if count == 3 {
				if withJokers {
					switch jokers {
					case 3:
						return fourOfKind
					case 2:
						return fiveOfKind
					case 1:
						return fourOfKind
					}
				}
				return threeOfKind
			}
		}
		if withJokers {
			switch jokers {
			case 2:
				return fourOfKind
			case 1:
				return fullHouse
			}
		}
		return twoPair
	case 4:
		if withJokers && jokers > 0 {
			return threeOfKind
		}
		return onePair
	default:
		if withJokers {
			switch jokers {
			case 2:
				return threeOfKind
			case 1:
				return onePair
			}
		}
		return highCard
	}
}

func parse(lines []string) []Hand {
	var hands []Hand
	for _, line := range lines {
		parsed := strings.Fields(line)
		bid, _ := strconv.Atoi(parsed[1])
		var cards [5]CardType
		for i, card := range parsed[0] {
			cards[i] = CardType(card)
		}
		hands = append(hands, Hand{
			cards: cards,
			bid:   bid,
		})
	}
	return hands
}
