package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(fn string) (Grid, error) {
	var grid Grid

	fh, err := os.Open(fn)
	if err != nil {
		return grid, fmt.Errorf("unable to open file: %w", err)
	}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		grid = append(grid, toRow(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return grid, fmt.Errorf("unable to scan file: %w", err)
	}

	return addNeighbours(grid), nil
}

func toRow(str string) Row {
	var row Row

	for _, char := range str {
		row = append(row, &Cell{
			Value: toValue(char),
		})
	}

	return row
}

func addNeighbours(g Grid) Grid {
	for y, row := range g {
		for x := range row {
			cell := g[y][x]

			cell.Neighbours = make(map[Direction]*Cell, 4)
			cell.Neighbours[Up] = getCell(g, x, y-1)
			cell.Neighbours[Left] = getCell(g, x-1, y)
			cell.Neighbours[Right] = getCell(g, x+1, y)
			cell.Neighbours[Down] = getCell(g, x, y+1)

			g[y][x] = cell
		}
	}

	return g
}

func getCell(g Grid, x, y int) *Cell {
	if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
		return nil
	}

	return g[y][x]
}

func toValue(r rune) Value {
	switch r {
	case '.':
		return Empty
	case '#':
		return Obstruction
	case '^':
		return Guard
	}

	return Unset
}
