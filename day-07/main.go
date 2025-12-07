package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	if os.Args[len(os.Args)-1] == "example" {
		input = example
	}

	fmt.Println("p1:", p1(strings.TrimSuffix(input, "\n")))
	fmt.Println("p2:", p2(strings.TrimSuffix(input, "\n")))
}

type Position struct {
	row, column int
}

func p1(input string) int {
	lines := strings.Split(input, "\n")

	splitters := make(map[Position]struct{})
	for row, line := range lines {
		for column, char := range line {
			if char == '^' {
				splitters[Position{row, column}] = struct{}{}
			}
		}
	}

	start := strings.IndexByte(lines[0], 'S')

	tachyons := make(map[Position]struct{})
	tachyons[Position{0, start}] = struct{}{}

	var split int
	for range len(lines) {
		newTachyons := make(map[Position]struct{})
		for tachyon := range tachyons {
			if _, found := splitters[Position{tachyon.row + 1, tachyon.column}]; found {
				newTachyons[Position{tachyon.row + 1, tachyon.column - 1}] = struct{}{}
				newTachyons[Position{tachyon.row + 1, tachyon.column + 1}] = struct{}{}

				split++
			} else {
				newTachyons[Position{tachyon.row + 1, tachyon.column}] = struct{}{}
			}
		}

		tachyons = newTachyons
	}

	return split
}

func p2(input string) int {
	lines := strings.Split(input, "\n")

	splitters := make(map[Position]struct{})
	for row, line := range lines {
		for column, char := range line {
			if char == '^' {
				splitters[Position{row, column}] = struct{}{}
			}
		}
	}

	start := strings.IndexByte(lines[0], 'S')

	tachyons := make(map[Position]int)
	tachyons[Position{0, start}] = 1

	for range len(lines) {
		newTachyons := make(map[Position]int)
		for tachyon, timelines := range tachyons {
			if _, found := splitters[Position{tachyon.row + 1, tachyon.column}]; found {
				newTachyons[Position{tachyon.row + 1, tachyon.column - 1}] += timelines
				newTachyons[Position{tachyon.row + 1, tachyon.column + 1}] += timelines
			} else {
				newTachyons[Position{tachyon.row + 1, tachyon.column}] += timelines
			}
		}

		tachyons = newTachyons
	}

	var sum int
	for _, timelines := range tachyons {
		sum += timelines
	}

	return sum
}
