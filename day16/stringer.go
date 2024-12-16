package main

import (
	"fmt"
	"image"
	"strings"
)

func (m Maze) String() string {
	var strs []string

	w, h := m.dimensions()

	var tiles [][]string

	for range h {
		tiles = append(tiles, make([]string, w))
	}

	for pt, t := range m {
		tiles[pt.Y][pt.X] = fmt.Sprint(t)
	}

	for _, row := range tiles {
		strs = append(strs, strings.Join(row, ""))
	}

	return strings.Join(strs, "\n")
}

func (t Tile) String() string {
	switch t.Type {
	case Empty:
		return "."
	case End:
		return "E"
	case Start:
		return "S"
	case Wall:
		return "#"
	default:
		return "?"
	}
}

func (m Maze) dimensions() (w, h int) {
	if len(m) == 0 {
		return 0, 0
	}

	maxPt := image.Point{X: 0, Y: 0}

	for pt := range m {
		if pt.X > maxPt.X {
			maxPt.X = pt.X
		}

		if pt.Y > maxPt.Y {
			maxPt.Y = pt.Y
		}
	}

	return maxPt.X + 1, maxPt.Y + 1
}
