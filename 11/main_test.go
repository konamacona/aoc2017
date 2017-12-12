package main

import "testing"
import "reflect"

type testpair struct {
	input  []string
	output int
}

var tests = []testpair{
	{[]string{"ne", "ne", "ne"}, 3},
	{[]string{"ne", "ne", "sw", "sw"}, 0},
	{[]string{"ne", "ne", "s", "s"}, 2},
	{[]string{"se", "sw", "se", "sw", "sw"}, 3},
}

func TestPart1(t *testing.T) {
	for _, test := range tests {
		result := Part1(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Error(
				"For", test.input,
				"Expected", test.output,
				"Got", result,
			)
		}
	}
}
