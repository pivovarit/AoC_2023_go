package day25

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkSnowverloadPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		snowverloadPart1(input)
	}
}

func Test_snowverloadPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"jqt: rhn xhk nvd",
			"rsh: frs pzl lsr",
			"xhk: hfx",
			"cmg: qnr nvd lhk bvb",
			"rhn: xhk bvb hfx",
			"bvb: xhk hfx",
			"pzl: lsr hfx nvd",
			"qnr: nvd",
			"ntq: jqt hfx bvb xhk",
			"nvd: lhk",
			"lsr: lhk",
			"rzs: qnr cmg lsr rsh",
			"frs: qnr lhk lsr",
		}, want: 54},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snowverloadPart1(tt.input); got != tt.want {
				t.Errorf("snowverloadPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
