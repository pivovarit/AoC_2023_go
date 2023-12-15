package day1

import (
	"github.com/pivovarit/aoc/util"
)

var numberHashToInt = getNumberStringBuckets()

func hash(s *string) int {
	return int((*s)[1]+(*s)[2]) - 208
}

func getNumberStringBuckets() [23]int {
	return [23]int{8, -1, -2, 1, -1, -1, -1, 9, -1, -1, 3, 8, -1, -1, -1, 5, -1, 6, -1, -1, 4, -1, 2}
}

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
			if isDigit(char) {
				digits[0] = int(char - '0')
				break
			}
		}

		for idx := range entry {
			char := entry[len(entry)-1-idx]
			if isDigit(rune(char)) {
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
			if isDigit(char) {
				digits[0] = int(char - '0')
				break
			} else if isInFirstChars(&char) {
				digit, found := getDigits(idx, entry)
				if found {
					digits[0] = digit
					break
				}
			}
		}

		entryLength := len(entry)
		for idx := range entry {
			char := entry[entryLength-1-idx]
			r := rune(char)
			if isDigit(r) {
				digits[1] = int(char - '0')
				break
			} else if isInLastChars(&char) {
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

var wordsToNumbersKeys = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func isInFirstChars(char *int32) bool {
	switch *char {
	case 'e':
		return true
	case 'f':
		return true
	case 'n':
		return true
	case 'o':
		return true
	case 's':
		return true
	case 't':
		return true
	default:
		return false
	}
}

func isInLastChars(char *uint8) bool {
	switch *char {
	case 'e':
		return true
	case 'n':
		return true
	case 'o':
		return true
	case 'r':
		return true
	case 't':
		return true
	case 'x':
		return true
	default:
		return false
	}
}

func getDigits(idx int, entry string) (int, bool) {
	adjustedEntryLength := len(entry) - idx
	for _, word := range wordsToNumbersKeys {
		if len(word) <= adjustedEntryLength && word[0] == entry[idx] {
			match := true
			for i := range word {
				if entry[i+idx] != word[i] {
					match = false
					break
				}
			}
			if match {
				return numberHashToInt[hash(&word)], true
			}
		}
	}
	return 0, false
}

func getDigitsBackwards(idx int, entry string) (int, bool) {
	entryLength := len(entry)
	adjustedEntryLength := entryLength - idx
	for _, word := range wordsToNumbersKeys {
		if len(word) <= adjustedEntryLength && word[len(word)-1] == entry[entryLength-1-idx] {
			match := true
			for i := range word {
				if entry[entryLength-1-idx-i] != word[len(word)-1-i] {
					match = false
					break
				}
			}
			if match {
				return numberHashToInt[hash(&word)], true
			}
		}
	}
	return 0, false
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
