package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)

	j := Task("input1.txt", 2)
	log.Println("Part 2:", j)
}

type (
	Grid []Row
	Row  []string
)

func Task(file string, part int) int {
	grid, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return search(grid)
	case 2:
		return 0
	}

	return 0
}

func (g Grid) String() string {
	var sb strings.Builder

	for _, row := range g {
		fmt.Fprintln(&sb, strings.Join(row, ""))
	}

	return sb.String()
}

func search(g Grid) int {
	x := find(g, "X")
	s := find(g, "S")

	var paths [][][]int

	for _, i := range x {
		for _, j := range s {
			paths = append(paths, path(i, j))
		}
	}

	var res [][][]int
	for _, p := range paths {
		if isValid(g, p) {
			res = append(res, p)
		}
	}

	return len(res)
}

func path(a, b []int) [][]int {
	var res [][]int

	ranges := connect(a, b)

	for idx := range steps(ranges[0], ranges[1]) {
		res = append(res, []int{
			next(ranges[0], idx),
			next(ranges[1], idx),
		})
	}

	return res
}

func find(g Grid, str string) [][]int {
	var loc [][]int

	for y, row := range g {
		for x, char := range row {
			if char == str {
				loc = append(loc, []int{y, x})
			}
		}
	}

	return loc
}

func connect(i, j []int) [][]int {
	var res [][]int

	res = append(res, between(i[0], j[0]))
	res = append(res, between(i[1], j[1]))

	return res
}

func between(a, b int) []int {
	var res []int

	for i := a; i != b; i = nextInt(i, b) {
		res = append(res, i)
	}
	res = append(res, b)

	return res
}

func nextInt(i, j int) int {
	switch {
	case i < j:
		return i + 1
	case i > j:
		return i - 1
	default:
		return i
	}
}

func steps(i, j []int) int {
	if len(i) > len(j) {
		return len(i)
	}

	return len(j)
}

func next(nums []int, idx int) int {
	if idx >= len(nums) {
		return nums[len(nums)-1]
	}

	return nums[idx]
}

func isValid(g Grid, points [][]int) bool {
	// Correct length
	if len(points) != 4 {
		return false
	}

	// Within bounds
	for _, point := range points {
		if point[0] < 0 || point[1] < 0 {
			return false
		}

		if point[0] > len(g) || point[0] > len(g[0]) {
			return false
		}
	}

	// Correct letters
	for idx := range 4 {
		y := points[idx][0]
		x := points[idx][1]

		switch idx {
		case 0:
			if g[y][x] != "X" {
				return false
			}
		case 1:
			if g[y][x] != "M" {
				return false
			}
		case 2:
			if g[y][x] != "A" {
				return false
			}
		case 3:
			if g[y][x] != "S" {
				return false
			}
		}
	}

	// Correct direction
	xy := points[0][0]
	xx := points[0][1]

	sy := points[3][0]
	sx := points[3][1]

	if !(abs(xy-sy) == 0 || abs(xy-sy) == 3) {
		return false
	}

	if !(abs(xx-sx) == 0 || abs(xx-sx) == 3) {
		return false
	}

	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
