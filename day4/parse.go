package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func parse(fn string, filt []string) (Grid, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	var grid Grid

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		row := toRow(scanner.Text(), filt)

		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return addNeighbours(grid), nil
}

func toRow(str string, filt []string) Row {
	var row Row

	for _, char := range str {
		cell := Cell{
			Value: ".",
		}

		if slices.Contains(filt, string(char)) {
			cell.Value = string(char)
		}

		row = append(row, cell)
	}

	return row
}

func addNeighbours(g Grid) Grid {
	for y, row := range g {
		for x := range row {
			cell := g[y][x]

			cell.Neighbours.TopLeft = getCell(g, x-1, y-1)
			cell.Neighbours.Top = getCell(g, x, y-1)
			cell.Neighbours.TopRight = getCell(g, x+1, y-1)
			cell.Neighbours.Left = getCell(g, x-1, y)
			cell.Neighbours.Right = getCell(g, x+1, y)
			cell.Neighbours.BottomLeft = getCell(g, x-1, y+1)
			cell.Neighbours.Bottom = getCell(g, x, y+1)
			cell.Neighbours.BottomRight = getCell(g, x+1, y+1)

			g[y][x] = cell
		}
	}

	return g
}

func getCell(g Grid, x, y int) *Cell {
	if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
		return nil
	}

	return &g[y][x]
}
