package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample1 string

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "actual",
			input: cleanInput(sample1),
			want:  6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			println("PART 1")
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			} else {
				println("  ", got)
			}
			println("")
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "actual",
			input: cleanInput(sample1),
			want:  5905,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			println("PART 2")
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			} else {
				println("  ", got)
			}
			println("")
		})
	}
}
