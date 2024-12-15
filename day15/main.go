package main

import (
	"log"
)

type Puzzle struct {
	Warehouse Warehouse
	Movements Movements
}

type (
	Movements []Direction
	Warehouse []Row
	Row       []*Cell
)

type Cell struct {
	Object     Object
	Neighbours map[Direction]*Cell
}

type Object int

const (
	Empty Object = iota
	Box
	Robot
	Wall
)

type Direction int

const (
	Unknown Direction = iota
	Down
	Left
	Right
	Up
)

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)
}

func Task(file string, part int) int {
	puz, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return sum(puz)
	case 2:
		return 0
	}

	return 0
}

func sum(p Puzzle) int {
	r := p.Warehouse.home()

	for _, d := range p.Movements {
		r, _ = r.move(d)
	}

	return p.Warehouse.calculate()
}

func (c *Cell) move(d Direction) (*Cell, bool) {
	nc := c.Neighbours[d]

	if nc.Object == Wall {
		return c, false
	}

	if nc.Object == Empty {
		return c.swap(d, true), true
	}

	ec := nc.findEmpty(d)
	if ec == nil {
		return c, false
	}

	d = d.inverse()
	for ec.Neighbours[d].Object != Robot {
		ec = ec.swap(d, true)
	}

	return ec.swap(d, false), true
}

func (c *Cell) swap(d Direction, move bool) *Cell {
	nc := c.Neighbours[d]

	c.Object, nc.Object = nc.Object, c.Object

	if move {
		return nc
	}

	return c
}

func (c *Cell) findEmpty(d Direction) *Cell {
	nc := c.Neighbours[d]

	for {
		if nc.Object == Empty {
			return nc
		}

		if nc.Object == Wall {
			return nil
		}

		nc = nc.Neighbours[d]
	}
}

func (d Direction) inverse() Direction {
	switch d {
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	case Up:
		return Down
	}

	return Unknown
}

func (w Warehouse) calculate() int {
	var sum int

	for y, row := range w {
		for x, cell := range row {
			if cell.Object == Box {
				sum += y*100 + x
			}
		}
	}

	return sum
}
