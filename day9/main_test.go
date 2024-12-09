package main_test

import (
	"strings"
	"testing"

	main "github.com/mikelorant/advent-of-code-2024/day9"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	tests := map[string]struct {
		diskmap string
		part    int
		want    int
	}{
		"part1 demo1": {
			diskmap: "12345",
			part:    1,
			want:    60,
		},
		"part1 example1": {
			diskmap: "2333133121414131402",
			part:    1,
			want:    1928,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := main.Task(strings.NewReader(tt.diskmap), tt.part)

			assert.Equal(t, tt.want, got)
		})
	}
}
