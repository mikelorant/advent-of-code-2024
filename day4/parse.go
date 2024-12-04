package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(fn string) (Grid, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	var grid Grid

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		row := toRow(scanner.Text())

		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return grid, nil
}

func toRow(str string) Row {
	var row Row

	for _, char := range str {
		switch char {
		case 'X', 'M', 'A', 'S':
			row = append(row, string(char))
		default:
			row = append(row, ".")
		}
	}

	return row
}
