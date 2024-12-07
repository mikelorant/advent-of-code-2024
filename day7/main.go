package main

import (
	"log"
	"strconv"
	"strings"
)

type (
	Equations map[int]Numbers
	Numbers   []int
	Operators []Operator
	Operator  int
)

const (
	Unset Operator = iota
	Add
	Multiply
	Concatenation
)

func main() {
	i := Task("input1.txt", 1)
	log.Println("Part 1:", i)

	j := Task("input1.txt", 2)
	log.Println("Part 2:", j)
}

func Task(file string, part int) int {
	eqs, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return sum(eqs, 2)
	case 2:
		return sum(eqs, 3)
	}

	return 0
}

func sum(eqs Equations, cmds int) int {
	var sum int

	for val, nums := range eqs {
		sum += result(val, nums, cmds)
	}

	return sum
}

func result(val int, nums Numbers, cmds int) int {
	var sum int

	opers := len(nums) - 1
	perms := power(cmds, opers)

	for n := range perms {
		ops := operators(n, opers, cmds)

		if evaluate(nums, ops) != val {
			continue
		}

		sum += val

		break
	}

	return sum
}

func operators(n, pad, cmds int) Operators {
	var ops Operators

	for _, char := range convert(n, cmds, pad) {
		ops = append(ops, toOperator(char))
	}

	return ops
}

func evaluate(nums Numbers, ops Operators) int {
	i := nums[0]

	for idx, num := range nums[1:] {
		i = calculate(i, ops[idx], num)
	}

	return i
}

func calculate(i int, op Operator, num int) int {
	switch op {
	case Add:
		return i + num
	case Multiply:
		return i * num
	case Concatenation:
		return mustInt(strconv.Itoa(i) + strconv.Itoa(num))
	}

	return 0
}

func power(x, y int) int {
	i := 1

	for range y {
		i *= x
	}

	return i
}

func toOperator(r rune) Operator {
	switch r {
	case '0':
		return Add
	case '1':
		return Multiply
	case '2':
		return Concatenation
	}

	return Unset
}

func (o Operator) String() string {
	switch o {
	case Add:
		return "+"
	case Multiply:
		return "*"
	case Concatenation:
		return "||"
	}

	return ""
}

func convert(n, base, pad int) string {
	str := strconv.FormatInt(int64(n), base)
	pad -= len(str)

	return strings.Repeat("0", pad) + str
}
