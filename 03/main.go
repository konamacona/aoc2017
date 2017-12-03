package main

import "fmt"

func main() {
	fmt.Printf("%d\n", ManDist(265149))
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

	// 	fmt.Printf(`input: %d
	// 	layer: %d
	// 	side: %d
	// 	bottomRight: %d
	// 	progress: %d
	// 	halfSide: %d
	// `, input, layers, side, bottomRight, progress, halfSide)

	return (layers - 1) + Abs((progress%(side-1))-halfSide)
}

func Abs(input int) int {
	if input >= 0 {
		return input
	}

	return input * -1
}
