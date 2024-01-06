package day22

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkSandSlabsPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sandSlabsPart1(input)
	}
}

func BenchmarkSandSlabsPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sandSlabsPart2(input)
	}
}

func Test_sandSlabsPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"1,0,1~1,2,1",
			"0,0,2~2,0,2",
			"0,2,3~2,2,3",
			"0,0,4~0,2,4",
			"2,0,5~2,2,5",
			"0,1,6~2,1,6",
			"1,1,8~1,1,9",
		}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sandSlabsPart1(tt.input); got != tt.want {
				t.Errorf("sandSlabsPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sandSlabsPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"1,0,1~1,2,1",
			"0,0,2~2,0,2",
			"0,2,3~2,2,3",
			"0,0,4~0,2,4",
			"2,0,5~2,2,5",
			"0,1,6~2,1,6",
			"1,1,8~1,1,9",
		}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sandSlabsPart2(tt.input); got != tt.want {
				t.Errorf("sandSlabsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
