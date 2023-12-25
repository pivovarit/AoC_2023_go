package day24

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkNeverTellMeTheOddsPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		neverTellMeTheOddsPart1(input)
	}
}

func BenchmarkNeverTellMeTheOddsPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		neverTellMeTheOddsPart2(input)
	}
}

func Test_neverTellMeTheOddsPart1(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		wantCount int
	}{
		{name: "example", input: input, wantCount: 17244},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := neverTellMeTheOddsPart1(tt.input); gotCount != tt.wantCount {
				t.Errorf("neverTellMeTheOddsPart1() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_neverTellMeTheOddsPart2(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		wantCount int
	}{
		{name: "example", input: input, wantCount: 1025019997186820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := neverTellMeTheOddsPart2(tt.input); gotCount != tt.wantCount {
				t.Errorf("neverTellMeTheOddsPart2() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
