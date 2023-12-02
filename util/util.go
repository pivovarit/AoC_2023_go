package util

import (
	"bufio"
	"log"
	"os"
)

func ReadInput() []string {
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
	return inputArray
}
