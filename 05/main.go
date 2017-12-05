package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input1, input2 []int
	for scanner.Scan() {
		text := scanner.Text()

		val, err := strconv.Atoi(text)

		if err != nil {
			log.Fatal("ERROR: Non int string!")
		}

		input1 = append(input1, val)
		input2 = append(input2, val)
	}

	fmt.Println(Part1(input1))
	fmt.Println(Part2(input2))
}

func Part1(input []int) int {
	steps := 0
	i := 0
	max := len(input) - 1

	var jump int
	for {
		jump = input[i]
		input[i]++
		i += jump
		steps++
		if i < 0 || i > max {
			break
		}
	}
	return steps
}

func Part2(input []int) int {
	steps := 0
	i := 0
	max := len(input) - 1

	var jump int
	for {
		jump = input[i]
		if jump >= 3 {
			input[i]--
		} else {
			input[i]++
		}
		i += jump
		steps++
		if i < 0 || i > max {
			break
		}
	}
	return steps
}
