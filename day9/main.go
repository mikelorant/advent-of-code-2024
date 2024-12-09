package main

import (
	"io"
	"log"
	"slices"
)

type Disk []Block

type Block struct {
	Type  Type
	Value int
}

type Type int

const (
	Empty Type = iota
	Allocated
)

func main() {
	i := Task(load("input1.txt"), 1)
	log.Println("Part 1:", i)
}

func Task(r io.Reader, part int) int {
	disk := parse(r)

	switch part {
	case 1:
		disk.defragment()

		return disk.checksum()
	case 2:
		return 0
	}

	return 0
}

func (d Disk) defragment() {
	empty := d.empty()
	alloc := d.allocated()

	slices.Reverse(alloc)

	for !d.isDone() {
		for idx := range empty {
			if idx >= len(alloc) {
				break
			}

			if d.isDone() {
				break
			}

			d.swap(empty[idx], alloc[idx])
		}
	}
}

func (d Disk) empty() []int {
	var idxs []int

	for idx, b := range d {
		if b.Type != Empty {
			continue
		}

		idxs = append(idxs, idx)
	}

	return idxs
}

func (d Disk) allocated() []int {
	var idxs []int

	for idx, b := range d {
		if b.Type != Allocated {
			continue
		}

		idxs = append(idxs, idx)
	}

	return idxs
}

func (d Disk) isDone() bool {
	var em bool

	for _, b := range d {
		switch b.Type {
		case Allocated:
			if em {
				return false
			}
		case Empty:
			em = true
		}
	}

	return true
}

func (d Disk) checksum() int {
	var sum int

	for idx, b := range d {
		if b.Type == Empty {
			break
		}

		sum += b.Value * idx
	}

	return sum
}

func (d Disk) swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
