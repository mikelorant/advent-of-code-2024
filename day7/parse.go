package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(fn string) (Equations, error) {
	eqs := make(Equations, 0)

	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		value, nums := equation(scanner.Text())

		eqs[value] = nums
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return eqs, nil
}

func equation(str string) (int, Numbers) {
	ss := strings.Split(str, ":")

	value := mustInt(ss[0])
	nums := toNumbers(ss[1])

	return value, nums
}

func toNumbers(str string) Numbers {
	var nums Numbers

	str = strings.TrimSpace(str)

	for _, i := range strings.Split(str, " ") {
		nums = append(nums, mustInt(i))
	}

	return nums
}

func mustInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("unable to convert string to int: %v", err.Error())
	}

	return i
}
