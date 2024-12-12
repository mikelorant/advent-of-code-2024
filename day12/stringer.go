package main

import (
	"fmt"
	"strings"
)

func (g Garden) String() string {
	var ss []string

	for _, row := range g {
		ss = append(ss, fmt.Sprint(row))
	}

	return strings.Join(ss, "\n")
}

func (r Row) String() string {
	var str string

	for _, plot := range r {
		str += fmt.Sprint(plot)
	}

	return str
}

func (p Plot) String() string {
	return string(p.Value)
}
