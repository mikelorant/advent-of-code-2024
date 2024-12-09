package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func parse(r io.Reader) Disk {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("unable to read all: %v", err.Error())
	}

	return allocate(mustInts(b))
}

func allocate(ints []int) Disk {
	var disk Disk
	var id int

	alloc := true

	for _, i := range ints {
		var b Block

		switch alloc {
		case true:
			b.Type = Allocated
			b.Value = id

			id++
		default:
			b.Value = -1
		}

		alloc = !alloc

		bs := duplicate(b, i)
		disk = append(disk, bs...)
	}

	return disk
}

func load(file string) io.Reader {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal("unable to open file:", err.Error())
	}

	return fh
}

func duplicate(b Block, size int) []Block {
	var bs []Block

	for range size {
		bs = append(bs, b)
	}

	return bs
}

func mustInts(bs []byte) []int {
	var ints []int

	for _, b := range bs {
		str := string(b)

		if !unicode.IsDigit(rune(b)) {
			continue
		}

		ints = append(ints, mustInt(str))
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
