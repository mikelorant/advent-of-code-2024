package main

import (
	"fmt"
	"log"
)

type (
	Grid   []Row
	Row    []*Cell
	Robots []*Robot
)

type Cell struct {
	Location Coordinate
	Quadrant int
	Robots   []*Robot
}

type Robot struct {
	Location Coordinate
	Velocity Coordinate
	Updated  bool
}

type Coordinate struct {
	X int
	Y int
}

func main() {
	i := Task("input1.txt", 1, 101, 103)
	log.Println("Part 1:", i)

	j := Task("input1.txt", 2, 101, 103)
	log.Println("Part 2:", j)
}

func Task(file string, part int, x, y int) int {
	rs, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return safety(rs, x, y)
	case 2:
		return anomaly(rs, x, y)
	}

	return 0
}

func safety(rs Robots, x, y int) int {
	loops := 100

	for range loops {
		rs.update(x, y)
	}

	return rs.factor(x, y)
}

func anomaly(rs Robots, x, y int) int {
	var i int

	for {
		rs.update(x, y)

		i++

		grid := makeGrid(x, y)
		grid.add(rs)

		if grid.isAnomaly() {
			fmt.Println(grid)

			break
		}
	}

	return i
}

func (rs Robots) factor(x, y int) int {
	var factor int

	for q := 1; q < 5; q++ {
		num := rs.count(x, y, q)

		if factor == 0 {
			factor = num

			continue
		}

		factor *= num
	}

	return factor
}

func (rs Robots) count(maxX, maxY int, quad int) int {
	var num int

	for _, r := range rs {
		if quadrant(maxX, maxY, r.Location.X, r.Location.Y) != quad {
			continue
		}

		num++
	}

	return num
}

func (rs Robots) update(maxX, maxY int) {
	for _, r := range rs {
		r.move(maxX, maxY)
	}
}

func (r *Robot) move(maxX, maxY int) {
	velX, velY := r.Velocity.X, r.Velocity.Y

	newX := (r.Location.X + velX + maxX) % maxX
	newY := (r.Location.Y + velY + maxY) % maxY

	r.Location.X, r.Location.Y = newX, newY
}

func (g Grid) add(rs Robots) {
	for _, r := range rs {
		x, y := r.Location.X, r.Location.Y

		g[y][x].Robots = append(g[y][x].Robots, r)
	}
}

func (g Grid) isAnomaly() bool {
	var seq int

	const anomaly = 6

	for _, row := range g {
		for _, cell := range row {
			if len(cell.Robots) == 0 {
				seq = 0

				continue
			}

			seq++

			if seq > anomaly {
				return true
			}
		}
	}

	return false
}

func makeGrid(x, y int) Grid {
	var g Grid

	for row := range y {
		var r Row

		for col := range x {
			c := Cell{
				Location: Coordinate{X: col, Y: row},
				Quadrant: quadrant(x, y, col, row),
			}

			r = append(r, &c)
		}

		g = append(g, r)
	}

	return g
}

func quadrant(maxX, maxY, x, y int) int {
	midX := maxX / 2
	midY := maxY / 2

	switch {
	case y < midY && x < midX:
		return 1
	case y < midY && x > midX:
		return 2
	case y > midY && x < midX:
		return 3
	case y > midY && x > midX:
		return 4
	}

	return 0
}
