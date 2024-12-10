package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func parse(r io.Reader) (Topography, error) {
	var topo Topography

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := toRow(scanner.Text())

		topo = append(topo, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return addNeighbours(topo), nil
}

func load(file string) io.Reader {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal("unable to open file:", err.Error())
	}

	return fh
}

func toRow(str string) Row {
	var row Row

	for _, char := range str {
		switch {
		case unicode.IsDigit(char):
			row = append(row, &Cell{
				Value: mustInt(string(char)),
			})
		case char == '.':
			row = append(row, &Cell{
				Value: -1,
			})
		default:
		}
	}

	return row
}

func addNeighbours(t Topography) Topography {
	for y, row := range t {
		for x := range row {
			cell := t[y][x]

			cell.Neighbours = make(map[Direction]*Cell, 4)
			cell.Neighbours[Up] = getCell(t, x, y-1)
			cell.Neighbours[Left] = getCell(t, x-1, y)
			cell.Neighbours[Right] = getCell(t, x+1, y)
			cell.Neighbours[Down] = getCell(t, x, y+1)

			t[y][x] = cell
		}
	}

	return t
}

func mustInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("unable to convert string to int: %v", err.Error())
	}

	return i
}

func getCell(t Topography, x, y int) *Cell {
	if y < 0 || y >= len(t) || x < 0 || x >= len(t[y]) {
		return nil
	}

	return t[y][x]
}
