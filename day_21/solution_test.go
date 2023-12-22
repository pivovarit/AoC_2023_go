package day21

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkStepCounterPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stepCounterPart1(input)
	}
}

func BenchmarkStepCounterPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stepCounterPart2(input)
	}
}

func Test_stepCounterPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"...........",
			".....###.#.",
			".###.##..#.",
			"..#.#...#..",
			"....#.#....",
			".##..S####.",
			".##..#...#.",
			".......##..",
			".##.#.####.",
			".##..##.##.",
			"...........",
		}, want: 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stepCounterPart1(tt.input); got != tt.want {
				t.Errorf("stepCounterPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traverse(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		steps int
		want  int
	}{
		{name: "example", input: []string{
			"...........",
			".....###.#.",
			".###.##..#.",
			"..#.#...#..",
			"....#.#....",
			".##..S####.",
			".##..#...#.",
			".......##..",
			".##.#.####.",
			".##..##.##.",
			"...........",
		}, steps: 6, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traverse(tt.input, Location{5, 5}, tt.steps, false); got != tt.want {
				t.Errorf("traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stepCounterPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: input, want: 636350496972143},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stepCounterPart2(tt.input); got != tt.want {
				t.Errorf("stepCounterPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
