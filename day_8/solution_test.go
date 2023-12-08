package day8

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	run()
}

func Test_hauntedWasteLandPart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example 1", args: args{[]string{
			"RL",
			"",
			"AAA = (BBB, CCC)",
			"BBB = (DDD, EEE)",
			"CCC = (ZZZ, GGG)",
			"DDD = (DDD, DDD)",
			"EEE = (EEE, EEE)",
			"GGG = (GGG, GGG)",
			"ZZZ = (ZZZ, ZZZ)",
		}}, want: 2},
		{name: "example 2", args: args{[]string{
			"LLR",
			"",
			"AAA = (BBB, BBB)",
			"BBB = (AAA, ZZZ)",
			"ZZZ = (ZZZ, ZZZ)",
		}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hauntedWasteLandPart1(tt.args.input); got != tt.want {
				t.Errorf("hauntedWasteLandPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_parse(t *testing.T) {
	type fields struct {
		input []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  map[string]Directions
	}{
		{name: "example 1", fields: fields{[]string{
			"LLR",
			"",
			"AAA = (BBB, BBB)",
			"BBB = (AAA, ZZZ)",
			"ZZZ = (ZZZ, ZZZ)",
		}}, want: "LLR", want1: map[string]Directions{
			"AAA": {"BBB", "BBB"},
			"BBB": {"AAA", "ZZZ"},
			"ZZZ": {"ZZZ", "ZZZ"},
		}},
		{name: "example 2", fields: fields{[]string{
			"RL",
			"",
			"AAA = (BBB, CCC)",
			"BBB = (DDD, EEE)",
			"CCC = (ZZZ, GGG)",
			"DDD = (DDD, DDD)",
			"EEE = (EEE, EEE)",
			"GGG = (GGG, GGG)",
			"ZZZ = (ZZZ, ZZZ)",
		}}, want: "RL", want1: map[string]Directions{
			"AAA": {"BBB", "CCC"},
			"BBB": {"DDD", "EEE"},
			"CCC": {"ZZZ", "GGG"},
			"DDD": {"DDD", "DDD"},
			"EEE": {"EEE", "EEE"},
			"GGG": {"GGG", "GGG"},
			"ZZZ": {"ZZZ", "ZZZ"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parse(tt.fields.input)
			if got != tt.want {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_hauntedWasteLandPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"LR",
			"",
			"11A = (11B, XXX)",
			"11B = (XXX, 11Z)",
			"11Z = (11B, XXX)",
			"22A = (22B, XXX)",
			"22B = (22C, 22C)",
			"22C = (22Z, 22Z)",
			"22Z = (22B, 22B)",
			"XXX = (XXX, XXX)",
		}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hauntedWasteLandPart2(tt.args.input); got != tt.want {
				t.Errorf("hauntedWasteLandPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
