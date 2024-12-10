package main

import (
	"io"
	"log"
)

type (
	Topography []Row
	Row        []*Cell
	Direction  int
	Trailheads int
)

type Cell struct {
	Value      int
	Visited    bool
	Neighbours map[Direction]*Cell
}

const (
	Unset Direction = iota
	Up
	Right
	Down
	Left
)

const (
	Score Trailheads = iota
	Ratings
)

func main() {
	i := Task(load("input1.txt"), 1)
	log.Println("Part 1:", i)

	j := Task(load("input1.txt"), 2)
	log.Println("Part 2:", j)
}

func Task(r io.Reader, part int) int {
	topo, err := parse(r)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return topo.sum(Score)
	case 2:
		return topo.sum(Ratings)
	}

	return 0
}

func (t *Topography) sum(th Trailheads) int {
	var sum int

	for _, c := range t.start() {
		cs := []*Cell{c}

		for len(cs) != 0 {
			nc := c.next()
			cs = append(cs, nc...)

			c, cs = cs[len(cs)-1], cs[:len(cs)-1]

			if c.Value == 9 {
				sum++
			}

			if th == Score {
				c.Visited = true
			}
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
