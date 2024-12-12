package main

import (
	"log"
)

type (
	Garden    []Row
	Row       []*Plot
	Area      int
	Perimiter int
	Direction int
)

type Plot struct {
	Value      rune
	Visited    bool
	Neighbours map[Direction]*Plot
}

const (
	Up Direction = iota
	Right
	Down
	Left
)

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)
}

func Task(file string, part int) int {
	garden, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return price(garden)
	case 2:
		return 0
	}

	return 0
}

func price(g Garden) int {
	var sum int

	for _, row := range g {
		for _, plot := range row {
			if plot.Visited {
				continue
			}

			a, p := measure(plot)

			sum += int(a) * int(p)
		}
	}

	return sum
}

func measure(p *Plot) (Area, Perimiter) {
	var area Area
	var perimiter Perimiter

	plots := []*Plot{p}

	for len(plots) != 0 {
		p, plots = plots[len(plots)-1], plots[:len(plots)-1]

		if p.Visited {
			continue
		}

		area++
		perimiter += p.perimeter()

		p.Visited = true

		ps := p.connected()

		plots = append(plots, ps...)
	}

	return area, perimiter
}

func (p *Plot) connected() []*Plot {
	var ps []*Plot

	for _, dir := range []Direction{Up, Right, Down, Left} {
		np := p.Neighbours[dir]

		if np == nil {
			continue
		}

		if np.Visited {
			continue
		}

		if np.Value != p.Value {
			continue
		}

		ps = append(ps, np)
	}

	return ps
}

func (p *Plot) perimeter() Perimiter {
	var i Perimiter

	for _, dir := range []Direction{Up, Right, Down, Left} {
		if p.Neighbours[dir] == nil {
			i++

			continue
		}

		if p.Neighbours[dir].Value == p.Value {
			continue
		}

		i++
	}

	return i
}
