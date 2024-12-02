package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(fn string) (Reports, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	var reps Reports

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		rep, err := toReport(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("unable to read report: %w", err)
		}

		reps = append(reps, rep)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return reps, nil
}

func toReport(str string) (Report, error) {
	var rep Report

	ss := strings.Split(str, " ")
	for _, s := range ss {
		lvl, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to integer: %w", err)
		}

		rep = append(rep, lvl)
	}

	return rep, nil
}
