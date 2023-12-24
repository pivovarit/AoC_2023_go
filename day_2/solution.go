package day2

import (
	"github.com/pivovarit/aoc/util"
	"runtime"
	"sync"
)

func run() {
	input := util.ReadInput()

	util.Timed("cubeConundrumPart1Sequential", func() int {
		return cubeConundrumPart1Sequential(input)
	})
	util.Timed("cubeConundrumPart1Parallel", func() int {
		return cubeConundrumPart1Parallel(input)
	})
	util.Timed("cubeConundrumPart2Sequential", func() int {
		return cubeConundrumPart2Sequential(input)
	})

	util.Timed("cubeConundrumPart2Parallel", func() int {
		return cubeConundrumPart2Parallel(input)
	})
}

const (
	redCubesCount   = 12
	greenCubesCount = 13
	blueCubesCount  = 14
)

type State int

const (
	processingGameId State = iota
	processingCubeCount
	processingCubeColour
)

const (
	emptySeparator      = ' '
	gameIdSeparator     = ':'
	cubeChoiceSeparator = ','
	gameSeparator       = ';'
)

func cubeConundrumPart1Sequential(input []string) (sum int) {
LineLoop:
	for _, line := range input {
		state := processingGameId
		gameId := 0
		cubeCount := 0
		cubeColourProcessed := false
		colourPrefix := true
		for _, char := range line[5:] {
			switch state {
			case processingGameId:
				if char != gameIdSeparator {
					gameId = gameId*10 + int(char-'0')
				} else {
					state = processingCubeCount
				}
			case processingCubeCount:
				switch {
				case char == emptySeparator && colourPrefix:
					colourPrefix = false

				case char != emptySeparator:
					cubeCount = cubeCount*10 + int(char-'0')

				case char == emptySeparator && !colourPrefix:
					state = processingCubeColour
				}
			case processingCubeColour:
				if char != cubeChoiceSeparator && char != gameSeparator {
					if cubeColourProcessed {
						continue
					}
					switch char {
					case 'r':
						if cubeCount > redCubesCount {
							continue LineLoop
						}
					case 'b':
						if cubeCount > blueCubesCount {
							continue LineLoop
						}
					case 'g':
						if cubeCount > greenCubesCount {
							continue LineLoop
						}
					}
					cubeColourProcessed = true
				} else {
					cubeColourProcessed = false
					cubeCount = 0
					state = processingCubeCount
					colourPrefix = true
				}
			}
		}
		sum += gameId
	}
	return sum
}

func cubeConundrumPart1Parallel(input []string) (sum int) {
	var wg sync.WaitGroup
	parallelism := runtime.NumCPU()
	inputLen := len(input)
	parallelism = min(inputLen, parallelism)

	var idSumPerCPU = make([]int, parallelism)

	for cpuId := 0; cpuId < parallelism; cpuId++ {
		wg.Add(1)
		go func(cpu int) {
			defer wg.Done()
		LineLoop:
			for i := cpu; i < inputLen; i = i + parallelism {
				line := input[i]
				state := processingGameId
				gameId := 0
				cubeCount := 0
				cubeColourProcessed := false
				colourPrefix := true
				for _, char := range line[5:] {
					switch state {
					case processingGameId:
						if char != gameIdSeparator {
							gameId = gameId*10 + int(char-'0')
						} else {
							state = processingCubeCount
						}
					case processingCubeCount:
						if char == emptySeparator && colourPrefix {
							colourPrefix = false
							break
						}

						if char != emptySeparator {
							cubeCount = cubeCount*10 + int(char-'0')
							break
						}

						if char == emptySeparator && !colourPrefix {
							state = processingCubeColour
						}
					case processingCubeColour:
						if char != cubeChoiceSeparator && char != gameSeparator {
							if cubeColourProcessed {
								continue
							}
							switch char {
							case 'r':
								if cubeCount > redCubesCount {
									continue LineLoop
								}
							case 'b':
								if cubeCount > blueCubesCount {
									continue LineLoop
								}
							case 'g':
								if cubeCount > greenCubesCount {
									continue LineLoop
								}
							}
							cubeColourProcessed = true
						} else {
							cubeColourProcessed = false
							cubeCount = 0
							state = processingCubeCount
							colourPrefix = true
						}
					}
				}
				idSumPerCPU[cpu] += gameId
			}
		}(cpuId)
	}

	wg.Wait()

	for _, sumPerCpu := range idSumPerCPU {
		sum += sumPerCpu
	}

	return sum
}

