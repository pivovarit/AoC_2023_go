package day1

import (
	"github.com/pivovarit/aoc/util"
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
		var digits [2]int
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
		var digits [2]int
		entryLength := len(entry)
		for idx, char := range entry {
			if isDigit(char) {
				digits[0] = int(char - '0')
				break
			} else {
				value, found := findDigit(char, entryLength, idx, entry)
				if found {
					digits[0] = value
					break
				}
			}
		}

		for idx := range entry {
			char := entry[entryLength-1-idx]
			r := rune(char)
			if isDigit(r) {
				digits[1] = int(char - '0')
				break
			} else {
				digit, found := findDigitBackwards(char, entryLength, idx, entry)
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

func findDigit(char int32, entryLength int, idx int, entry string) (int, bool) {
	adjustedEntryLength := entryLength - idx
	switch char {
	case 'e':
		if adjustedEntryLength >= 5 {
			// eight
			if entry[idx+1] == 'i' && entry[idx+2] == 'g' && entry[idx+3] == 'h' && entry[idx+4] == 't' {
				return 8, true
			}
		}
	case 'f':
		if adjustedEntryLength >= 4 {
			// four
			if entry[idx+1] == 'o' && entry[idx+2] == 'u' && entry[idx+3] == 'r' {
				return 4, true
			}

			// five
			if entry[idx+1] == 'i' && entry[idx+2] == 'v' && entry[idx+3] == 'e' {
				return 5, true
			}

		}
	case 'n':
		if adjustedEntryLength >= 4 {
			// nine
			if entry[idx+1] == 'i' && entry[idx+2] == 'n' && entry[idx+3] == 'e' {
				return 9, true
			}
		}
	case 'o':
		if adjustedEntryLength >= 3 {
			// one
			if entry[idx+1] == 'n' && entry[idx+2] == 'e' {
				return 1, true
			}
		}
	case 's':
		//six
		if adjustedEntryLength >= 3 {
			if entry[idx+1] == 'i' && entry[idx+2] == 'x' {
				return 6, true
			}
		}

		// seven
		if adjustedEntryLength >= 5 {
			if entry[idx+1] == 'e' && entry[idx+2] == 'v' && entry[idx+3] == 'e' && entry[idx+4] == 'n' {
				return 7, true
			}
		}
	case 't':
		//two
		if adjustedEntryLength >= 3 {
			if entry[idx+1] == 'w' && entry[idx+2] == 'o' {
				return 2, true
			}
		}

		// three
		if adjustedEntryLength >= 5 {
			if entry[idx+1] == 'h' && entry[idx+2] == 'r' && entry[idx+3] == 'e' && entry[idx+4] == 'e' {
				return 3, true
			}
		}
	}
	return -1, false
}

func findDigitBackwards(char uint8, entryLength int, idx int, entry string) (int, bool) {
	lastIdx := entryLength - idx - 1
	switch char {
	case 'e':
		// one
		if entryLength >= 3 && entry[lastIdx-1] == 'n' && entry[lastIdx-2] == 'o' {
			return 1, true
		}
		if entryLength >= 4 {
			// five
			if entry[lastIdx-1] == 'v' && entry[lastIdx-2] == 'i' && entry[lastIdx-3] == 'f' {
				return 5, true
			}
			// nine
			if entry[lastIdx-1] == 'n' && entry[lastIdx-2] == 'i' && entry[lastIdx-3] == 'n' {
				return 9, true
			}
		}
		if entryLength >= 5 {
			// three
			if entry[lastIdx-1] == 'e' && entry[lastIdx-2] == 'r' && entry[lastIdx-3] == 'h' && entry[lastIdx-4] == 't' {
				return 3, true
			}
		}

	case 'n':
		// seven
		if entryLength >= 5 && entry[lastIdx-1] == 'e' && entry[lastIdx-2] == 'v' && entry[lastIdx-3] == 'e' && entry[lastIdx-4] == 's' {
			return 7, true
		}
	case 'o':
		// two
		if entryLength >= 3 && entry[lastIdx-1] == 'w' && entry[lastIdx-2] == 't' {
			return 2, true
		}
	case 'r':
		// four
		if entryLength >= 4 && entry[lastIdx-1] == 'u' && entry[lastIdx-2] == 'o' && entry[lastIdx-3] == 'f' {
			return 4, true
		}
	case 't':
		// eight
		if entryLength >= 5 && entry[lastIdx-1] == 'h' && entry[lastIdx-2] == 'g' && entry[lastIdx-3] == 'i' && entry[lastIdx-4] == 'e' {
			return 8, true
		}
	case 'x':
		// six
		if entryLength >= 3 && entry[lastIdx-1] == 'i' && entry[lastIdx-2] == 's' {
			return 6, true
		}
	default:
		return -1, false
	}
	return -1, false
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
