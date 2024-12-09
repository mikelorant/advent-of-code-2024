package main

import (
	"fmt"
	"strconv"
)

func (d Disk) String() string {
	var str string

	for _, b := range d {
		str += fmt.Sprint(b)
	}

	return str
}

func (b Block) String() string {
	if b.Type == Allocated {
		return strconv.Itoa(b.Value)
	}

	return "."
}
