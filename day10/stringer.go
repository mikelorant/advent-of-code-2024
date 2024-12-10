package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (t Topography) String() string {
	var ss []string

	for _, row := range t {
		ss = append(ss, fmt.Sprint(row))
	}

	return strings.Join(ss, "\n")
}

func (r Row) String() string {
	var str string

	for _, c := range r {
		str += fmt.Sprint(c)
	}

	return str
}

func (c Cell) String() string {
	switch c.Value {
	case -1:
		return "."
	default:
		return strconv.Itoa(c.Value)
	}
}
