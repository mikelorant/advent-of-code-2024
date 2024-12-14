package main

import (
	"log"
)

type Machines []Machine

type Machine struct {
	Button map[rune]Coordinate
	Prize  Coordinate
}

type Coordinate struct {
	X, Y int
}

type Pairs [2]int

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)
}

func Task(file string, part int) int {
	machines, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return sum(machines, true)
	case 2:
		for idx := range machines {
			machines[idx].Prize.X += 10000000000000
			machines[idx].Prize.Y += 10000000000000
		}

		return sum(machines, false)
	}

	return 0
}

func sum(ms Machines, limit bool) int {
	var sum int

	for _, m := range ms {
		sum += presses(m, limit)
	}

	return sum
}

func presses(m Machine, limit bool) int {
	var minp int

	ax, bx := m.Button['A'].X, m.Button['B'].X

	for _, ps := range combinations(ax, bx, m.Prize.X) {
		ap, bp := ps[0], ps[1]
		ay, by := m.Button['A'].Y, m.Button['B'].Y

		if !isValid(ay, ap, by, bp, m.Prize.Y) {
			continue
		}

		if limit {
			if ap > 100 || bp > 100 {
				continue
			}
		}

		t := tokens(ap, bp)
		if t < minp || minp == 0 {
			minp = t
		}
	}

	return minp
}

func tokens(a, b int) int {
	return a*3 + b
}

func isValid(a, x, b, y, c int) bool {
	return a*x+b*y == c
}

func combinations(a, b, c int) []Pairs {
	var ints []Pairs

	g := gcd(a, b)

	if c%g != 0 {
		return nil
	}

	i := c / a
	if i < 0 {
		return nil
	}

	for x := range i + 1 {
		if (c - a*x) < 0 {
			continue
		}

		if (c-a*x)%b != 0 {
			continue
		}

		y := (c - a*x) / b

		ints = append(ints, Pairs{x, y})
	}

	return ints
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
