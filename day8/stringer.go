package main

import (
	"fmt"
	"strings"
)

func (g Grid) String() string {
	var sb strings.Builder

	for _, row := range g {
		fmt.Fprintln(&sb, row)
	}

	return sb.String()
}

func (r Row) String() string {
	var sb strings.Builder

	for _, cell := range r {
		fmt.Fprint(&sb, cell)
	}

	return sb.String()
}

func (c Cell) String() string {
	if c.Antenna == nil {
		return "."
	}

	return string(*c.Antenna)
}
