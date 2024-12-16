package main

import (
	"fmt"
	"image"
	"log"
)

type Maze map[image.Point]*Tile

type Tile struct {
	Type Type
}

type Type int

const (
	Unknown Type = iota
	Empty
	End
	Start
	Wall
)

func main() {
	do := Task("input1.txt", 1)
	log.Println("Part 1:", do)
}

func Task(file string, part int) int {
	maze, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	fmt.Println(maze)

	switch part {
	case 1:
		return score(maze)
	case 2:
		return 0
	}

	return 0
}

func score(_ Maze) int {
	return 0
}
