package main_test

import (
	"testing"

	main "github.com/mikelorant/advent-of-code-2024/day14"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	tests := map[string]struct {
		file string
		part int
		x    int
		y    int
		want int
	}{
		"part1 demo1": {
			file: "demo1.txt",
			part: 1,
			x:    11,
			y:    7,
			want: 12,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := main.Task(tt.file, tt.part, tt.x, tt.y)

			assert.Equal(t, tt.want, got)
		})
	}
}
