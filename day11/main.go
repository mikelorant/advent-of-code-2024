package main

import (
	"io"
	"log"
	"strconv"
)

type Stones map[int]int

func main() {
	i := Task(load("input1.txt"), 25, 1)
	log.Println("Part 1:", i)

	j := Task(load("input1.txt"), 75, 2)
	log.Println("Part 2:", j)
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

func count(s Stones, blinks int) int {
	for range blinks {
		s = blink(s)
	}

	var c int

	for _, i := range s {
		c += i
	}

	return c
}

func blink(s Stones) Stones {
	ss := make(Stones, 0)

	for num, count := range s {
		for _, i := range compute(num) {
			ss[i] += count
		}
	}

	return ss
}

func split(i int) []int {
	s := strconv.Itoa(i)
	l := len(s) / 2

	return []int{
		mustInt(s[:l]),
		mustInt(s[l:]),
	}
}

func compute(i int) []int {
	switch {
	case i == 0:
		return []int{1}
	case len(strconv.Itoa(i))%2 == 0:
		return split(i)
	default:
		return []int{i * 2024}
	}
}
