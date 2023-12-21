package day16

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkTheFloorWillBeLavaPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		theFloorWillBeLavaPart1(input)
	}
}

func BenchmarkTheFloorWillBeLavaPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		theFloorWillBeLavaPart2(input)
	}
}

func Test_theFloorWillBeLavaPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			`.|...\....`,
			`|.-.\.....`,
			`.....|-...`,
			`........|.`,
			`..........`,
			`.........\`,
			`..../.\\..`,
			`.-.-/..|..`,
			`.|....-|.\`,
			`..//.|....`,
		}, want: 46},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := theFloorWillBeLavaPart1(tt.input); got != tt.want {
				t.Errorf("theFloorWillBeLavaPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_theFloorWillBeLavaPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			`.|...\....`,
			`|.-.\.....`,
			`.....|-...`,
			`........|.`,
			`..........`,
			`.........\`,
			`..../.\\..`,
			`.-.-/..|..`,
			`.|....-|.\`,
			`..//.|....`,
		}, want: 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := theFloorWillBeLavaPart2(tt.input); got != tt.want {
				t.Errorf("theFloorWillBeLavaPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
