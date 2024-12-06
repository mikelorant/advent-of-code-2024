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
		return steps(cell, Up)
	case 2:
		return 0
	}

	return 0
}

func guard(g Grid) *Cell {
	for _, row := range g {
		for _, cell := range row {
			if cell.Value == Guard {
				cell.Visited = true

				return cell
			}
		}
	}

	return nil
}

func steps(c *Cell, dir Direction) int {
	i := 1 // Include the guard's starting cell.

	for isCell(c, dir) {
		c, dir = step(c, dir)

		if !c.Visited {
			c.Visited = true

			i++
		}
	}

	return i
}

func step(c *Cell, d Direction) (*Cell, Direction) {
	if c.Neighbours[d].Value != Obstruction {
		return c.Neighbours[d], d
	}

	switch d {
	case Up:
		return c.Neighbours[Right], Right
	case Right:
		return c.Neighbours[Down], Down
	case Down:
		return c.Neighbours[Left], Left
	case Left:
		return c.Neighbours[Up], Up
	}

	return c, d
}

func isCell(c *Cell, d Direction) bool {
	switch d {
	case Up:
		return c.Neighbours[Up] != nil
	case Right:
		return c.Neighbours[Right] != nil
	case Down:
		return c.Neighbours[Down] != nil
	case Left:
		return c.Neighbours[Left] != nil
	}

	return false
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
