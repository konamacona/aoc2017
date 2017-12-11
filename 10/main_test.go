package main

import "testing"
import "reflect"

type testpair struct {
	input  string
	output string
}

var tests = []testpair{
	{"", "a2582a3a0e66e6e86e3812dcb672a272"},
	{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
	{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
	{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
}

func TestGetScore(t *testing.T) {
	for _, test := range tests {
		result := Part2(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Error(
				"For", test.input,
				"Expected", test.output,
				"Got", result,
			)
		}
	}
}

// a2582a3a0e66e6e86e3812dcb672a272
// a2582a3a e66e6e86e3812dcb672a272

// 3efbe78a8d82f29979031a4aa0b16a9d
// 3efbe78a8d82f29979 31a4aa0b16a9d

// 63960835bcdc130f0b66d7ff4f6a5a8e
// 6396 835bcdc13 f b66d7ff4f6a5a8e
