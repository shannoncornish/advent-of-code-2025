package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
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
}

type Position struct {
	row, column int
}

type Shape map[Position]struct{}

type Region struct {
	width, length int
	quantities    []int
}

func p1(input string) int {
	components := strings.Split(input, "\n\n")

	shapes := make(map[int]Shape)
	for _, component := range components[0 : len(components)-1] {
		lines := strings.Split(component, "\n")

		shape := make(Shape)
		for row, line := range lines[1:] {
			for column, char := range line {
				if char == '#' {
					shape[Position{row, column}] = struct{}{}
				}
			}
		}

		id, _ := strconv.Atoi(strings.TrimSuffix(lines[0], ":"))
		shapes[id] = shape
	}

	var regions []Region
	for line := range strings.SplitSeq(components[len(components)-1], "\n") {
		sizeString, quantitiesString, _ := strings.Cut(line, ": ")

		parts := strings.Split(sizeString, "x")
		width, _ := strconv.Atoi(parts[0])
		length, _ := strconv.Atoi(parts[1])

		quantitiesStrings := strings.Split(quantitiesString, " ")

		quantities := make([]int, len(quantitiesStrings))
		for i, quantityString := range quantitiesStrings {
			quantity, _ := strconv.Atoi(quantityString)
			quantities[i] = quantity
		}

		regions = append(regions, Region{
			width,
			length,
			quantities,
		})
	}

	var count int
	for _, region := range regions {
		area := region.width * region.length

		var minimalArea int
		for i, quantity := range region.quantities {
			minimalArea += (len(shapes[i]) * quantity)
		}

		if minimalArea < area {
			count++
		}
	}

	return count
}
