package main_test

import (
	"strings"
	"testing"

	main "github.com/mikelorant/advent-of-code-2024/day11"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	tests := map[string]struct {
		input  string
		blinks int
		part   int
		want   int
	}{
		"part1 demo1": {
			input:  "0 1 10 99 999",
			blinks: 1,
			part:   1,
			want:   7,
		},
		"part1 demo2 blink1": {
			input:  "125 17",
			blinks: 1,
			part:   1,
			want:   3,
		},
		"part1 demo2 blink2": {
			input:  "125 17",
			blinks: 2,
			part:   1,
			want:   4,
		},
		"part1 demo2 blink3": {
			input:  "125 17",
			blinks: 3,
			part:   1,
			want:   5,
		},
		"part1 demo2 blink4": {
			input:  "125 17",
			blinks: 4,
			part:   1,
			want:   9,
		},
		"part1 demo2 blink5": {
			input:  "125 17",
			blinks: 5,
			part:   1,
			want:   13,
		},
		"part1 demo2 blink6": {
			input:  "125 17",
			blinks: 6,
			part:   1,
			want:   22,
		},
		"part1 demo2 blink25": {
			input:  "125 17",
			blinks: 25,
			part:   1,
			want:   55312,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := main.Task(strings.NewReader(tt.input), tt.blinks, tt.part)

			assert.Equal(t, tt.want, got)
		})
	}
}
