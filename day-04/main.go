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

type Position struct {
	row, column int
}

var adjacentOffsets = []Position{
	{-1, -1}, {-1, +0}, {-1, +1},
	{+0, -1} /*+0,+0*/, {+0, +1},
	{+1, -1}, {+1, +0}, {+1, +1},
}

func main() {
	if os.Args[len(os.Args)-1] == "example" {
		input = example
	}

	fmt.Println("p1:", p1(strings.TrimSuffix(input, "\n")))
	fmt.Println("p2:", p2(strings.TrimSuffix(input, "\n")))
}

func p1(input string) int {
	var count int

	rolls := make(map[Position]struct{})
	for row, line := range strings.Split(input, "\n") {
		for column, char := range line {
			if char == '@' {
				rolls[Position{row, column}] = struct{}{}
			}
		}
	}

	for roll := range rolls {
		var adjacent int
		for _, adjacentOffset := range adjacentOffsets {
			position := Position{roll.row + adjacentOffset.row, roll.column + adjacentOffset.column}
			if _, found := rolls[position]; found {
				adjacent++
			}
		}

		if adjacent < 4 {
			count++
		}
	}

	return count
}

func p2(input string) int {
	var count int

	rolls := make(map[Position]struct{})
	for row, line := range strings.Split(input, "\n") {
		for column, char := range line {
			if char == '@' {
				rolls[Position{row, column}] = struct{}{}
			}
		}
	}

	candidates := rolls

	for {
		var removable []Position
		for candidate := range candidates {
			var adjacent int
			for _, adjacentOffset := range adjacentOffsets {
				position := Position{candidate.row + adjacentOffset.row, candidate.column + adjacentOffset.column}
				if _, found := rolls[position]; found {
					adjacent++
				}
			}

			if adjacent < 4 {
				removable = append(removable, candidate)
			}
		}

		if len(removable) == 0 {
			break
		}

		count += len(removable)

		for _, roll := range removable {
			delete(rolls, roll)
		}

		candidates = make(map[Position]struct{})
		for _, roll := range removable {
			for _, adjacentOffset := range adjacentOffsets {
				position := Position{roll.row + adjacentOffset.row, roll.column + adjacentOffset.column}
				if _, found := rolls[position]; found {
					candidates[position] = struct{}{}
				}
			}
		}
	}

	return count
}
