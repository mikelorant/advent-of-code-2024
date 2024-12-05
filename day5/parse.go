package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ruleSplitter   = "|"
	updateSplitter = ","
)

func parse(fn string) (Manual, error) {
	man := Manual{
		Rules: make(Rules, 0),
	}

	fh, err := os.Open(fn)
	if err != nil {
		return man, fmt.Errorf("unable to open file: %w", err)
	}

	splitter := ruleSplitter
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			splitter = updateSplitter

			continue
		}

		ints := toInts(strings.Split(line, splitter))

		switch {
		case splitter == ruleSplitter:
			page := ints[0]
			num := ints[1]

			man.Rules[page] = append(man.Rules[page], num)
		default:
			man.Updates = append(man.Updates, ints)
		}
	}
	if err := scanner.Err(); err != nil {
		return man, fmt.Errorf("unable to scan file: %w", err)
	}

	return man, nil
}

func toInts(ss []string) []int {
	var ints []int

	for _, s := range ss {
		ints = append(ints, mustInt(s))
	}

	return ints
}

func mustInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("unable to convert string to int: %v", err.Error())
	}

	return i
}
