package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type tower struct {
	name       string
	baseWeight int
	children   []string

	actualWeight int
}

func main() {
	var towers []tower
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var t tower
		fmt.Sscanf(line, "%s (%d) -> ", &t.name, &t.baseWeight)

		if len(line) > (len(t.name) + 4) {
			halves := strings.Split(line, "-> ")
			if len(halves) > 1 {
				t.children = strings.Split(halves[1], ", ")
			}
		}

		towers = append(towers, t)
		// fmt.Println(t.name, t.baseWeight, t.children)
	}

	fmt.Println(GetBottom(towers))
	fmt.Println(GetImbalance(towers))
}

func GetBottom(input []tower) string {
	// Towers remaining as valid possibilities
	valid := make(map[string]bool)

	for _, t := range input {
		// Throw out anything without children
		if len(t.children) > 0 {

			// Invalidate children
			for _, c := range t.children {
				valid[c] = false
			}

			// Add as valid if it hasn't been already
			value, exists := valid[t.name]
			valid[t.name] = (value || !exists)
		}
	}

	result := "FAILED"

	// Anything left is the answer
	for name, isValid := range valid {
		if isValid {
			result = name
		}
	}

	return result
}

func GetImbalance(input []tower) int {
	// Towers by name
	towers := make(map[string]tower)

	for _, t := range input {
		towers[t.name] = t
	}

	bottom := GetBottom(input)

	weight, err := GetWeight(towers[bottom], towers)

	if err != -1 {
		return err
	}

	// Should never get here
	return weight
}

func GetWeight(t tower, m map[string]tower) (int, int) {
	if t.actualWeight != 0 {
		return t.actualWeight, -1
	}

	if len(t.children) == 0 {
		t.actualWeight = t.baseWeight
		return t.actualWeight, -1
	}

	result := t.baseWeight
	var cw = make([]int, len(t.children))
	var err int
	for i, c := range t.children {
		cw[i], err = GetWeight(m[c], m)
		if err != -1 {
			// bubble up
			return -1, err
		}
		result += cw[i]
	}

	for i, v := range cw {
		if i > 0 && v != cw[i-1] {
			expected := cw[i-1]
			actual := v
			actualName := t.children[i]
			for j, jv := range cw {
				if j != i && v == jv {
					expected = v
					actual = cw[i-1]
					actualName = t.children[i-1]
				}
			}

			answer := m[actualName].baseWeight + (expected - actual)
			return -1, answer
		}
	}

	return result, -1
}
