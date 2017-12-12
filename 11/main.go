package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)

	commands := strings.Split(input, ",")
	fmt.Println(follow(commands))
}

//
//      \______/
//      /      \
//   __/        \__
//     \        /
//      \______/
//      /       \
//

//   \ n  /
// nw +--+ ne
//   /    \
// -+      +-
//   \    /
// sw +--+ se
//   / s  \

// dir |  q  |  r
// ----+-----+-----
// n   |   0 |  -1
// ne  |  +1 |  -1
// se  |  +1 |   0
// s   |   0 |  +1
// sw  |  -1 |   1
// nw  |  -1 |   0

func follow(input []string) (int, int) {
	q := 0
	r := 0

	maxDist := 0
	var x, y, z int

	for _, d := range input {
		switch d {
		case "n":
			r--
		case "ne":
			q++
			r--
		case "se":
			q++
		case "s":
			r++
		case "sw":
			q--
			r++
		case "nw":
			q--
		}

		x, y, z := axial_to_cube(q, r)
		maxDist = max(abs(x), abs(y), abs(z), maxDist)
	}

	x, y, z = axial_to_cube(q, r)
	// fmt.Println(input, q, r, "->", x, y, z)

	return max(abs(x), abs(y), abs(z)), maxDist
}

func axial_to_cube(q int, r int) (int, int, int) {
	var x = q
	var z = r
	var y = -x - z
	return x, y, z
}

func abs(i int) int {
	if i >= 0 {
		return i
	}

	return -1 * i
}

func max(args ...int) int {
	result := args[0]
	for _, x := range args {
		if x > result {
			result = x
		}
	}
	return result
}
