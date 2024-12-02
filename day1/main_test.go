package main_test

import (
	"testing"

	main "github.com/mikelorant/advent-of-code-2024/day1"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	tests := map[string]struct {
		file string
		part int
		want int
	}{
		"part1 demo1": {
			file: "demo1.txt",
			part: 1,
			want: 11,
		},
		"part2 demo1": {
			file: "demo1.txt",
			part: 2,
			want: 31,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := main.Task(tt.file, tt.part)

			assert.Equal(t, tt.want, got)
		})
	}
}