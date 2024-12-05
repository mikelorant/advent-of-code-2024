package main

import (
	"log"
	"slices"
)

type Manual struct {
	Updates []Update
	Rules   Rules
}

type (
	Update []int
	Rules  map[int][]int
)

type State bool

const (
	Valid   State = true
	Invalid State = false
)

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)

	j := Task("input1.txt", 2)
	log.Println("Part 2:", j)
}

func Task(file string, part int) int {
	manual, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return sum(manual, Valid)
	case 2:
		return sum(manual, Invalid)
	}

	return 0
}

func sum(man Manual, s State) int {
	var sum int

	for _, up := range man.Updates {
		if isValid(up, man.Rules) && s == Valid {
			sum += middle(up)
		}

		if !isValid(up, man.Rules) && s == Invalid {
			sum += middle(reorder(up, man.Rules))
		}
	}

	return sum
}

func isValid(up Update, rs Rules) bool {
	for idx, page := range up {
		for _, val := range rs[page] {
			if !slices.Contains(up, val) {
				continue
			}

			if slices.Index(up, val) < idx {
				return false
			}
		}
	}

	return true
}

func reorder(up Update, rs Rules) Update {
START:
	for idx, page := range up {
		for _, num := range rs[page] {
			if !slices.Contains(up, num) {
				continue
			}

			if pos := slices.Index(up, num); pos < idx {
				move(up, idx, pos)

				goto START
			}
		}
	}

	return up
}

func move(up Update, i, j int) Update {
	val := up[i]

	_ = slices.Delete(up, i, i+1)
	_ = slices.Insert(up, j, val)

	return up
}

func middle(up Update) int {
	return up[len(up)/2]
}
