package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (rs Robots) String() string {
	var strs []string

	for _, r := range rs {
		strs = append(strs, fmt.Sprint(r))
	}

	return strings.Join(strs, "\n")
}

func (r Robot) String() string {
	return fmt.Sprintf("Loc: %v, %v Vel: %v, %v", r.Location.X, r.Location.Y, r.Velocity.X, r.Velocity.Y)
}

func (g Grid) String() string {
	var sb strings.Builder

	for _, row := range g {
		for _, cell := range row {
			if cell.Quadrant == 0 {
				fmt.Fprintf(&sb, " ")

				continue
			}

			c := len(cell.Robots)

			switch c {
			case 0:
				fmt.Fprint(&sb, ".")
			default:
				fmt.Fprint(&sb, strconv.Itoa(c))
			}
		}

		fmt.Fprintf(&sb, "\n")
	}

	return sb.String()
}
