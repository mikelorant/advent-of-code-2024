package main

import (
	"io"
	"log"
	"os"
)

type Commands []int

func main() {
	i := Task(load("input1.txt"), 1)
	log.Println("Part 1:", i)

	j := Task(load("input1.txt"), 2)
	log.Println("Part 2:", j)
}

func Task(r io.Reader, part int) int {
	cmds := parse(r)

	switch part {
	case 1:
		return sum(cmds, false)
	case 2:
		return sum(cmds, true)
	}

	return 0
}

func sum(cmds Commands, cond bool) int {
	if cond {
		cmds = filter(cmds)
	}

	return scan(cmds)
}

func scan(cmds []int) int {
	var sum int

	for _, i := range cmds {
		if i > 0 {
			sum += i
		}
	}

	return sum
}

func filter(cmds Commands) Commands {
	var filtCmds Commands

	enabled := true

	for _, cmd := range cmds {
		switch {
		case cmd == -1:
			enabled = true
		case cmd == -2:
			enabled = false
		case enabled:
			filtCmds = append(filtCmds, cmd)
		}
	}

	return filtCmds
}

func load(file string) io.Reader {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal("unable to open file:", err.Error())
	}

	return fh
}