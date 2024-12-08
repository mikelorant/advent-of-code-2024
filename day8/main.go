package main

import (
	"fmt"
	"log"
)

type (
	Grid        []Row
	Row         Cells
	Antenna     rune
	Cells       []*Cell
	Coordinates []Coordinate
	Harmonics   int
)

type Cell struct {
	X         int
	Y         int
	Antenna   *Antenna
	Antinodes []Antenna
}

type Coordinate struct {
	X int
	Y int
}

const (
	NoHarmonics Harmonics = iota
	WithHarmonics
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

	fmt.Println(grid)

	switch part {
	case 1:
		return locations(grid, NoHarmonics)
	case 2:
		return locations(grid, WithHarmonics)
	}

	return 0
}

func locations(g Grid, h Harmonics) int {
	g.setAntinodes(h)

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

func (g Grid) setAntinodes(h Harmonics) {
	for ant, cells := range g.antennas() {
		for idx, cell := range cells {
			if idx >= len(cells) {
				break
			}

			for _, c := range cells[idx+1:] {
				bound := Coordinate{X: len(g[0]), Y: len(g)}
				for _, a := range antinodes(cell, c, bound, h) {
					g.addAntinode(a, ant)
				}
			}
		}
	}
}

func (g Grid) addAntinode(coord Coordinate, ant Antenna) bool {
	cell := g.cell(coord.X, coord.Y)

	if cell == nil {
		return false
	}

	cell.Antinodes = append(cell.Antinodes, ant)

	return true
}

func antinodes(a, b *Cell, bound Coordinate, h Harmonics) Coordinates {
	cA := toCoord(a)
	cB := toCoord(b)

	if h == WithHarmonics {
		return extendAllInclusive(cA, cB, bound)
	}

	return extendBoth(cA, cB)
}
