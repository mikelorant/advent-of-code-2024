package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(fn string) (Puzzle, error) {
	var puz Puzzle

	fh, err := os.Open(fn)
	if err != nil {
		return puz, fmt.Errorf("unable to open file: %w", err)
	}

	var wh Warehouse

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		wh = append(wh, toRow(line))
	}
	if err := scanner.Err(); err != nil {
		return puz, fmt.Errorf("unable to scan file: %w", err)
	}

	puz.Warehouse = wh.addNeighbours()

	for scanner.Scan() {
		movs := toMovements(scanner.Text())

		puz.Movements = append(puz.Movements, movs...)
	}
	if err := scanner.Err(); err != nil {
		return puz, fmt.Errorf("unable to scan file: %w", err)
	}

	return puz, nil
}

func toRow(str string) Row {
	var row Row

	for _, char := range str {
		row = append(row, &Cell{
			Object: toObject(char),
		})
	}

	return row
}

func toObject(r rune) Object {
	switch r {
	case 'O':
		return Box
	case '@':
		return Robot
	case '#':
		return Wall
	}

	return Empty
}

func toMovements(str string) Movements {
	var ms Movements

	for _, char := range str {
		ms = append(ms, toDirection(char))
	}

	return ms
}

func toDirection(r rune) Direction {
	switch r {
	case 'v':
		return Down
	case '<':
		return Left
	case '>':
		return Right
	case '^':
		return Up
	}

	return Unknown
}

func (w Warehouse) addNeighbours() Warehouse {
	for y, row := range w {
		for x := range row {
			cell := w[y][x]

			cell.Neighbours = make(map[Direction]*Cell, 4)
			cell.Neighbours[Up] = w.getCell(x, y-1)
			cell.Neighbours[Left] = w.getCell(x-1, y)
			cell.Neighbours[Right] = w.getCell(x+1, y)
			cell.Neighbours[Down] = w.getCell(x, y+1)

			w[y][x] = cell
		}
	}

	return w
}

func (w Warehouse) getCell(x, y int) *Cell {
	if y < 0 || y >= len(w) || x < 0 || x >= len(w[y]) {
		return nil
	}

	return w[y][x]
}

func (w Warehouse) home() *Cell {
	for _, row := range w {
		for _, cell := range row {
			if cell.Object == Robot {
				return cell
			}
		}
	}

	return nil
}
