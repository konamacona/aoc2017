package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		tokens := strings.Split(input, "\t")
		var values = make([]int, len(tokens))
		for i, t := range tokens {
			fmt.Println(t)
			value, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal("ERROR: Unable to parse input")
			}
			values[i] = value
		}
		fmt.Println(Both(values))
	}
}

func Both(input []int) (int, int) {
	result := 0
	configs := make(map[string]int)
	max := len(input)
	for {
		result++
		index := PickNext(input)
		total := input[index]
		input[index] = 0

		for total > 0 {
			index = (index + 1) % max
			input[index] += 1
			total -= 1
		}

		hash := Hash(input)
		i, exists := configs[hash]
		if exists {
			return result, result - i
		}
		configs[hash] = result
	}
}

func PickNext(input []int) int {
	max := math.MinInt32
	result := -1
	for i, v := range input {
		if v > max || (v == max && i < result) {
			result = i
			max = v
		}
	}
	return result
}

func Hash(input []int) string {
	// https://stackoverflow.com/a/37533144
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input)), ","), "[]")
}
