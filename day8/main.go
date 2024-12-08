package main

import (
	"fmt"
	"log"
)

type (
	Grid    []Row
	Row     Cells
	Cells   []*Cell
	Antenna rune
)

type Cell struct {
	X         int
	Y         int
	Antenna   *Antenna
	Antinodes []Antenna
}

type Antinodes struct {
	Ax, Ay int
	Bx, By int
}

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)
}

func Task(file string, part int) int {
	grid, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	fmt.Println(grid)

	switch part {
	case 1:
		return locations(grid)
	case 2:
		return 0
	}

	return 0
}

func locations(g Grid) int {
	g.setAntinodes()

	return len(g.antinodes())
}

func (g Grid) antennas() map[Antenna]Cells {
	ants := make(map[Antenna]Cells, 0)

	for _, row := range g {
		for _, cell := range row {
			if cell.Antenna != nil {
				ants[*cell.Antenna] = append(ants[*cell.Antenna], cell)
			}
		}
	}

	return ants
}

func (g Grid) antinodes() Cells {
	var cells Cells

	for _, row := range g {
		for _, cell := range row {
			if len(cell.Antinodes) > 0 {
				cells = append(cells, cell)
			}
		}
	}

	return cells
}

func (g Grid) cell(x, y int) *Cell {
	if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
		return nil
	}

	return g[y][x]
}

func (g Grid) setAntinodes() {
	for ant, cells := range g.antennas() {
		for idx, cell := range cells {
			if idx >= len(cells) {
				break
			}

			for _, c := range cells[idx+1:] {
				antis := antinodes(cell, c)

				c1 := g.cell(antis.Ax, antis.Ay)
				c2 := g.cell(antis.Bx, antis.By)

				if c1 != nil {
					c1.Antinodes = append(c1.Antinodes, ant)
				}

				if c2 != nil {
					c2.Antinodes = append(c2.Antinodes, ant)
				}
			}
		}
	}
}

func antinodes(a, b *Cell) Antinodes {
	cx, dx := sequence(a.X, b.X)
	cy, dy := sequence(a.Y, b.Y)

	return Antinodes{
		Ax: cx, Ay: cy,
		Bx: dx, By: dy,
	}
}

func sequence(i, j int) (int, int) {
	if i == j {
		return i, j
	}

	s := abs(i - j)

	s1 := i - s
	s2 := i + s

	for s1 == i || s1 == j {
		s1 -= s
	}
	for s2 == i || s2 == j {
		s2 += s
	}

	if i > j {
		return s2, s1
	}

	return s1, s2
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
