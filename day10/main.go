package main

import (
	"io"
	"log"
)

type (
	Topography []Row
	Row        []*Cell
)

type Cell struct {
	Value      int
	Visited    bool
	Neighbours map[Direction]*Cell
}

type Direction int

const (
	Unset Direction = iota
	Up
	Right
	Down
	Left
)

func main() {
	i := Task(load("input1.txt"), 1)
	log.Println("Part 1:", i)
}

func Task(r io.Reader, part int) int {
	topo, err := parse(r)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return topo.sum()
	case 2:
		return 0
	}

	return 0
}

func (t *Topography) sum() int {
	var sum int
	var cs []*Cell

	for _, c := range t.start() {
		for {
			nc := c.next()
			cs = append(cs, nc...)

			if len(cs) == 0 {
				break
			}

			c, cs = cs[len(cs)-1], cs[:len(cs)-1]

			if c.Value == 9 {
				sum++
			}

			c.Visited = true
		}

		t.reset()
	}

	return sum
}

func (t *Topography) start() []*Cell {
	var cells []*Cell

	for _, row := range *t {
		for idx, cell := range row {
			if cell.Value != 0 {
				continue
			}

			cells = append(cells, row[idx])
		}
	}

	return cells
}

func (c *Cell) next() []*Cell {
	var cells []*Cell

	for _, dir := range []Direction{Up, Right, Down, Left} {
		if c.Neighbours[dir] == nil {
			continue
		}

		nc := c.Neighbours[dir]

		if nc.Visited {
			continue
		}

		if c.Value != nc.Value-1 {
			continue
		}

		cells = append(cells, nc)
	}

	return cells
}

func (t *Topography) reset() {
	for _, row := range *t {
		for _, cell := range row {
			cell.Visited = false
		}
	}
}
