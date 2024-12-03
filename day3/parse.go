package main

import (
	"io"
	"log"
	"regexp"
	"strconv"
)

const (
	validExp    = `(mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\))`
	multiplyExp = `mul\((\d+),(\d+)\)`
)

const (
	doStr   = "do()"
	dontStr = "don't()"
)

func parse(r io.Reader) Commands {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("unable to read all: %v", err.Error())
	}

	return convert(string(b))
}

func convert(str string) Commands {
	var cmds Commands

	re := regexp.MustCompile(validExp)

	for _, val := range re.FindAllString(str, -1) {
		switch val {
		case doStr:
			cmds = append(cmds, -1)
		case dontStr:
			cmds = append(cmds, -2)
		default:
			cmds = append(cmds, multiply(val))
		}
	}

	return cmds
}

func multiply(str string) int {
	re := regexp.MustCompile(multiplyExp)
	matches := re.FindStringSubmatch(str)

	return mustInt(matches[1]) * mustInt(matches[2])
}

func mustInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("unable to convert string to int: %v", err.Error())
	}

	return i
}
