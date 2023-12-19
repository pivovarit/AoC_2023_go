package day19

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkAplentyPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aplentyPart1(input)
	}
}

func BenchmarkAplentyPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aplentyPart2(input)
	}
}

func Test_aplentyPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: input, want: 432427},
		{name: "example", input: []string{
			"px{a<2006:qkq,m>2090:A,rfg}",
			"pv{a>1716:R,A}",
			"lnx{m>1548:A,A}",
			"rfg{s<537:gd,x>2440:R,A}",
			"qs{s>3448:A,lnx}",
			"qkq{x<1416:A,crn}",
			"crn{x>2662:A,R}",
			"in{s<1351:px,qqz}",
			"qqz{s>2770:qs,m<1801:hdj,R}",
			"gd{a>3333:R,R}",
			"hdj{m>838:A,pv}",
			"",
			"{x=787,m=2655,a=1222,s=2876}",
			"{x=1679,m=44,a=2067,s=496}",
			"{x=2036,m=264,a=79,s=2244}",
			"{x=2461,m=1339,a=466,s=291}",
			"{x=2127,m=1623,a=2188,s=1013}",
		}, want: 19114},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aplentyPart1(tt.input); got != tt.want {
				t.Errorf("aplentyPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aplentyPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: input, want: 143760172569135},
		{name: "example", input: []string{
			"px{a<2006:qkq,m>2090:A,rfg}",
			"pv{a>1716:R,A}",
			"lnx{m>1548:A,A}",
			"rfg{s<537:gd,x>2440:R,A}",
			"qs{s>3448:A,lnx}",
			"qkq{x<1416:A,crn}",
			"crn{x>2662:A,R}",
			"in{s<1351:px,qqz}",
			"qqz{s>2770:qs,m<1801:hdj,R}",
			"gd{a>3333:R,R}",
			"hdj{m>838:A,pv}",
		}, want: 167409079868000},
		{name: "example", input: []string{
			"px1{a<2006:qkq1,px2}",
			"px2{m>2090:A,rfg1}",
			"pv1{a>1716:R,A}",
			"lnx1{m>1548:A,A}",
			"rfg1{s<537:gd1,rfg2}",
			"rfg2{x>2440:R,A}",
			"qs1{s>3448:A,lnx1}",
			"qkq1{x<1416:A,crn1}",
			"crn1{x>2662:A,R}",
			"in{s<1351:px1,qqz1}",
			"qqz1{s>2770:qs1,qqz2}",
			"qqz2{m<1801:hdj1,R}",
			"gd1{a>3333:R,R}",
			"hdj1{m>838:A,pv1}",
		}, want: 167409079868000},
		{name: "all combinations", input: []string{"in{a<9000:A,A}"}, want: 256000000000000},
		{name: "example", input: []string{
			"in{x>1:in2,R}",
			"in2{m>1:in3,A}",
			"in3{a>1:in4,A}",
			"in4{s>1:A,R}",
		}, want: 255872047988001},
		{name: "example", input: []string{
			"in{x<4001:aaa,R}",
			"aaa{x<1001:R,m<2001:R,A}",
		}, want: 96000000000000},
		{name: "example", input: []string{"in{x<2001:R,a<2001:A,x<3001:R,A}"}, want: 96000000000000},
		{name: "single combination", input: []string{
			"in{x<2:in2,R}",
			"in2{m<2:in3,R}",
			"in3{a<2:in4,R}",
			"in4{s<2:A,R}",
		}, want: 1},
		{name: "edge case", input: []string{
			"in{x>3:R,x<2:R,A}",
		}, want: 128000000000},
		{name: "edge case", input: []string{
			"in{m>3:R,m<2:R,A}",
		}, want: 128000000000},
		{name: "edge case", input: []string{
			"in{a>3:R,a<2:R,A}",
		}, want: 128000000000},
		{name: "edge case", input: []string{
			"in{s>3:R,s<2:R,A}",
		}, want: 128000000000},
		{name: "edge case", input: []string{
			"in{x<2001:A,R}",
		}, want: 128000000000000},
		{name: "edge case", input: []string{
			"in{m<2001:A,R}",
		}, want: 128000000000000},
		{name: "edge case", input: []string{
			"in{a<2001:A,R}",
		}, want: 128000000000000},
		{name: "edge case", input: []string{
			"in{s<2001:A,R}",
		}, want: 128000000000000},
		{name: "edge case", input: []string{
			"in{x<4000:R,m<4000:R,a<4000:R,s<4000:R,A}",
		}, want: 1},
		{name: "edge case", input: []string{
			"in{x>2:R,x>1:R,m>1:R,a>1:R,s>1:R,A}",
		}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aplentyPart2(tt.input); got != tt.want {
				t.Errorf("aplentyPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
