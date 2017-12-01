package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	input  string
	output int
}

type testarrayrpair struct {
	input  string
	output []int
}

var getValueTests = []testarrayrpair{
	{"1122", []int{1, 1, 2, 2}},
	{"1111", []int{1, 1, 1, 1}},
	{"1234", []int{1, 2, 3, 4}},
	{"91212129", []int{9, 1, 2, 1, 2, 1, 2, 9}},
}

func TestGetValueArray(t *testing.T) {
	for _, test := range getValueTests {
		result := GetValueArray(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}

var part1tests = []testpair{
	{"1122", 3},
	{"1111", 4},
	{"1234", 0},
	{"91212129", 9},
}

func TestPart1(t *testing.T) {
	for _, test := range part1tests {
		result := Part1(GetValueArray(test.input))
		if result != test.output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}

var part2tests = []testpair{
	{"1212", 6},
	{"1221", 0},
	{"123425", 4},
	{"123123", 12},
	{"12131415", 4},
}

func TestPart2(t *testing.T) {
	for _, test := range part2tests {
		result := Part2(GetValueArray(test.input))
		if result != test.output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}
