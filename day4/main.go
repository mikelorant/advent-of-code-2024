package main

import (
	"fmt"
	"log"
	"strings"
)

type (
	Grid []Row
	Row  []Cell
)

type Cell struct {
	Value      string
	Neighbours Neighbours
}

type Neighbours struct {
	TopLeft     *Cell
	Top         *Cell
	TopRight    *Cell
	Left        *Cell
	Right       *Cell
	BottomLeft  *Cell
	Bottom      *Cell
	BottomRight *Cell
}

type Direction int

type Adjacent struct {
	Cell      *Cell
	Direction Direction
}

var (
	xmas = []string{"X", "M", "A", "S"}
	mas  = []string{"M", "A", "S"}
)

const (
	Undefined Direction = iota
	TopLeft
	Top
	TopRight
	Left
	Right
	BottomLeft
	Bottom
	BottomRight
)

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)

	j := Task("input1.txt", 2)
	log.Println("Part 2:", j)
}

func Task(file string, part int) int {
	filter := xmas
	if part == 2 {
		filter = mas
	}

	grid, err := parse(file, filter)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return searchXMAS(grid)
	case 2:
		return searchMAS(grid)
	}

	return 0
}

func searchXMAS(g Grid) int {
	var count int

	for _, x := range candidates(g, "X") {
		ms := adjacents(x, "M")

		for _, m := range ms {
			a := lookupAdjacent(*m.Cell, "A", m.Direction)
			if a == nil {
				continue
			}

			s := lookupAdjacent(*a, "S", m.Direction)
			if s == nil {
				continue
			}

			count++
		}
	}

	return count
}

func searchMAS(g Grid) int {
	var count int

	for _, a := range candidates(g, "A") {
		if checkDiagonals(a) {
			count++
		}
	}

	return count
}

func candidates(g Grid, str string) []Cell {
	var res []Cell

	for _, row := range g {
		for _, cell := range row {
			if cell.Value == str {
				res = append(res, cell)
			}
		}
	}

	return res
}

func checkDiagonals(cell Cell) bool {
	var count int

	diagonals := []Direction{TopLeft, TopRight, BottomLeft, BottomRight}

	for _, dir := range diagonals {
		if lookupAdjacent(cell, "M", dir) == nil {
			continue
		}

		if lookupAdjacent(cell, "S", opposite(dir)) == nil {
			continue
		}

		count++
	}

	return count == 2
}

func adjacents(cell Cell, str string) []Adjacent {
	var adjs []Adjacent

	around := []Direction{TopLeft, Top, TopRight, Left, Right, BottomLeft, Bottom, BottomRight}

	for _, dir := range around {
		adjCell := lookupAdjacent(cell, str, dir)
		if adjCell == nil {
			continue
		}

		adj := Adjacent{
			Cell:      adjCell,
			Direction: dir,
		}

		adjs = append(adjs, adj)
	}

	return adjs
}

func lookupAdjacent(cell Cell, str string, dir Direction) *Cell {
	adj := adjacent(cell, dir)
	if adj == nil || adj.Value != str {
		return nil
	}

	return adj
}

func adjacent(cell Cell, d Direction) *Cell {
	switch d {
	case TopLeft:
		return cell.Neighbours.TopLeft
	case Top:
		return cell.Neighbours.Top
	case TopRight:
		return cell.Neighbours.TopRight
	case Left:
		return cell.Neighbours.Left
	case Right:
		return cell.Neighbours.Right
	case BottomLeft:
		return cell.Neighbours.BottomLeft
	case Bottom:
		return cell.Neighbours.Bottom
	case BottomRight:
		return cell.Neighbours.BottomRight
	}

	return &cell
}

func opposite(d Direction) Direction {
	switch d {
	case TopLeft:
		return BottomRight
	case TopRight:
		return BottomLeft
	case BottomLeft:
		return TopRight
	case BottomRight:
		return TopLeft
	}

	return d
}

func (g Grid) String() string {
	var sb strings.Builder

	for _, row := range g {
		for _, cell := range row {
			fmt.Fprintf(&sb, "%v", cell.Value)
		}

		fmt.Fprintln(&sb, "")
	}

	return sb.String()
}
