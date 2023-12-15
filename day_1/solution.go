package main

import (
	"github.com/pivovarit/aoc/util"
  "fmt"
)

var numberHashToInt = getNumberStringBuckets()

func hash(s *string) int {
	return int((*s)[1]+(*s)[2]) - 208
}

func getNumberStringBuckets() [23]int {
	return [23]int{8, -1, -2, 1, -1, -1, -1, 9, -1, -1, 3, 8, -1, -1, -1, 5, -1, 6, -1, -1, 4, -1, 2}
}

func main() {
	input := util.ReadInput()

	util.Timed("trebuchetPart1", func() int {
		return trebuchetPart1(input)
	})
	util.Timed("trebuchetPart2", func() int {
		return trebuchetPart2(input)
	})

	util.Timed("trebuchetPart2 Jacek", func() int {
		return solve(input)
	})
  for i, line := range input {
    sol1 := trebuchetPart2(input[i:i+1])
    sol2 := solve(input[i:i+1])
    if sol1 != sol2 {
      fmt.Println("Mismatch in two solutions", line, sol1, sol2)
    }
  }
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

type State struct {
	line       string
	position   int
	firstDigit int
	lastDigit  int
}

func newState(line string) State {
	return State{line: line, position: 0, firstDigit: -1, lastDigit: -1}
}

func (s *State) updateDigit(digit int) {
	if s.firstDigit == -1 {
		s.firstDigit = digit
	}
	s.lastDigit = digit
}

func (s *State) parseDigitNew() {
	for s.position < len(s.line) {
		c := s.line[s.position]
		s.position += 1
		switch c {
		case '1':
			s.updateDigit(1)
		case '2':
			s.updateDigit(2)
		case '3':
			s.updateDigit(3)
		case '4':
			s.updateDigit(4)
		case '5':
			s.updateDigit(5)
		case '6':
			s.updateDigit(6)
		case '7':
			s.updateDigit(7)
		case '8':
			s.updateDigit(8)
		case '9':
			s.updateDigit(9)
		case 'o':
			if s.position+1 < len(s.line) && s.line[s.position] == 'n' {
				if s.line[s.position+1] == 'e' {
					s.updateDigit(1)
					s.position += 1 // it can be eight later
				}
			}
		case 't':
			if s.position+1 < len(s.line) && s.line[s.position] == 'w' {
				if s.line[s.position+1] == 'o' {
					s.updateDigit(2)
					s.position += 1 // it can be one later
				}
			} else if s.position+3 < len(s.line) && s.line[s.position] == 'h' {
				if s.line[s.position+1] == 'r' {
					if s.line[s.position+2] == 'e' {
						if s.line[s.position+3] == 'e' {
							s.updateDigit(3)
							s.position += 2 // it can be eight later
						}
					}
				}
			}

		case 'f':
			if s.position+2 < len(s.line) && s.line[s.position] == 'o' {
				if s.line[s.position+1] == 'u' {
					if s.line[s.position+2] == 'r' {
						s.updateDigit(4)
						s.position += 3
					}
				}
			} else if s.position+2 < len(s.line) && s.line[s.position] == 'i' {
				if s.line[s.position+1] == 'v' {
					if s.line[s.position+2] == 'e' {
						s.updateDigit(5)
						s.position += 2 // it can be eight later
					}
				}
			}
		case 's':
			if s.position+1 < len(s.line) && s.line[s.position] == 'i' {
				if s.line[s.position+1] == 'x' {
					s.updateDigit(6)
					s.position += 2
				}
			} else if s.position+3 < len(s.line) && s.line[s.position] == 'e' {
				if s.line[s.position+1] == 'v' {
					if s.line[s.position+2] == 'e' {
						if s.line[s.position+3] == 'n' {
							s.updateDigit(7)
							s.position += 3 // it can be nine later
						}
					}
				}
			}
		case 'e':
			if s.position+3 < len(s.line) && s.line[s.position] == 'i' {
				if s.line[s.position+1] == 'g' {
					if s.line[s.position+2] == 'h' {
						if s.line[s.position+3] == 't' {
							s.updateDigit(8)
							s.position += 3 // it can be three later
						}
					}
				}
			}
		case 'n':
			if s.position+2 < len(s.line) && s.line[s.position] == 'i' {
				if s.line[s.position+1] == 'n' {
					if s.line[s.position+2] == 'e' {
						s.updateDigit(9)
						s.position += 2 // it can be eight later
					}
				}
			}
		default:
		}
	}
}

func solve(input []string) int {
	sum := 0
	for _, line := range input {
		state := newState(line)
		state.parseDigitNew()
		sum += state.firstDigit*10 + state.lastDigit
	}
	return sum
}
