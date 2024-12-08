package main

func extendBoth(a, b Coordinate) Coordinates {
	var coords Coordinates

	coords = append(coords, next(a, b))
	coords = append(coords, next(b, a))

	return coords
}

func extendAllInclusive(a, b, bound Coordinate) Coordinates {
	var coords Coordinates

	fwd := extend(a, b, bound)
	rev := extend(b, a, bound)

	coords = append(coords, a, b)
	coords = append(coords, fwd...)
	coords = append(coords, rev...)

	return coords
}

func extend(a, b, bound Coordinate) Coordinates {
	var coords []Coordinate

	for {
		c := next(a, b)

		if c.X < 0 || c.X > bound.X || c.Y < 0 || c.Y > bound.Y {
			break
		}

		coords = append(coords, c)

		a, b = b, c
	}

	return coords
}

func next(a, b Coordinate) Coordinate {
	return Coordinate{
		sequence(a.X, b.X),
		sequence(a.Y, b.Y),
	}
}

func sequence(i, j int) int {
	if i == j {
		return i
	}

	if j > i {
		return j + (j - i)
	}

	return j - (i - j)
}

func toCoord(c *Cell) Coordinate {
	return Coordinate{
		X: c.X,
		Y: c.Y,
	}
}