func cubeConundrumPart2Sequential(input []string) (sum int) {
	for _, line := range input {
		state := processingGameId
		gameId := 0
		cubeCount := 0
		cubeColourProcessed := false
		colourPrefix := true
		maxGreen := -1
		maxBlue := -1
		maxRed := -1
		for _, char := range line[5:] {
			switch state {
			case processingGameId:
				if char != gameIdSeparator {
					gameId = gameId*10 + int(char-'0')
				} else {
					state = processingCubeCount
				}
			case processingCubeCount:
				if char == emptySeparator && colourPrefix {
					colourPrefix = false
					break
				}

				if char != emptySeparator {
					cubeCount = cubeCount*10 + int(char-'0')
					break
				}

				if char == emptySeparator && !colourPrefix {
					state = processingCubeColour
				}
			case processingCubeColour:
				if char != cubeChoiceSeparator && char != gameSeparator {
					if cubeColourProcessed {
						continue
					}
					switch char {
					case 'r':
						maxRed = max(maxRed, cubeCount)
					case 'b':
						maxBlue = max(maxBlue, cubeCount)
					case 'g':
						maxGreen = max(maxGreen, cubeCount)
					}
					cubeColourProcessed = true
				} else {
					cubeColourProcessed = false
					cubeCount = 0
					state = processingCubeCount
					colourPrefix = true
				}
			}
		}
		sum += maxBlue * maxGreen * maxRed
	}
	return sum
}

func cubeConundrumPart2Parallel(input []string) (sum int) {
	var wg sync.WaitGroup
	parallelism := runtime.NumCPU()
	inputLen := len(input)
	parallelism = min(inputLen, parallelism)

	var powerPerCPU = make([]int, parallelism)

	for cpuId := 0; cpuId < parallelism; cpuId++ {
		wg.Add(1)
		go func(cpu int) {
			defer wg.Done()
			for i := cpu; i < inputLen; i = i + parallelism {
				line := input[i]
				state := processingGameId
				gameId := 0
				cubeCount := 0
				cubeColourProcessed := false
				colourPrefix := true
				maxGreen := -1
				maxBlue := -1
				maxRed := -1
				for _, char := range line[5:] {
					switch state {
					case processingGameId:
						if char != gameIdSeparator {
							gameId = gameId*10 + int(char-'0')
						} else {
							state = processingCubeCount
						}
					case processingCubeCount:
						if char == emptySeparator && colourPrefix {
							colourPrefix = false
							break
						}

						if char != emptySeparator {
							cubeCount = cubeCount*10 + int(char-'0')
							break
						}

						if char == emptySeparator && !colourPrefix {
							state = processingCubeColour
						}
					case processingCubeColour:
						if char != cubeChoiceSeparator && char != gameSeparator {
							if cubeColourProcessed {
								continue
							}
							switch char {
							case 'r':
								maxRed = max(maxRed, cubeCount)
							case 'b':
								maxBlue = max(maxBlue, cubeCount)
							case 'g':
								maxGreen = max(maxGreen, cubeCount)
							}
							cubeColourProcessed = true
						} else {
							cubeColourProcessed = false
							cubeCount = 0
							state = processingCubeCount
							colourPrefix = true
						}
					}
				}
				powerPerCPU[cpu] += maxBlue * maxGreen * maxRed
			}
		}(cpuId)
	}
	wg.Wait()

	for _, sumPerCpu := range powerPerCPU {
		sum += sumPerCpu
	}

	return sum
}
