package main_test

import (
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc"
	main "github.com/mikelorant/advent-of-code-2024/day10"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	tests := map[string]struct {
		topography string
		part       int
		want       int
	}{
		"part1 demo1": {
			topography: heredoc.Doc(`
				0123
				1234
				8765
				9876
			`),
			part: 1,
			want: 1,
		},
		"part1 demo2": {
			topography: heredoc.Doc(`
				...0...
				...1...
				...2...
				6543456
				7.....7
				8.....8
				9.....9
			`),
			part: 1,
			want: 2,
		},
		"part1 demo3": {
			topography: heredoc.Doc(`
				..90..9
				...1.98
				...2..7
				6543456
				765.987
				876....
				987....
			`),
			part: 1,
			want: 4,
		},
		"part1 demo4": {
			topography: heredoc.Doc(`
				10..9..
				2...8..
				3...7..
				4567654
				...8..3
				...9..2
				.....01
			`),
			part: 1,
			want: 3,
		},
		"part1 example1": {
			topography: heredoc.Doc(`
				89010123
				78121874
				87430965
				96549874
				45678903
				32019012
				01329801
				10456732
			`),
			part: 1,
			want: 36,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := main.Task(strings.NewReader(tt.topography), tt.part)

			assert.Equal(t, tt.want, got)
		})
	}
}
