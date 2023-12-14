package day2

import (
	"github.com/pivovarit/aoc/util"
	"reflect"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkCubeConundrumPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cubeConundrumPart1(input)
	}
}

func BenchmarkCubeConundrumPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cubeConundrumPart2(input)
	}
}

func Test_cubeConundrum(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Game 1", args: args{[]string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}}, want: 1},
		{name: "Game 2", args: args{[]string{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"}}, want: 2},
		{name: "Game 3", args: args{[]string{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"}}, want: 0},
		{name: "Game 4", args: args{[]string{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"}}, want: 0},
		{name: "Game 5", args: args{[]string{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}}, want: 5},
		{name: "example", args: args{[]string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cubeConundrumPart1(tt.args.input); got != tt.want {
				t.Errorf("cubeConundrumPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cubeConundrumPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Game 1", args: args{[]string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}}, want: 48},
		{name: "Game 2", args: args{[]string{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"}}, want: 12},
		{name: "Game 3", args: args{[]string{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"}}, want: 1560},
		{name: "Game 4", args: args{[]string{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"}}, want: 630},
		{name: "Game 5", args: args{[]string{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}}, want: 36},
		{name: "example", args: args{[]string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}}, want: 2286},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cubeConundrumPart2(tt.args.input); got != tt.want {
				t.Errorf("cubeConundrumPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseEntry(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		want       int
		wantRounds []GameRound
	}{
		{
			name: "Game 1",
			args: args{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			want: 1,
			wantRounds: []GameRound{
				Round(Entry(3, BLUE), Entry(4, RED)),
				Round(Entry(1, RED), Entry(2, GREEN), Entry(6, BLUE)),
				Round(Entry(2, GREEN)),
			}},
		{
			name: "Game 2",
			args: args{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"},
			want: 2,
			wantRounds: []GameRound{
				Round(Entry(1, BLUE), Entry(2, GREEN)),
				Round(Entry(3, GREEN), Entry(4, BLUE), Entry(1, RED)),
				Round(Entry(1, GREEN), Entry(1, BLUE)),
			}},
		{
			name: "Game 3",
			args: args{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"},
			want: 3,
			wantRounds: []GameRound{
				Round(Entry(8, GREEN), Entry(6, BLUE), Entry(20, RED)),
				Round(Entry(5, BLUE), Entry(4, RED), Entry(13, GREEN)),
				Round(Entry(5, GREEN), Entry(1, RED)),
			}},
		{
			name: "Game 4",
			args: args{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"},
			want: 4,
			wantRounds: []GameRound{
				Round(Entry(1, GREEN), Entry(3, RED), Entry(6, BLUE)),
				Round(Entry(3, GREEN), Entry(6, RED)),
				Round(Entry(3, GREEN), Entry(15, BLUE), Entry(14, RED)),
			}},
		{
			name: "Game 5",
			args: args{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"},
			want: 5,
			wantRounds: []GameRound{
				Round(Entry(6, RED), Entry(1, BLUE), Entry(3, GREEN)),
				Round(Entry(2, BLUE), Entry(1, RED), Entry(2, GREEN)),
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseEntry(tt.args.input)
			if got != tt.want {
				t.Errorf("parseEntry() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantRounds) {
				t.Errorf("parseEntry() got = %v, want %v", got1, tt.wantRounds)
			}
		})
	}
}
