package day18

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkLavaductLagoonPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lavaductLagoonPart1(input)
	}
}

func BenchmarkLavaductLagoonPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lavaductLagoonPart2(input)
	}
}

func Test_lavaductLagoonPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"R 6 (#70c710)",
			"D 5 (#0dc571)",
			"L 2 (#5713f0)",
			"D 2 (#d2c081)",
			"R 2 (#59c680)",
			"D 2 (#411b91)",
			"L 5 (#8ceee2)",
			"U 2 (#caa173)",
			"L 1 (#1b58a2)",
			"U 2 (#caa171)",
			"R 2 (#7807d2)",
			"U 3 (#a77fa3)",
			"L 2 (#015232)",
			"U 2 (#7a21e3)",
		}, want: 62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lavaductLagoonPart1(tt.input); got != tt.want {
				t.Errorf("lavaductLagoonPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lavaductLagoonPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"R 6 (#70c710)",
			"D 5 (#0dc571)",
			"L 2 (#5713f0)",
			"D 2 (#d2c081)",
			"R 2 (#59c680)",
			"D 2 (#411b91)",
			"L 5 (#8ceee2)",
			"U 2 (#caa173)",
			"L 1 (#1b58a2)",
			"U 2 (#caa171)",
			"R 2 (#7807d2)",
			"U 3 (#a77fa3)",
			"L 2 (#015232)",
			"U 2 (#7a21e3)",
		}, want: 952408144115},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lavaductLagoonPart2(tt.input); got != tt.want {
				t.Errorf("lavaductLagoonPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
