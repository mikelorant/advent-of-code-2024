package main

import (
	"fmt"
	"strings"
)

func (w Warehouse) String() string {
	var strs []string

	for _, row := range w {
		strs = append(strs, fmt.Sprint(row))
	}

	return strings.Join(strs, "\n")
}

func (r Row) String() string {
	var str string

	for _, cell := range r {
		str += fmt.Sprint(cell)
	}

	return str
}

func (c Cell) String() string {
	return fmt.Sprint(c.Object)
}

func (d Direction) String() string {
	switch d {
	case Down:
		return "v"
	case Left:
		return "<"
	case Right:
		return ">"
	case Up:
		return "^"
	}

	return "?"
}

func (o Object) String() string {
	switch o {
	case Empty:
		return "."
	case Box:
		return "O"
	case Robot:
		return "@"
	case Wall:
		return "#"
	}

	return "?"
}
