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

	fmt.Println("p1: ", p1(strings.TrimSuffix(input, "\n")))
	fmt.Println("p2: ", p2(strings.TrimSuffix(input, "\n")))
}

func p1(input string) int64 {
	return calculateTotalOutputJoltage(input, 2)
}

func p2(input string) int64 {
	return calculateTotalOutputJoltage(input, 12)
}

func calculateTotalOutputJoltage(input string, batteries int) int64 {
	var totalOutputJoltage int64

	for bank := range strings.SplitSeq(input, "\n") {
		var joltageDigits []byte

		var offset int
		for remainingDigits := batteries; remainingDigits > 0; remainingDigits-- {
			available := bank[offset:]

			availableBytes := []byte(available)
			slices.Sort(availableBytes)

			for _, availableByte := range slices.Backward(availableBytes) {
				index := offset + strings.IndexByte(available, availableByte)
				if index <= len(bank)-remainingDigits {
					offset = index + 1
					joltageDigits = append(joltageDigits, bank[index])
					break
				}
			}
		}

		joltage, _ := strconv.ParseInt(string(joltageDigits), 10, 64)

		totalOutputJoltage += joltage
	}

	return totalOutputJoltage
}
