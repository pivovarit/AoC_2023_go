package day5

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	run()
}

var example = []string{
	"seeds: 79 14 55 13",
	"",
	"seed-to-soil map:",
	"50 98 2",
	"52 50 48",
	"",
	"soil-to-fertilizer map:",
	"0 15 37",
	"37 52 2",
	"39 0 15",
	"",
	"fertilizer-to-water map:",
	"49 53 8",
	"0 11 42",
	"42 0 7",
	"57 7 4",
	"",
	"water-to-light map:",
	"88 18 7",
	"18 25 70",
	"",
	"light-to-temperature map:",
	"45 77 23",
	"81 45 19",
	"68 64 13",
	"",
	"temperature-to-humidity map:",
	"0 69 1",
	"1 0 69",
	"",
	"humidity-to-location map:",
	"60 56 37",
	"56 93 4",
}

func TestParser_parse(t *testing.T) {
	type fields struct {
		tokenizer Tokenizer
	}
	tests := []struct {
		name   string
		fields fields
		want   State
	}{
		{
			name: "example",
			fields: fields{
				tokenizer: Tokenizer{input: example},
			},
			want: State{
				seeds: []int{79, 14, 55, 13},
				seedToSoil: []AlmanacMap{
					{DestinationRange: Range{start: 52, end: 100}, SourceRange: Range{start: 50, end: 98}},
					{DestinationRange: Range{start: 50, end: 52}, SourceRange: Range{start: 98, end: 100}},
				},
				soilToFertilizer: []AlmanacMap{
					{DestinationRange: Range{start: 39, end: 54}, SourceRange: Range{start: 0, end: 15}},
					{DestinationRange: Range{start: 0, end: 37}, SourceRange: Range{start: 15, end: 52}},
					{DestinationRange: Range{start: 37, end: 39}, SourceRange: Range{start: 52, end: 54}},
				},
				fertilizerToWater: []AlmanacMap{
					{DestinationRange: Range{start: 42, end: 49}, SourceRange: Range{start: 0, end: 7}},
					{DestinationRange: Range{start: 57, end: 61}, SourceRange: Range{start: 7, end: 11}},
					{DestinationRange: Range{start: 0, end: 42}, SourceRange: Range{start: 11, end: 53}},
					{DestinationRange: Range{start: 49, end: 57}, SourceRange: Range{start: 53, end: 61}},
				},
				waterToLight: []AlmanacMap{
					{DestinationRange: Range{start: 88, end: 95}, SourceRange: Range{start: 18, end: 25}},
					{DestinationRange: Range{start: 18, end: 88}, SourceRange: Range{start: 25, end: 95}},
				},
				lightToTemperature: []AlmanacMap{
					{DestinationRange: Range{start: 81, end: 100}, SourceRange: Range{start: 45, end: 64}},
					{DestinationRange: Range{start: 68, end: 81}, SourceRange: Range{start: 64, end: 77}},
					{DestinationRange: Range{start: 45, end: 68}, SourceRange: Range{start: 77, end: 100}},
				},
				temperatureToHumidity: []AlmanacMap{
					{DestinationRange: Range{start: 1, end: 70}, SourceRange: Range{start: 0, end: 69}},
					{DestinationRange: Range{start: 0, end: 1}, SourceRange: Range{start: 69, end: 70}},
				},
				humidityToLocation: []AlmanacMap{
					{DestinationRange: Range{start: 60, end: 97}, SourceRange: Range{start: 56, end: 93}},
					{DestinationRange: Range{start: 56, end: 60}, SourceRange: Range{start: 93, end: 97}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				tokenizer: tt.fields.tokenizer,
			}
			if got := p.parse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_fertilizerPart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{input: example}, want: 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := almanacPart1(tt.args.input); got != tt.want {
				t.Errorf("almanacPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fertilizerPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{input: example}, want: 46},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := almanacPart2(tt.args.input); got != tt.want {
				t.Errorf("almanacPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtract(t *testing.T) {
	type args struct {
		r1     Range
		ranges []Range
	}
	tests := []struct {
		name string
		args args
		want []Range
	}{
		{name: "no_overlap", args: args{r1: Range{1, 5}, ranges: []Range{{7, 10}}}, want: []Range{{1, 5}}},
		{name: "overlap", args: args{r1: Range{1, 10}, ranges: []Range{{7, 15}}}, want: []Range{{1, 7}}},
		{name: "full_overlap", args: args{r1: Range{1, 10}, ranges: []Range{{3, 7}}}, want: []Range{{1, 3}, {7, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtract(tt.args.r1, tt.args.ranges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
