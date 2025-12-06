package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	fmt.Println("p1: ", p1(strings.TrimSuffix(input, "\n")))
	fmt.Println("p2: ", p2(strings.TrimSuffix(input, "\n")))
}

func p1(input string) int {
	sum := 0

	for r := range strings.SplitSeq(input, ",") {
		startString, endString, _ := strings.Cut(r, "-")
		start, _ := strconv.Atoi(startString)
		end, _ := strconv.Atoi(endString)

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			if s[0:len(s)/2] == s[(len(s)/2):] {
				sum += i
			}
		}
	}

	return sum
}

func p2(input string) int {
	sum := 0

	for r := range strings.SplitSeq(input, ",") {
		startString, endString, _ := strings.Cut(r, "-")
		start, _ := strconv.Atoi(startString)
		end, _ := strconv.Atoi(endString)

		for i := start; i <= end; i++ {
			if i < 10 {
				continue
			}

			s := strconv.Itoa(i)

			for chunkSize := len(s) / 2; chunkSize > 0; chunkSize-- {
				if len(s)%chunkSize != 0 {
					continue
				}

				repeating := true
				for j := range len(s) / chunkSize {
					if s[0:chunkSize] != s[j*chunkSize:j*chunkSize+chunkSize] {
						repeating = false
						break
					}
				}

				if repeating {
					sum += i
					break
				}
			}
		}
	}

	return sum
}
