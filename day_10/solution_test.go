package day10

import (
	"testing"
)

func Test_run(t *testing.T) {
	run()
}

func Test_pipeMazePart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			".....",
			".F-7.",
			".S.|.",
			".L-J.",
			".....",
		}, want: 4},
		{name: "example", input: []string{
			".....",
			".F-7.",
			".|.|.",
			".S-J.",
			".....",
		}, want: 4},
		{name: "example", input: []string{
			".....",
			".F-7.",
			".|.|.",
			".L-S.",
			".....",
		}, want: 4},
		{name: "example", input: []string{
			".....",
			".F-7.",
			".|.S.",
			".L-J.",
			".....",
		}, want: 4},
		{name: "example", input: []string{
			".....",
			".F-S.",
			".|.|.",
			".L-J.",
			".....",
		}, want: 4},
		{name: "example", input: []string{
			".....",
			".FS7.",
			".|.|.",
			".L-J.",
			".....",
		}, want: 4},
		{name: "example", input: []string{
			".....",
			".F-7.",
			".|.|.",
			".LSJ.",
			".....",
		}, want: 4},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"SJ.L7",
			"|F--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FS.L7",
			"|F--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".SJ|.",
			"FJ.L7",
			"|F--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FS|.",
			"FJ.L7",
			"|F--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..S7.",
			".FJ|.",
			"FJ.L7",
			"|F--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..FS.",
			".FJ|.",
			"FJ.L7",
			"|F--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJS.",
			"FJ.L7",
			"|F--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FJ.S7",
			"|F--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FJ.LS",
			"|F--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FJ.L7",
			"|F--S",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FJ.L7",
			"|F-SJ",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FJ.L7",
			"|FS-J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FJ.L7",
			"|S--J",
			"LJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FJ.L7",
			"|F--J",
			"LS...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FJ.L7",
			"|F--J",
			"SJ...",
		}, want: 8},
		{name: "example 2", input: []string{
			"..F7.",
			".FJ|.",
			"FJ.L7",
			"SF--J",
			"LJ...",
		}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pipeMazePart1(tt.input); got != tt.want {
				t.Errorf("pipeMazePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pipeMazePart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "example", input: []string{
			"FF7FSF7F7F7F7F7F---7",
			"L|LJ||||||||||||F--J",
			"FL-7LJLJ||||||LJL-77",
			"F--JF--7||LJLJIF7FJ-",
			"L---JF-JLJIIIIFJLJJ7",
			"|F|F-JF---7IIIL7L|7|",
			"|FFJF7L7F-JF7IIL---7",
			"7-L-JL7||F7|L7F-7F7|",
			"L.L7LFJ|||||FJL7||LJ",
			"L7JLJL-JLJLJL--JLJ.L",
		}, want: 10},
		{name: "example", input: []string{
			"OF----7F7F7F7F-7OOOO",
			"O|F--7||||||||FJOOOO",
			"O||OFJ||||||||L7OOOO",
			"FJL7L7LJLJ||LJIL-7OO",
			"L--JOL7IIILJS7F-7L7O",
			"OOOOF-JIIF7FJ|L7L7L7",
			"OOOOL7IF7||L7|IL7L7|",
			"OOOOO|FJLJ|FJ|F7|OLJ",
			"OOOOFJL-7O||O||||OOO",
			"OOOOL---JOLJOLJLJOOO",
		}, want: 8},
		{name: "example", input: []string{
			"...........",
			".S-------7.",
			".|F-----7|.",
			".||.....||.",
			".||.....||.",
			".|L-7.F-J|.",
			".|..|.|..|.",
			".L--J.L--J.",
			"...........",
		}, want: 4},
		{name: "example", input: []string{
			"..........",
			".S------7.",
			".|F----7|.",
			".||OOOO||.",
			".||OOOO||.",
			".|L-7F-J|.",
			".|II||II|.",
			".L--JL--J.",
			"..........",
		}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pipeMazePart2(tt.input); got != tt.want {
				t.Errorf("pipeMazePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
