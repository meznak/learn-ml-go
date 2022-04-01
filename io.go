package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x []int
	y float64
}

func ReadInputLines(fname string) []string {
	lines := make([]string, 0)

	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadInputInts(fname string) []int {
	lines := ReadInputLines(fname)
	ints := make([]int, len(lines))
	var err error

	for i, v := range lines {
		ints[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}

	return ints
}

func ReadInputPoints(fname string) []point {
	lines := ReadInputLines(fname)
	points := make([]point, len(lines))

	for i, v := range lines {
		splitLine := strings.Split(v, ";")
		xString := strings.Split(splitLine[0], ",")

		var p point
		p.x = make([]int, len(xString))

		for j, vString := range xString {
			x, err := strconv.Atoi(vString)
			if err != nil {
				panic(err)
			}
			p.x[j] = x
		}

		y, err := strconv.ParseFloat(splitLine[1], 32)
		if err != nil {
			panic(err)
		}
		p.y = y

		points[i] = p

	}

	return points
}
