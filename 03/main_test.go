package main

import "testing"

type testpair struct {
	input  int
	output int
}

var testsChecksum = []testpair{
	{1, 0},
	{10, 3},
	{12, 3},
	{17, 4},
	{23, 2},
	{25, 4},
	{1024, 31},
}

func TestManDist(t *testing.T) {
	for _, test := range testsChecksum {
		result := ManDist(test.input)
		if result != test.output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}
