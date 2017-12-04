package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	part1count := 0
	part2count := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if Part1Valid(text) {
			part1count++
		}
		if Part2Valid(text) {
			part2count++
		}
	}
	fmt.Println(part1count)
	fmt.Println(part2count)
}

func Part1Valid(input string) bool {
	tokens := strings.Split(input, " ")
	m := make(map[string]bool)
	for _, t := range tokens {
		_, exists := m[t]
		if exists {
			return false
		}
		m[t] = true
	}
	return true
}

func Part2Valid(input string) bool {
	tokens := strings.Split(input, " ")
	m := make(map[string]bool)
	for _, t := range tokens {

		split := strings.Split(t, "")
		sort.Strings(split)
		t = strings.Join(split, "")

		_, exists := m[t]
		if exists {
			return false
		}
		m[t] = true
	}
	return true
}
