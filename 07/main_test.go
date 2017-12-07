package main

import "testing"

type testpair struct {
	input  []tower
	output string
	weight int
}

var part1tests = []testpair{
	{[]tower{
		{"pbga", 66, []string{}, 0},
		{"xhth", 57, []string{}, 0},
		{"ebii", 61, []string{}, 0},
		{"havc", 66, []string{}, 0},
		{"ktlj", 57, []string{}, 0},
		{"fwft", 72, []string{"ktlj", "cntj", "xhth"}, 0},
		{"qoyq", 66, []string{}, 0},
		{"padx", 45, []string{"pbga", "havc", "qoyq"}, 0},
		{"tknk", 41, []string{"ugml", "padx", "fwft"}, 0},
		{"jptl", 61, []string{}, 0},
		{"ugml", 68, []string{"gyxo", "ebii", "jptl"}, 0},
		{"gyxo", 61, []string{}, 0},
		{"cntj", 57, []string{}, 0},
	}, "tknk", 60},
}

func TestGetBottom(t *testing.T) {
	for _, test := range part1tests {
		result := GetBottom(test.input)
		if result != test.output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}

func TestGetImbalance(t *testing.T) {
	for _, test := range part1tests {
		result := GetImbalance(test.input)
		if result != test.weight {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}
