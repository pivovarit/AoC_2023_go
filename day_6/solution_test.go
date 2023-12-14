package day6

import (
	"github.com/pivovarit/aoc/util"
	"reflect"
	"testing"
)

var example = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

var input = util.ReadInput()

func Test_run(t *testing.T) {
	run()
}

func BenchmarkWaitForItPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		waitForItPart1(input)
	}
}

func BenchmarkWaitForItPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		waitForItPart2(input)
	}
}

func TestParser_parse(t *testing.T) {
	type fields struct {
		input []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []Race
	}{
		{name: "example", fields: fields{example}, want: []Race{
			{time: 7, record: 9},
			{time: 15, record: 40},
			{time: 30, record: 200},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				input: tt.fields.input,
			}
			if got := p.parse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_waitForItPart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{input: example}, want: 288},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waitForItPart1(tt.args.input); got != tt.want {
				t.Errorf("waitForItPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_waitForItPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{input: example}, want: 71503},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waitForItPart2(tt.args.input); got != tt.want {
				t.Errorf("waitForItPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
