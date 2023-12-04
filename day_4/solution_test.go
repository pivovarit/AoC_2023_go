package day4

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	run()
}

func Test_scratchCardsPart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		}}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scratchCardsPart1(tt.args.input); got != tt.want {
				t.Errorf("scratchCardsPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCards(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "example 1", args: args{" 41 48 83 86 17 "}, want: []int{41, 48, 83, 86, 17}},
		{name: "example 2", args: args{"   41  48   83 86   17  "}, want: []int{41, 48, 83, 86, 17}},
		{name: "example 3", args: args{"  "}, want: []int{}},
		{name: "example 4", args: args{""}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseCards(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_roundId(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "example 1", fields: fields{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, want: 1},
		{name: "example 2", fields: fields{"Card   12: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, want: 12},
		{name: "example 3", fields: fields{"Card     1876: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, want: 1876},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				input: tt.fields.input,
			}
			if got := p.roundId(); got != tt.want {
				t.Errorf("roundId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_winningNumbers(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{name: "example 1", fields: fields{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, want: []int{41, 48, 83, 86, 17}},
		{name: "example 2", fields: fields{"Card 12: 1   2   3  4   5 | 83 86  6 31 17  9 48 53"}, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				input: tt.fields.input,
			}
			if got := p.winningNumbers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("winningNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_chosenNumbers(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{name: "example 1", fields: fields{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, want: []int{83, 86, 6, 31, 17, 9, 48, 53}},
		{name: "example 2", fields: fields{"Card 12: 1   2   3  4   5 | 10 11  13 15 17  90 12 12787"}, want: []int{10, 11, 13, 15, 17, 90, 12, 12787}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				input: tt.fields.input,
			}
			if got := p.chosenNumbers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chosenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scratchCardsPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		}}, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scratchCardsPart2(tt.args.input); got != tt.want {
				t.Errorf("scratchCardsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
