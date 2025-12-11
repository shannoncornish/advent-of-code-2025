package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed example1.txt
var example1 string

//go:embed example2.txt
var example2 string

//go:embed input.txt
var input string

func main() {
	input1 := input
	input2 := input
	if os.Args[len(os.Args)-1] == "example" {
		input1 = example1
		input2 = example2
	}

	fmt.Println("p1:", p1(strings.TrimSuffix(input1, "\n")))
	fmt.Println("p2:", p2(strings.TrimSuffix(input2, "\n")))
}

func p1(input string) int {
	devices := make(map[string][]string)
	for line := range strings.SplitSeq(input, "\n") {
		device, outputs, _ := strings.Cut(line, ":")
		devices[device] = strings.Fields(outputs)
	}

	start := "you"
	goal := "out"

	cache := make(map[string]int)

	var fn func(current string) int
	fn = func(current string) int {
		if sum, found := cache[current]; found {
			return sum
		}

		if current == goal {
			return 1
		}

		var sum int
		for _, output := range devices[current] {
			sum += fn(output)
		}

		cache[current] = sum
		return sum
	}

	return fn(start)
}

type P2CacheKey struct {
	device   string
	dac, fft bool
}

func p2(input string) int {
	devices := make(map[string][]string)
	for line := range strings.SplitSeq(input, "\n") {
		device, outputs, _ := strings.Cut(line, ":")
		devices[device] = strings.Fields(outputs)
	}

	start := "svr"
	goal := "out"

	cache := make(map[P2CacheKey]int)

	var fn func(current string, dac, fft bool) int
	fn = func(current string, dac, fft bool) int {

		key := P2CacheKey{current, dac, fft}
		if sum, found := cache[key]; found {
			return sum
		}

		if dac && fft && current == goal {
			return 1
		}

		var sum int
		for _, output := range devices[current] {
			dac := dac || output == "dac"
			fft := fft || output == "fft"

			sum += fn(output, dac, fft)
		}

		cache[key] = sum
		return sum
	}

	return fn(start, false, false)
}
