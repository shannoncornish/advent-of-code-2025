package main

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
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
	fmt.Println("p2:", p2(strings.TrimSuffix(input, "\n")))
}

type InclusiveRange struct {
	start, end int
}

func (r InclusiveRange) Contains(i int) bool {
	return i >= r.start && i <= r.end
}

func p1(input string) int {
	rangesInput, ingredientsInput, _ := strings.Cut(input, "\n\n")

	var ranges []InclusiveRange
	for rangeInput := range strings.SplitSeq(rangesInput, "\n") {
		startString, endString, _ := strings.Cut(rangeInput, "-")

		start, _ := strconv.Atoi(startString)
		end, _ := strconv.Atoi(endString)

		ranges = append(ranges, InclusiveRange{start, end})
	}

	var count int
	for ingredientString := range strings.SplitSeq(ingredientsInput, "\n") {
		ingredient, _ := strconv.Atoi(ingredientString)

		for _, r := range ranges {
			if r.Contains(ingredient) {
				count++
				break
			}
		}
	}

	return count
}

func p2(input string) int {
	rangesInput, _, _ := strings.Cut(input, "\n\n")

	var ranges []InclusiveRange
	var boundaries []int
	for rangeInput := range strings.SplitSeq(rangesInput, "\n") {
		startString, endString, _ := strings.Cut(rangeInput, "-")

		start, _ := strconv.Atoi(startString)
		end, _ := strconv.Atoi(endString)

		ranges = append(ranges, InclusiveRange{start, end})

		boundaries = append(boundaries, start, end+1)
	}

	slices.Sort(boundaries)

	var sum int
	for i := 0; i < len(boundaries)-1; i++ {
		current := boundaries[i]
		next := boundaries[i+1]

		for _, r := range ranges {
			if r.Contains(current) {
				sum += next - current
				break
			}
		}
	}

	return sum
}
