package main

import (
	"log"
	"slices"
)

type (
	Reports []Report
	Report  []int
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
	if slices.IsSorted(rep) && isAdjacent(rep) {
		return true
	}

	slices.Reverse(rep)

	if slices.IsSorted(rep) && isAdjacent(rep) {
		return true
	}

	return false
}

func isAdjacent(rep Report) bool {
	for idx, lvl := range rep[1:] {
		diff := lvl - rep[idx]
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func remove(rep Report, idx int) Report {
	return slices.Delete(slices.Clone(rep), idx, idx+1)
}
