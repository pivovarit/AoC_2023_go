package day23

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkALongWalkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aLongWalkPart1(input)
	}
}

func BenchmarkALongWalkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aLongWalkPart2(input)
	}
}

func Test_aLongWalkPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"#.#####################",
			"#.......#########...###",
			"#######.#########.#.###",
			"###.....#.>.>.###.#.###",
			"###v#####.#v#.###.#.###",
			"###.>...#.#.#.....#...#",
			"###v###.#.#.#########.#",
			"###...#.#.#.......#...#",
			"#####.#.#.#######.#.###",
			"#.....#.#.#.......#...#",
			"#.#####.#.#.#########v#",
			"#.#...#...#...###...>.#",
			"#.#.#v#######v###.###v#",
			"#...#.>.#...>.>.#.###.#",
			"#####v#.#.###v#.#.###.#",
			"#.....#...#...#.#.#...#",
			"#.#########.###.#.#.###",
			"#...###...#...#...#.###",
			"###.###.#.###v#####v###",
			"#...#...#.#.>.>.#.>.###",
			"#.###.###.#.###.#.#v###",
			"#.....###...###...#...#",
			"#####################.#",
		}, want: 94},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aLongWalkPart1(tt.input); got != tt.want {
				t.Errorf("aLongWalkPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aLongWalkPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"#.#####################",
			"#.......#########...###",
			"#######.#########.#.###",
			"###.....#.>.>.###.#.###",
			"###v#####.#v#.###.#.###",
			"###.>...#.#.#.....#...#",
			"###v###.#.#.#########.#",
			"###...#.#.#.......#...#",
			"#####.#.#.#######.#.###",
			"#.....#.#.#.......#...#",
			"#.#####.#.#.#########v#",
			"#.#...#...#...###...>.#",
			"#.#.#v#######v###.###v#",
			"#...#.>.#...>.>.#.###.#",
			"#####v#.#.###v#.#.###.#",
			"#.....#...#...#.#.#...#",
			"#.#########.###.#.#.###",
			"#...###...#...#...#.###",
			"###.###.#.###v#####v###",
			"#...#...#.#.>.>.#.>.###",
			"#.###.###.#.###.#.#v###",
			"#.....###...###...#...#",
			"#####################.#",
		}, want: 154},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aLongWalkPart2(tt.input); got != tt.want {
				t.Errorf("aLongWalkPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
