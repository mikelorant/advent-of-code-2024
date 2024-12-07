package main

import (
	"fmt"
	"log"
)

type (
	Grid []Row
	Row  []*Cell
)

type Cell struct {
	Value      Value
	Visited    bool
	Neighbours map[Direction]*Cell
}

type Value rune

const (
	Unset       Value = ' '
	Empty       Value = '.'
	Obstruction Value = '#'
	Guard       Value = '^'
)

type Direction int

const (
	Unknown Direction = iota
	Up
	Right
	Down
	Left
)

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)

	j := Task("input1.txt", 2)
	log.Println("Part 2:", j)
}

func Task(file string, part int) int {
	grid, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	cell := guard(grid)
	if cell == nil {
		log.Fatal("Unable to find guard cell:", err.Error())
	}

	switch part {
	case 1:
		return positions(cell, Up)
	case 2:
		return paradoxes(grid, cell, Up)
	}

	return 0
}

func guard(g Grid) *Cell {
	for _, row := range g {
		for _, cell := range row {
			if cell.Value == Guard {
				return cell
			}
		}
	}

	return nil
}

func positions(c *Cell, dir Direction) int {
	var i int

	for c != nil {
		if !c.Visited {
			c.Visited = true

			i++
		}

		c, dir = step(c, dir)
	}

	return i
}

func paradoxes(g Grid, c *Cell, dir Direction) int {
	var i int

	positions(c, Up)

	size := len(g) * len(g[0])

	for _, row := range g {
		for _, cell := range row {
			if !(cell.Visited && cell.Value == Empty) {
				continue
			}

			cell.Value = Obstruction

			if isLoop(c, dir, size) {
				i++
			}

			cell.Value = Empty
		}
	}

	return i
}

func isLoop(c *Cell, dir Direction, maxIter int) bool {
	var i int

	for c != nil {
		c, dir = step(c, dir)

		i++

		if i > maxIter {
			return true
		}
	}

	return false
}

func step(c *Cell, d Direction) (*Cell, Direction) {
	if c.Neighbours[d] == nil {
		return nil, d
	}

	for {
		if c.Neighbours[d].Value != Obstruction {
			break
		}

		d = changeDirection(d)
	}

	return c.Neighbours[d], d
}

func changeDirection(d Direction) Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}

	return d
}

func (v Value) String() string {
	switch v {
	case Unset:
		return " "
	case Empty:
		return "."
	case Obstruction:
		return "#"
	case Guard:
		return "^"
	}

	return " "
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Right:
		return "Right"
	case Down:
		return "Down"
	case Left:
		return "Left"
	}

	return ""
}

func (c Cell) String() string {
	return fmt.Sprintf("%v", c.Value)
}
