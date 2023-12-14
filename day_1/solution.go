package day1

import (
	"github.com/pivovarit/aoc/util"
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
				digit := char - '0'
				if digits[0] == -1 {
					digits[0] = int(digit)
				}

				digits[1] = int(digit)
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
				digit := char - '0'
				if digits[0] == -1 {
					digits[0] = int(digit)
				}

				digits[1] = int(digit)
			} else if maybeWord(char) {
				digit := getDigits(idx, entry)
				if digit != 0 {
					if digits[0] == -1 {
						digits[0] = digit
					}
					digits[1] = digit
				}
			}
		}

		result += digits[0]*10 + digits[1]
	}

	return result
}

var firstChars = []int32{'o', 't', 'f', 's', 'e', 'n'}

func maybeWord(char int32) bool {
	for _, c := range firstChars {
		if c == char {
			return true
		}
	}
	return false
}

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

func getDigits(idx int, entry string) int {
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
				return wordsToNumbers[word]
			}
		}
	}
	return 0
}
