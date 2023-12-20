package day20

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkPulsePropagationPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pulsePropagationPart1(input)
	}
}

func BenchmarkPulsePropagationPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pulsePropagationPart2(input)
	}
}

func Test_pulsePropagationPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: input, want: 929810733},
		{name: "example 1", input: []string{
			"broadcaster -> a, b, c",
			"%a -> b",
			"%b -> c",
			"%c -> inv",
			"&inv -> a",
		}, want: 32000000},
		{name: "example 2", input: []string{
			"broadcaster -> a",
			"%a -> inv, con",
			"&inv -> b",
			"%b -> con",
			"&con -> output",
		}, want: 11687500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pulsePropagationPart1(tt.input); got != tt.want {
				t.Errorf("pulsePropagationPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pulsePropagationPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: input, want: 231657829136023},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pulsePropagationPart2(tt.input); got != tt.want {
				t.Errorf("pulsePropagationPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
