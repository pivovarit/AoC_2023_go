package day12

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkHotSpringsPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hotSpringsPart1(input)
	}
}

func BenchmarkHotSpringsPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hotSpringsPart2(input)
	}
}

func Test_hotSpringsPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"???.### 1,1,3",
			".??..??...?##. 1,1,3",
			"?#?#?#?#?#?#?#? 1,3,1,6",
			"????.#...#... 4,1,1",
			"????.######..#####. 1,6,5",
			"?###???????? 3,2,1",
		}, want: 21},
		{name: "example", input: input, want: 7032},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hotSpringsPart1(tt.input); got != tt.want {
				t.Errorf("hotSpringsPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hotSpringsPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"???.### 1,1,3",
			".??..??...?##. 1,1,3",
			"?#?#?#?#?#?#?#? 1,3,1,6",
			"????.#...#... 4,1,1",
			"????.######..#####. 1,6,5",
			"?###???????? 3,2,1",
		}, want: 525152},
		{name: "example", input: input, want: 1493340882140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hotSpringsPart2(tt.input); got != tt.want {
				t.Errorf("hotSpringsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
