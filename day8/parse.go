package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func parse(fn string) (Grid, error) {
	var grid Grid

	fh, err := os.Open(fn)
	if err != nil {
		return grid, fmt.Errorf("unable to open file: %w", err)
	}

	scanner := bufio.NewScanner(fh)
	for i := 0; scanner.Scan(); i++ {
		row := toRow(scanner.Text(), i)

		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return grid, fmt.Errorf("unable to scan file: %w", err)
	}

	return grid, nil
}

func toRow(str string, i int) Row {
	var row Row

	for idx, char := range str {
		row = append(row, &Cell{
			X:       idx,
			Y:       i,
			Antenna: toAntenna(char),
		})
	}

	return row
}

func toAntenna(r rune) *Antenna {
	switch {
	case unicode.IsUpper(r):
		fallthrough
	case unicode.IsLower(r):
		fallthrough
	case unicode.IsDigit(r):
		ant := Antenna(r)

		return &ant
	}

	return nil
}
