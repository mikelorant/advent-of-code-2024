package main

import (
	"log"
	"slices"
)

type (
	Reports   []Report
	Report    []int
	Direction int
)

const (
	Unknown Direction = iota
	Descending
	Ascending
	Flat
)

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)

	j := Task("input1.txt", 2)
	log.Println("Part 2:", j)
}

func Task(file string, part int) int {
	reps, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return check(reps, false)
	case 2:
		return check(reps, true)
	}

	return 0
}

func check(reps Reports, damp bool) int {
	var cnt int

	for _, rep := range reps {
		if isSafe(rep) {
			cnt++

			continue
		}

		if !damp {
			continue
		}

		for idx := range rep {
			if isSafe(remove(rep, idx)) {
				cnt++

				break
			}
		}
	}

	return cnt
}

func isSafe(rep Report) bool {
	for idx, lvl := range rep {
		if idx == 0 {
			continue
		}

		prevLvl := rep[idx-1]

		diff := lvl - prevLvl
		if direction(rep) == Descending {
			diff = -diff
		}

		if diff > 0 && diff <= 3 {
			continue
		}

		return false
	}

	return true
}

func direction(rep Report) Direction {
	for idx, lvl := range rep {
		if idx == 0 {
			continue
		}

		prevLvl := rep[idx-1]

		switch {
		case lvl < prevLvl:
			return Descending
		case lvl > prevLvl:
			return Ascending
		}
	}

	return Flat
}

func remove(rep Report, idx int) Report {
	r := make(Report, len(rep))
	copy(r, rep)

	return slices.Delete(r, idx, idx+1)
}
