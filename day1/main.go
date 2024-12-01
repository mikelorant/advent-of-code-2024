package main

import (
	"log"
	"sort"
)

type List []int

func main() {
	sum := Task("input1.txt", 1)
	log.Println("Part 1:", sum)

	score := Task("input1.txt", 2)
	log.Println("Part 2:", score)
}

func Task(file string, part int) int {
	lists, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return sumDiff(lists)
	case 2:
		return simScore(lists)
	}

	return 0
}

func sumDiff(lists []List) int {
	var sum int

	for idx, l := range lists {
		lists[idx] = sortList(l)
	}

	for idx := range len(lists[0]) {
		sum += diff(lists[0][idx], lists[1][idx])
	}

	return sum
}

func simScore(lists []List) int {
	var sum int

	for _, v := range lists[0] {
		sum += v * count(v, lists[1])
	}

	return sum
}

func sortList(l List) List {
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})

	return l
}

func diff(i, j int) int {
	d := i - j
	if d < 0 {
		return -d
	}

	return d
}

func count(val int, l List) int {
	var i int

	for _, v := range l {
		if v == val {
			i += 1
		}
	}

	return i
}
