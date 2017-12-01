package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	var part = -1
	flag.IntVar(&part, "part", -1, "Run part 1 or 2?")
	flag.Parse()

	var input string
	fmt.Scanln(&input)

	if part == 1 {
		fmt.Println(Part1(GetValueArray(input)))
	} else if part == 2 {
		fmt.Println(Part2(GetValueArray(input)))
	} else {
		log.Fatalf("ERROR: Part flag must be 1 or 2")
	}
}

// GetValueArray converts a string of integers into an array of ints
func GetValueArray(input string) []int {
	var values = make([]int, len(input))
	for i, r := range input {
		val, err := strconv.Atoi(string(r))
		if err != nil {
			log.Fatalf("ERROR: Could not convert list to integers")
		} else {
			values[i] = val
		}
	}
	return values
}

// Part1 converts an array of ints into a sum via the following rule:
// For each pair of matching siblings, add the number number value to the
// total. The list is circular, so include the comparison of last and first.
func Part1(values []int) int {
	var result int
	var last = len(values) - 1
	for i, r := range values {
		if i < last && r == values[i+1] {
			result += r
		} else if i == last && r == values[0] {
			result += r
		}
	}

	return result
}

// Part2 converts an array of ints into a sum via the following rule:
// For each pair of matching siblings*, add the number number value to the
// total. The list is circular, so include the comparison of first and last.
// In this case siblings are two numbers which are the furthest  from each other
// (length / 2 steps ahead) in the circular list, and the list will always be
// even length
func Part2(values []int) int {
	var result int
	var length = len(values)
	var offset = length / 2

	for i, r := range values {
		var siblingIndex = (i + offset) % length

		if r == values[siblingIndex] {
			result += r
		}
	}

	return result
}
