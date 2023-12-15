package day15

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkLensLibraryPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lensLibraryPart1(input)
	}
}

func BenchmarkLensLibraryPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lensLibraryPart2(input)
	}
}

func Test_lensLibraryPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{input: []string{"rn=1"}, want: 30},
		{input: []string{"cm-"}, want: 253},
		{input: []string{"qp=3"}, want: 97},
		{input: []string{"cm=2"}, want: 47},
		{input: []string{"qp-"}, want: 14},
		{input: []string{"pc=4"}, want: 180},
		{input: []string{"ot=9"}, want: 9},
		{input: []string{"ab=5"}, want: 197},
		{input: []string{"pc-"}, want: 48},
		{input: []string{"pc=6"}, want: 214},
		{input: []string{"ot=7"}, want: 231},
		{input: []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"}, want: 1320},
	}
	for _, tt := range tests {
		t.Run(tt.input[0], func(t *testing.T) {
			if got := lensLibraryPart1(tt.input); got != tt.want {
				t.Errorf("lensLibraryPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lensLibraryPart2(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{input: []string{"rn=1,rn-"}, want: 0},
		{input: []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"}, want: 145},
	}
	for _, tt := range tests {
		t.Run(tt.input[0], func(t *testing.T) {
			if got := lensLibraryPart2(tt.input); got != tt.want {
				t.Errorf("lensLibraryPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
