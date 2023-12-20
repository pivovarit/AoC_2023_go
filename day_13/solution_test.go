package day13

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkPointOfIncidencePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pointOfIncidencePart1(input)
	}
}

func BenchmarkPointOfIncidencePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pointOfIncidencePart2(input)
	}
}

func Test_pointOfIncidencePart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"#.##..##.",
			"..#.##.#.",
			"##......#",
			"##......#",
			"..#.##.#.",
			"..##..##.",
			"#.#.##.#.",
			"",
			"#...##..#",
			"#....#..#",
			"..##..###",
			"#####.##.",
			"#####.##.",
			"..##..###",
			"#....#..#",
		}, want: 405},
		{name: "example", input: []string{
			"##..######..#",
			"###.##..##.##",
			"###..####..##",
			"...#.#..#.#..",
			"..##.#..#.##.",
			"###.#.##.#.##",
			"##....##....#",
			"####.####.###",
			"...########..",
			"..#.#....#.#.",
			"###..####..##",
			"..##......##.",
			"###.#.##.#.##",
			"####..##..###",
			"#..#......#..",
			"...##.##.##..",
			"###.##..##.##",
		}, want: 7},
		{name: "example", input: []string{
			"###......####",
			"..##....##...",
			".#.######.#..",
			"#.###..###.##",
			"#.##.....#.##",
			".#.##..##.#..",
			".##......##..",
		}, want: 12},
		{name: "example", input: []string{
			"#.##..##.",
			"..#.##.#.",
			"##......#",
			"##......#",
			"..#.##.#.",
			"..##..##.",
			"#.#.##.#.",
			"",
			"#...##..#",
			"#....#..#",
			"..##..###",
			"#####.##.",
			"#####.##.",
			"..##..###",
			"#....#..#",
			"",
			".#.##.#.#",
			".##..##..",
			".#.##.#..",
			"#......##",
			"#......##",
			".#.##.#..",
			".##..##.#",
			"",
			"#..#....#",
			"###..##..",
			".##.#####",
			".##.#####",
			"###..##..",
			"#..#....#",
			"#..##...#",
		}, want: 709},
		{name: "example", input: []string{
			"###.##.##",
			"##.####.#",
			"##.#..#.#",
			"####..###",
			"....##...",
			"##.#..#.#",
			"...#..#..",
			"##..###.#",
			"##......#",
			"##......#",
			"..#.##.#.",
			"...#..#..",
			"##.####.#",
			"....##...",
			"...####..",
			"....##...",
			"##.####.#",
			"",
			".##.##...##...##.",
			"#####..##..##..##",
			".....##..##..##..",
			".##.#.#.####.#.#.",
			".##...#.#..#.#...",
			"....#..........#.",
			"#..#..#......#..#",
			"....###.....####.",
			".##...#.#..#.#...",
			".....#..####..#..",
			"#..#...##..##...#",
			"....#...#..#...#.",
			"#..#.##########.#",
			"#..##...####...##",
			"#####.##.##.##.##",
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointOfIncidencePart1(tt.input); got != tt.want {
				t.Errorf("pointOfIncidencePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_colEquals(t *testing.T) {
	type args struct {
		c1      int
		c2      int
		pattern Pattern
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "example", args: args{
			c1: 0,
			c2: 1,
			pattern: []string{
				"#.##..##.",
				"..#.##.#.",
				"##..#...#",
				"##...#..#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
		}, want: false},
		{name: "example", args: args{
			c1: 1,
			c2: 2,
			pattern: []string{
				"#.##..##.",
				"..#.##.#.",
				"##..#...#",
				"##...#..#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colEquals(tt.args.c1, tt.args.c2, tt.args.pattern); got != tt.want {
				t.Errorf("colEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pointOfIncidencePart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"#.##..##.",
			"..#.##.#.",
			"##......#",
			"##......#",
			"..#.##.#.",
			"..##..##.",
			"#.#.##.#.",
			"",
			"#...##..#",
			"#....#..#",
			"..##..###",
			"#####.##.",
			"#####.##.",
			"..##..###",
			"#....#..#",
			"",
			".#.##.#.#",
			".##..##..",
			".#.##.#..",
			"#......##",
			"#......##",
			".#.##.#..",
			".##..##.#",
			"",
			"#..#....#",
			"###..##..",
			".##.#####",
			".##.#####",
			"###..##..",
			"#..#....#",
			"#..##...#",
		}, want: 1400},
		{name: "example", input: []string{
			"#.##..##.",
			"..#.##.#.",
			"##......#",
			"##......#",
			"..#.##.#.",
			"..##..##.",
			"#.#.##.#.",
			"",
			"#...##..#",
			"#....#..#",
			"..##..###",
			"#####.##.",
			"#####.##.",
			"..##..###",
			"#....#..#",
		}, want: 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointOfIncidencePart2(tt.input); got != tt.want {
				t.Errorf("pointOfIncidencePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
