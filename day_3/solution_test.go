package day3

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	run()
}

func Test_gearRatiosPart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{[]string{
				"467..114..",
				"...*......",
				"..35...633",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*...*",
				".664...598",
			}},
			want: 4361,
		},
		{
			name: "0 parts",
			args: args{[]string{
				"467..114..",
			}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gearRatiosPart1(tt.args.input); got != tt.want {
				t.Errorf("gearRatiosPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gearRatiosPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		}}, want: 467835},
		{name: "example 2", args: args{[]string{
			"467.......",
			".*........",
			"35.......",
			"..*......",
			"..755.....",
			".....*....",
			".....598..",
		}}, want: 467*35 + 35*755 + 755*598},
		{name: "example 3", args: args{[]string{
			"12.......*..",
			"+.........34",
			".......-12..",
			"..78........",
			"..*....60...",
			"78..........",
			".......23...",
			"....90*12...",
			"............",
			"2.2......12.",
			".*.........*",
			"1.1.......56",
		}}, want: 6756},
		{name: "edge case #1", args: args{[]string{
			".........2",
			"........*.",
			".........2",
		}}, want: 4},
		{name: "edge case #2", args: args{[]string{
			"20........",
			".*........",
			"20........",
		}}, want: 400},
		{name: "edge case #3", args: args{[]string{
			"2*20......",
		}}, want: 40},
		{name: "edge case #4", args: args{[]string{
			"......20*2",
		}}, want: 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gearRatiosPart2(tt.args.input); got != tt.want {
				t.Errorf("gearRatiosPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSymbol(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "a", args: args{'a'}, want: true},
		{name: "1", args: args{'1'}, want: false},
		{name: ".", args: args{'.'}, want: false},
		{name: "#", args: args{'#'}, want: true},
		{name: "+", args: args{'+'}, want: true},
		{name: "*", args: args{'*'}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymbol(uint8(tt.args.char)); got != tt.want {
				t.Errorf("isSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractParts(t *testing.T) {
	type args struct {
		entry string
	}
	tests := []struct {
		name string
		args args
		want []GearPart
	}{
		{name: "no parts", args: args{"...*......"}, want: nil},
		{name: "single part", args: args{"...45....."}, want: []GearPart{
			{
				number: 45,
				start:  3,
				end:    5,
			},
		}},
		{name: "multiple parts", args: args{"...45..678"}, want: []GearPart{
			{
				number: 45,
				start:  3,
				end:    5,
			},
			{
				number: 678,
				start:  7,
				end:    10,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractParts(tt.args.entry, 0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
