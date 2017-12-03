package main

import "testing"

type testpair struct {
	input  []int
	output int
}

var testsChecksum = []testpair{
	{[]int{5, 1, 9, 5}, 8},
	{[]int{7, 5, 3}, 4},
	{[]int{2, 4, 6, 8}, 6},
}

func TestChecksum(t *testing.T) {
	for _, test := range testsChecksum {
		result := Checksum(test.input)
		if result != test.output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}

var testsDivsum = []testpair{
	{[]int{5, 9, 2, 8}, 4},
	{[]int{9, 4, 7, 3}, 3},
	{[]int{3, 8, 6, 5}, 2},
}

func TestDivsum(t *testing.T) {
	for _, test := range testsDivsum {
		result := Divsum(test.input)
		if result != test.output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}
