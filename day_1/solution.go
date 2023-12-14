package day1

import (
	"github.com/pivovarit/aoc/util"
	"slices"
	"unicode"
)

func run() {
	input := util.ReadInput()

	util.Timed("trebuchetPart1", func() int {
		return trebuchetPart1(input)
	})
	util.Timed("trebuchetPart2", func() int {
		return trebuchetPart2(input)
	})
}

func trebuchetPart1(input []string) int {
	result := 0

	for _, entry := range input {
		var digits = [2]int{-1, -1}
		for _, char := range entry {
			if unicode.IsDigit(char) {
				digits[0] = int(char - '0')
				break
			}
		}

		for idx := range entry {
			char := entry[len(entry)-1-idx]
			if unicode.IsDigit(rune(char)) {
				digits[1] = int(char - '0')
				break
			}
		}

		result += digits[0]*10 + digits[1]
	}

	return result
}

func trebuchetPart2(input []string) int {
	result := 0

	for _, entry := range input {
		var digits = [2]int{-1, -1}
		for idx, char := range entry {
			if unicode.IsDigit(char) {
				digits[0] = int(char - '0')
				break
			} else if slices.Contains(firstChars, uint8(char)) {
				digit, found := getDigits(idx, entry)
				if found {
					digits[0] = digit
					break
				}
			}
		}

		for idx := range entry {
			char := entry[len(entry)-1-idx]
			if unicode.IsDigit(rune(char)) {
				digits[1] = int(char - '0')
				break
			} else if slices.Contains(lastChars, char) {
				digit, found := getDigitsBackwards(idx, entry)
				if found {
					digits[1] = digit
					break
				}
			}
		}
		result += digits[0]*10 + digits[1]
	}

	return result
}

var firstChars = []uint8{'o', 't', 'f', 's', 'e', 'n'}
var lastChars = []uint8{'e', 'n', 'o', 'r', 't', 'x'}

var wordsToNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getDigits(idx int, entry string) (int, bool) {
	adjustedEntryLength := len(entry) - idx
	for word := range wordsToNumbers {
		if len(word) <= adjustedEntryLength && word[0] == entry[idx] {
			match := true
			for i := range word {
				if entry[i+idx] != word[i] {
					match = false
					break
				}
			}
			if match {
				return wordsToNumbers[word], true
			}
		}
	}
	return 0, false
}

func getDigitsBackwards(idx int, entry string) (int, bool) {
	adjustedEntryLength := len(entry) - idx
	for word := range wordsToNumbers {
		if len(word) <= adjustedEntryLength && word[len(word)-1] == entry[len(entry)-1-idx] {
			match := true
			for i := range word {
				if entry[len(entry)-1-idx-i] != word[len(word)-1-i] {
					match = false
					break
				}
			}
			if match {
				return wordsToNumbers[word], true
			}
		}
	}
	return 0, false
}
