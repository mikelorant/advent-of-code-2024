package main

import (
	"fmt"
	"strings"
)

func (ms Machines) String() string {
	var strs []string

	for _, m := range ms {
		strs = append(strs, fmt.Sprint(m))
	}

	return strings.Join(strs, "\n")
}

func (m Machine) String() string {
	var strs []string

	strs = append(strs, "Buttons:")

	for b, c := range m.Button {
		strs = append(strs, fmt.Sprintf("%v: %v, %v", string(b), c.X, c.Y))
	}

	strs = append(strs, fmt.Sprintf("Prize: %v, %v", m.Prize.X, m.Prize.Y))

	return strings.Join(strs, " ")
}
