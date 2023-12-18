package day17

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkClumsyCruciblePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clumsyCruciblePart1(input)
	}
}

func BenchmarkClumsyCruciblePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clumsyCruciblePart2(input)
	}
}

func Test_clumsyCruciblePart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"2413432311323",
			"3215453535623",
			"3255245654254",
			"3446585845452",
			"4546657867536",
			"1438598798454",
			"4457876987766",
			"3637877979653",
			"4654967986887",
			"4564679986453",
			"1224686865563",
			"2546548887735",
			"4322674655533",
		}, want: 102},
		{name: "2x2", input: []string{
			"12",
			"31",
		}, want: 3},
		{name: "3x3", input: []string{
			"123",
			"552",
			"553",
		}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clumsyCruciblePart1(tt.input); got != tt.want {
				t.Errorf("clumsyCruciblePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clumsyCruciblePart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"2413432311323",
			"3215453535623",
			"3255245654254",
			"3446585845452",
			"4546657867536",
			"1438598798454",
			"4457876987766",
			"3637877979653",
			"4654967986887",
			"4564679986453",
			"1224686865563",
			"2546548887735",
			"4322674655533",
		}, want: 94},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clumsyCruciblePart2(tt.input); got != tt.want {
				t.Errorf("clumsyCruciblePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
