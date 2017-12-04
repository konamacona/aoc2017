package main

import (
	"fmt"
)

func main() {
	fmt.Println(ManDist(265149))
	fmt.Println(Part2(265149))
}

func ManDist(input int) int {
	var bottomRight, side int
	layers := 0

	// Find the layer, side length and bottomRight value
	for {
		layers++
		side = layers*2 - 1
		bottomRight = side * side

		if input <= bottomRight {
			break
		}
	}

	if layers == 1 {
		return 0
	}

	// Where the input number falls in the current circle
	progress := bottomRight - input

	// The distance from a corner to a center
	halfSide := (side - 1) / 2

	return (layers - 1) + Abs((progress%(side-1))-halfSide)
}

func Part2(input int) int {
	s := NewSpiral()
	s.SetCurrent(1)

	for {
		s.Next()
		val := GetSurroundingSum(s)
		s.SetCurrent(val)

		if val > input {
			return val
		}
	}
}

func GetSurroundingSum(s *Spiral) int {
	result := 0
	for iX := s.x - 1; iX <= (s.x + 1); iX++ {
		for iY := s.y - 1; iY <= (s.y + 1); iY++ {
			if !(iX == s.x && iY == s.y) {
				result += s.Get(iX, iY)
			}
		}
	}
	return result
}

func Abs(input int) int {
	if input >= 0 {
		return input
	}

	return input * -1
}
