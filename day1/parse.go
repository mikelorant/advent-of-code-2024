package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parse(fn string) ([]List, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	lists := make([]List, 2)

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		pair, err := toPair(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("unable to convert to pairs: %w", err)
		}

		lists[0] = append(lists[0], pair[0])
		lists[1] = append(lists[1], pair[1])
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return lists, nil
}

func toPair(str string) ([]int, error) {
	var pair []int

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str, -1)

	for _, i := range matches {
		val, err := strconv.Atoi(i)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to integer: %w", err)
		}

		pair = append(pair, val)
	}

	return pair, nil
}
