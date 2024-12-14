package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parse(fn string) (Machines, error) {
	var ms Machines

	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	reButton := regexp.MustCompile(`^Button [A-Z]`)
	rePrize := regexp.MustCompile(`^Prize`)

	scanner := bufio.NewScanner(fh)

	var m Machine
	m.Button = make(map[rune]Coordinate, 0)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case reButton.MatchString(line):
			button, coord := toButton(line)
			m.Button[button] = coord
		case rePrize.MatchString(line):
			m.Prize = toPrize(line)
		case line == "":
			ms = append(ms, m)
			m = Machine{}
			m.Button = make(map[rune]Coordinate, 0)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	ms = append(ms, m)

	return ms, nil
}

func toButton(str string) (rune, Coordinate) {
	re := regexp.MustCompile(`Button ([A-Z]): X\+(\d+), Y\+(\d+)`)
	ss := re.FindStringSubmatch(str)

	c := Coordinate{
		X: mustInt(ss[2]),
		Y: mustInt(ss[3]),
	}

	b := []rune(ss[1])[0]

	return b, c
}

func toPrize(str string) Coordinate {
	re := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	ss := re.FindStringSubmatch(str)

	return Coordinate{
		X: mustInt(ss[1]),
		Y: mustInt(ss[2]),
	}
}

func mustInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("unable to convert string to int: %v", err.Error())
	}

	return i
}
