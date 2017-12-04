package main

import (
	"reflect"
	"testing"
)

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

//      -2  -1   0  +1  +2
//    --------------------
// +2 | 17  16  15  14  13
// +1 | 18   5   4   3  12
//  0 | 19   6   1   2  11
// -1 | 20   7   8   9  10
// -2 | 21  22  23  24  25

var testSpiral = map[int]map[int]int{
	+2: map[int]int{
		+2: 13,
		+1: 12,
		0:  11,
		-1: 10,
		-2: 25,
	},
	+1: map[int]int{
		+2: 14,
		+1: 3,
		0:  2,
		-1: 9,
		-2: 24,
	},
	0: map[int]int{
		+2: 15,
		+1: 4,
		0:  1,
		-1: 8,
		-2: 23,
	},
	-1: map[int]int{
		+2: 16,
		+1: 5,
		0:  6,
		-1: 7,
		-2: 22,
	},
	-2: map[int]int{
		+2: 17,
		+1: 18,
		0:  19,
		-1: 20,
		-2: 21,
	},
}

func TestSpiralCreation(t *testing.T) {
	s := NewSpiral()
	s.SetCurrent(1)
	for i := 1; i < 25; i++ {
		s.Next()
		s.SetCurrent(i + 1)
	}

	// s.Print()

	if !reflect.DeepEqual(s.m, testSpiral) {
		t.Error(
			"For Spiral",
			"expected", testSpiral,
			"got", s.m,
		)
	}
}

// 147  142  133  122   59
// 304    5    4    2   57
// 330   10    1    1   54
// 351   11   23   25   26
// 362  747  806--->   ...

var testMap = []int{
	1, 1, 2, 4, 5, 10, 11, 23, 25, 26, 54, 57, 59, 122, 133, 142, 147, 304, 330, 351, 362, 747, 806,
}

func TestGetSurroundingSum(t *testing.T) {
	s := NewSpiral()
	s.SetCurrent(1)

	for i := 1; i < len(testMap); i++ {
		s.Next()
		val := GetSurroundingSum(s)
		s.SetCurrent(val)

		if val != testMap[i] {
			t.Error(
				"For ", i,
				"expected", testMap[i],
				"got", val,
			)
		}
	}
}
