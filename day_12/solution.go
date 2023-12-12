package day12

import (
	"github.com/pivovarit/aoc/util"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("hotSpringsPart1", func() int {
		return hotSpringsPart1(input)
	})
	util.Timed("hotSpringsPart2", func() int {
		return hotSpringsPart2(input)
	})
}

const (
	damaged     = '#'
	operational = '.'
	unknown     = '?'
)

func countValidConfigurations(config string, nums []int) int {
	if config == "" {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	if len(nums) == 0 {
		if strings.ContainsRune(config, damaged) {
			return 0
		}
		return 1
	}

	result := 0

	if config[0] == operational || config[0] == unknown {
		result += countValidConfigurations(config[1:], nums)
	}

	if config[0] == damaged || config[0] == unknown {
		if nums[0] <= len(config) && !strings.ContainsRune(config[:nums[0]], operational) && (nums[0] == len(config) || config[nums[0]] != damaged) {
			if nums[0] == len(config) {
				if len(nums) == 1 {
					result++
				}
			} else {
				result += countValidConfigurations(config[nums[0]+1:], nums[1:])
			}
		}
	}

	return result
}

func hotSpringsPart1(input []string) int {
	var sum int
	for _, line := range input {
		parts := strings.Fields(line)
		config := parts[0]
		numsStr := strings.Split(parts[1], ",")
		var nums []int
		for _, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		sum += countValidConfigurations(config, nums)
	}
	return sum
}

func hotSpringsPart2(input []string) int {
	entries := parse(input)
	sum := 0
	for _, entry := range entries {
		sum += findAllCombinations(expand(entry))
	}
	return sum
}

type Entry struct {
	record string
	groups []int
}

func expand(entry Entry) Entry {
	var expandedRecord string
	for i := 0; i < len(entry.record)*5; i++ {
		if i != 0 && i%len(entry.record) == 0 {
			expandedRecord = expandedRecord + string(unknown)
		}
		expandedRecord = expandedRecord + string(entry.record[i%len(entry.record)])
	}

	var expandedGroup []int
	for i := 0; i < len(entry.groups)*5; i++ {
		expandedGroup = append(expandedGroup, entry.groups[i%len(entry.groups)])
	}

	return Entry{expandedRecord, expandedGroup}
}

func findAllCombinations(entry Entry) int {
	var cache [][]int
	for i := 0; i < len(entry.record); i++ {
		cache = append(cache, make([]int, len(entry.groups)+1))
		for j := 0; j < len(entry.groups)+1; j++ {
			cache[i][j] = -1
		}
	}

	return findCombinations(0, 0, entry, cache)
}

func findCombinations(i, j int, entry Entry, cache [][]int) int {
	if i >= len(entry.record) {
		if j < len(entry.groups) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	result := 0
	if entry.record[i] == operational {
		result = findCombinations(i+1, j, entry, cache)
	} else {
		if entry.record[i] == unknown {
			result += findCombinations(i+1, j, entry, cache)
		}
		if j < len(entry.groups) {
			count := 0
			for k := i; k < len(entry.record); k++ {
				if count > entry.groups[j] || entry.record[k] == operational || count == entry.groups[j] && entry.record[k] == unknown {
					break
				}
				count += 1
			}

			if count == entry.groups[j] {
				if i+count < len(entry.record) && entry.record[i+count] != damaged {
					result += findCombinations(i+count+1, j+1, entry, cache)
				} else {
					result += findCombinations(i+count, j+1, entry, cache)
				}
			}
		}
	}

	cache[i][j] = result
	return result
}

func parse(input []string) []Entry {
	var entries []Entry
	for _, line := range input {
		parts := strings.Fields(line)
		var group []int
		for _, number := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(number)
			group = append(group, num)
		}
		entries = append(entries, Entry{parts[0], group})
	}

	return entries
}
