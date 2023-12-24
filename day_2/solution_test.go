package day2

import (
	"github.com/pivovarit/aoc/util"
	"testing"
)

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkCubeConundrumPart1Sequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cubeConundrumPart1Sequential(input)
	}
}

func BenchmarkCubeConundrumPart1Parallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cubeConundrumPart1Parallel(input)
	}
}

func BenchmarkCubeConundrumPart2Sequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cubeConundrumPart2Sequential(input)
	}
}

func BenchmarkCubeConundrumPart2Parallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cubeConundrumPart2Parallel(input)
	}
}

func Test_cubeConundrumPart1_sequential(t *testing.T) {
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
			if got := cubeConundrumPart1Sequential(tt.args.input); got != tt.want {
				t.Errorf("cubeConundrumPart1_sequential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cubeConundrumPart2_parallel(t *testing.T) {
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
			if got := cubeConundrumPart1Parallel(tt.args.input); got != tt.want {
				t.Errorf("cubeConundrumPart1_parallel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cubeConundrumPart2Sequential(t *testing.T) {
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
			if got := cubeConundrumPart2Sequential(tt.args.input); got != tt.want {
				t.Errorf("cubeConundrumPart2Sequential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cubeConundrumPart2Parallel(t *testing.T) {
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
			if got := cubeConundrumPart2Parallel(tt.args.input); got != tt.want {
				t.Errorf("cubeConundrumPart2Parallel() = %v, want %v", got, tt.want)
			}
		})
	}
}
