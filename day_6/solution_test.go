package day6

import (
	"reflect"
	"testing"
)

var example = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func Test_run(t *testing.T) {
	run()
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
