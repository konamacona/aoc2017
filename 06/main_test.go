package main

import "testing"

type testpair struct {
	input   []int
	output1 int
	output2 int
}

var tests = []testpair{
	{[]int{0, 2, 7, 0}, 5, 4},
}

func TestBoth(t *testing.T) {
	for _, test := range tests {
		output1, output2 := Both(test.input)
		if output1 != test.output1 {
			t.Error(
				"For:", test.input,
				"Expected:", test.output1,
				"Got:", output1,
			)
		}

		if output2 != test.output2 {
			t.Error(
				"For:", test.input,
				"Expected:", test.output2,
				"Got:", output2,
			)
		}
	}
}
