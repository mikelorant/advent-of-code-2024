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
		return sum(cmds)
	case 2:
		return sum(filter(cmds))
	}

	return 0
}

func sum(cmds Commands) int {
	var val int

	for _, i := range cmds {
		if i > 0 {
			val += i
		}
	}

	return val
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
