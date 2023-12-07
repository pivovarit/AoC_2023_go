package util

import (
	"bufio"
	"log"
	"os"
	"time"
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

func Timed(name string, action func() int) {
	start := time.Now()
	result := action()
	elapsed := time.Since(start)
	log.Printf("%s() result: %d, execution time: %s", name, result, elapsed)
}
