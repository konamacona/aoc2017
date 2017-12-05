package main

import "testing"
import "fmt"

type testpair struct {
	input  []int
	output int
}

var testsPart1 = []testpair{
	{[]int{0, 3, 0, 1, -3}, 5},
}

func TestPart1(t *testing.T) {
	for _, test := range testsPart1 {
		output := Part1(test.input)
		if output != test.output {
			t.Error(
				"For:", test.input,
				"Expected:", test.output,
				"Got:", output,
			)
		}
	}
}

var testsPart2 = []testpair{
	{[]int{0, 3, 0, 1, -3}, 10},
}

func TestPart2(t *testing.T) {
	for _, test := range testsPart2 {
		output := Part2(test.input)
		fmt.Println(test.input)
		if output != test.output {
			t.Error(
				"For:", test.input,
				"Expected:", test.output,
				"Got:", output,
			)
		}
	}
}
