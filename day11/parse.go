package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(r io.Reader) (Stones, error) {
	stones := make(Stones, 0)

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("unable to read all: %v", err.Error())
	}

	for _, i := range mustInts(b) {
		stones[i]++
	}

	return stones, nil
}

func load(file string) io.Reader {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal("unable to open file:", err.Error())
	}

	return fh
}

func mustInts(bs []byte) []int {
	var ints []int

	ss := strings.Split(string(bs), " ")

	for _, str := range ss {
		ints = append(ints, mustInt(strings.TrimSpace(str)))
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
