package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(fn string) (Garden, error) {
	var garden Garden

	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		garden = append(garden, toRow(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return addNeighbours(garden), nil
}

func toRow(str string) Row {
	var row Row

	for _, char := range str {
		row = append(row, &Plot{Value: char})
	}

	return row
}

func addNeighbours(g Garden) Garden {
	for y, row := range g {
		for x := range row {
			plot := g[y][x]

			plot.Neighbours = make(map[Direction]*Plot, 4)
			plot.Neighbours[Up] = getPlot(g, x, y-1)
			plot.Neighbours[Left] = getPlot(g, x-1, y)
			plot.Neighbours[Right] = getPlot(g, x+1, y)
			plot.Neighbours[Down] = getPlot(g, x, y+1)

			g[y][x] = plot
		}
	}

	return g
}

func getPlot(g Garden, x, y int) *Plot {
	if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
		return nil
	}

	return g[y][x]
}
