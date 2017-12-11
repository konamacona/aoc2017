package main

import "testing"

type testpair struct {
	input   string
	output1 int
	output2 int
}

var tests = []testpair{
	{"{}", 1, 0},
	{"{{{}}}", 6, 0},
	{"{{},{}}", 5, 0},
	{"{{{},{},{{}}}}", 16, 0},
	{"{<a>,<a>,<a>,<a>}", 1, 4},
	{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},
	{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},
	{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17},
	{"<>", 0, 0},
	{"<random characters>", 0, 17},
	{"<<<<>", 0, 3},
	{"<{!>}>", 0, 2},
	{"<!!>", 0, 0},
	{"<!!!>>", 0, 0},
	{"<{o\"i!a,<{i<a>", 0, 10},
	{"<!!!ei!!!><>", 0, 2},
}

func TestGetScore(t *testing.T) {
	for _, test := range tests {

		result1, result2 := GetScore([]rune(test.input))
		if result1 != test.output1 || result2 != test.output2 {
			t.Error(
				"For", test.input,
				"Expected", test.output1, test.output2,
				"Got", result1, result2,
			)
		}
	}
}
