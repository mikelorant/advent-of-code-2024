package main_test

import (
	"strings"
	"testing"

	main "github.com/mikelorant/advent-of-code-2024/day3"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	tests := map[string]struct {
		program string
		part    int
		want    int
	}{
		"part1 example1": {
			program: "mul(44,46)",
			part:    1,
			want:    2024,
		},
		"part1 example2": {
			program: "mul(123,4)",
			part:    1,
			want:    492,
		},
		"part1 example3": {
			program: "mul(4*",
			part:    1,
			want:    0,
		},
		"part1 example4": {
			program: "mul(6,9!",
			part:    1,
			want:    0,
		},
		"part1 example5": {
			program: "?(12,34)",
			part:    1,
			want:    0,
		},
		"part1 example6": {
			program: "mul ( 2 , 4 )",
			part:    1,
			want:    0,
		},
		"part1 example7": {
			program: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			part:    1,
			want:    161,
		},
		"part2 example1": {
			program: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			part:    2,
			want:    48,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := main.Task(strings.NewReader(tt.program), tt.part)

			assert.Equal(t, tt.want, got)
		})
	}
}
