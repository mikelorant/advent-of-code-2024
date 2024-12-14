package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parse(fn string) (Robots, error) {
	var robots Robots

	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		robots = append(robots, toRobot(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return robots, fmt.Errorf("unable to scan file: %w", err)
	}

	return robots, nil
}

func toRobot(str string) *Robot {
	re := regexp.MustCompile(`p=([+-]?\d+),([+-]?\d+) v=([+-]?\d+),([+-]?\d+)`)
	ss := re.FindStringSubmatch(str)

	return &Robot{
		Location: Coordinate{X: mustInt(ss[1]), Y: mustInt(ss[2])},
		Velocity: Coordinate{X: mustInt(ss[3]), Y: mustInt(ss[4])},
	}
}

func mustInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("unable to convert string to int: %v", err.Error())
	}

	return i
}
