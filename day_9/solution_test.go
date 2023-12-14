package day9

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkMirageMaintenancePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mirageMaintenancePart1(input)
	}
}

func BenchmarkMirageMaintenancePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mirageMaintenancePart2(input)
	}
}

func Test_mirageMaintenancePart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
		}, want: 114},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirageMaintenancePart1(tt.input); got != tt.want {
				t.Errorf("mirageMaintenancePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mirageMaintenancePart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
		}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirageMaintenancePart2(tt.input); got != tt.want {
				t.Errorf("mirageMaintenancePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
