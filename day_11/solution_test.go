package day11

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkCosmicExpansionPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cosmicExpansionPart1(input)
	}
}

func BenchmarkCosmicExpansionPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cosmicExpansionPart2(input)
	}
}

func Test_cosmicExpansionPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}, want: 374},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cosmicExpansionPart1(tt.input); got != tt.want {
				t.Errorf("cosmicExpansionPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateExpandedPath(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		n     int
		want  int
	}{
		{name: "example", input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}, n: 10, want: 1030},
		{name: "example 2", input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}, n: 100, want: 8410},
		{name: "example 3", input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}, n: 1000, want: 82210},
		{name: "example 4", input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}, n: 10000, want: 820210},
		{name: "example 5", input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}, n: 100000, want: 8200210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateExpandedPath(tt.input, tt.n); got != tt.want {
				t.Errorf("cosmicExpansionPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
