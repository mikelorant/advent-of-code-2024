package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func parse(fn string) (Maze, error) {
	maze := make(Maze, 0)

	fh, err := os.Open(fn)
	if err != nil {
		return maze, fmt.Errorf("unable to open file: %w", err)
	}

	scanner := bufio.NewScanner(fh)
	for i := 0; scanner.Scan(); i++ {
		for pt, t := range toMaze(scanner.Text(), i) {
			maze[pt] = t
		}
	}
	if err := scanner.Err(); err != nil {
		return maze, fmt.Errorf("unable to scan file: %w", err)
	}

	return maze, nil
}

func toMaze(str string, y int) Maze {
	res := make(Maze, 0)

	for x, r := range str {
		pt := image.Point{X: x, Y: y}
		res[pt] = &Tile{
			Type: toType(r),
		}
	}

	return res
}

func toType(r rune) Type {
	switch r {
	case '.':
		return Empty
	case 'E':
		return End
	case 'S':
		return Start
	case '#':
		return Wall
	default:
		return Unknown
	}
}
