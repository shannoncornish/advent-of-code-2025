package main

import (
	_ "embed"
	"fmt"
	"log"
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
	fmt.Println("p2:", p2(strings.TrimSuffix(input, "\n")))
}

func p1(input string) int {
	var sheet [][]string
	for line := range strings.SplitSeq(input, "\n") {
		fields := strings.Fields(line)
		sheet = append(sheet, fields)
	}

	symbols := sheet[len(sheet)-1]
	sheet = sheet[:len(sheet)-1]

	var grantTotal int

	problems := len(symbols)
	for i := range problems {
		var numbers []int
		for _, line := range sheet {
			numberString := line[i]

			number, _ := strconv.Atoi(numberString)
			numbers = append(numbers, number)
		}

		switch symbols[i] {
		case "+":
			grantTotal += sum(numbers)
		case "*":
			grantTotal += product(numbers)
		default:
			log.Fatal("Unsupported symbol:", symbols[i])
		}
	}

	return grantTotal
}

func p2(input string) int {
	var sheet []string
	for line := range strings.SplitSeq(input, "\n") {
		sheet = append(sheet, line)
	}

	symbols := sheet[len(sheet)-1]
	sheet = sheet[:len(sheet)-1]

	var problems []problem

	maxRows := len(sheet)
	maxColumns := len(sheet[0])
	for column := range maxColumns {
		var digits []int

		for row := range maxRows {
			digitString := sheet[row][column]
			if digitString != ' ' {
				digit, _ := strconv.Atoi(string(digitString))
				digits = append(digits, digit)
			}
		}

		switch symbols[column] {
		case '+', '*':
			problems = append(problems, problem{symbols[column], nil})
		}

		if len(digits) > 0 {
			number := digitsToNumber(digits)
			problems[len(problems)-1].operands = append(problems[len(problems)-1].operands, number)
		}
	}

	var grantTotal int
	for _, problem := range problems {
		switch problem.operator {
		case '+':
			grantTotal += sum(problem.operands)
		case '*':
			grantTotal += product(problem.operands)
		default:
			log.Fatal("Unsupported symbol:", problem.operator)
		}
	}

	return grantTotal
}

type problem struct {
	operator byte
	operands []int
}

func digitsToNumber(digits []int) int {
	number := 0
	for _, digit := range digits {
		number = number*10 + digit
	}
	return number
}

func product(factors []int) int {
	product := 1
	for _, factor := range factors {
		product *= factor
	}

	return product
}

func sum(addends []int) int {
	var sum int
	for _, addend := range addends {
		sum += addend
	}
	return sum
}
