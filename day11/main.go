package main

import (
	"io"
	"log"
	"strconv"
)

func main() {
	i := Task(load("input1.txt"), 25, 1)
	log.Println("Part 1:", i)
}

func Task(r io.Reader, blinks int, part int) int {
	stones, err := parse(r)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return count(stones, blinks)
	case 2:
		return count(stones, blinks)
	}

	return 0
}

func count(stones []int, blinks int) int {
	for range blinks {
		stones = blink(stones)
	}

	return len(stones)
}

func blink(ss []int) []int {
	var ints []int

	for _, s := range ss {
		switch {
		case s == 0:
			ints = append(ints, 1)
		case len(strconv.Itoa(s))%2 == 0:
			s1, s2 := split(s)
			ints = append(ints, s1, s2)
		default:
			ints = append(ints, s*2024)
		}
	}

	return ints
}

func split(i int) (int, int) {
	str := strconv.Itoa(i)
	l := len(str) / 2

	return mustInt(str[:l]), mustInt(str[l:])
}
