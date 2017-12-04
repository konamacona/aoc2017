package main

import "testing"

type testpair struct {
	input  string
	output bool
}

var part1tests = []testpair{
	{"", true},
	{"aa bb cc dd ee", true},
	{"aa bb cc dd aa", false},
	{"aa bb cc dd aaa", true},
}

func TestPart1Valid(t *testing.T) {
	for _, test := range part1tests {
		result := Part1Valid(test.input)
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
	{"", true},
	{"abcde fghij", true},              // is a valid passphrase.
	{"abcde xyz ecdab", false},         // is not valid - the letters from the third word can be rearranged to form the first word.
	{"a ab abc abd abf abj", true},     // is a valid passphrase, because all letters need to be used when forming another word.
	{"iiii oiii ooii oooi oooo", true}, // is valid.
	{"oiii ioii iioi iiio", false},     // is not valid - any of these words can be rearranged to form any other word.
}

func TestPart2Valid(t *testing.T) {
	for _, test := range part2tests {
		result := Part2Valid(test.input)
		if result != test.output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}
