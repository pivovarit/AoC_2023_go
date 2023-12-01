package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func run() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputArray = append(inputArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(trebuchet_part_1(inputArray))
	println(trebuchet_part_2(inputArray))
}

func trebuchet(input []string, processor func(s string) string) int {
	result := 0

	for _, entry := range input {
		entry = processor(entry)

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

func trebuchet_part_1(input []string) int {
	return trebuchet(input, func(s string) string {
		return s
	})
}

func trebuchet_part_2(input []string) int {
	return trebuchet(input, func(s string) string {
		return replaceWordsWithDigits(s)
	})
}

func replaceWordsWithDigits(entry string) string {
	wordsToNumbers := map[string]int{
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

	for idx := range entry {
		for word := range wordsToNumbers {
			if strings.Index(entry, word) == idx {
				entry = strings.Replace(entry, word, strconv.Itoa(wordsToNumbers[word]), 1)
			}
		}
	}
	return entry
}
