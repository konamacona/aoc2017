package main

import (
	"fmt"
)

func main() {
	inputPart1 := []int32{
		230, 1, 2, 221, 97, 252, 168, 169, 57, 99, 0, 254, 181, 255, 235, 167,
	}

	inputPart2 := "230,1,2,221,97,252,168,169,57,99,0,254,181,255,235,167"

	fmt.Println(Part1(inputPart1), Part2(inputPart2))
}

func Part1(input []int32) int {
	arr, _, _ := Knot(GetStartingArray(256), input, 0, 0)
	return arr[0] * arr[1]
}

func Knot(arr []int, input []int32, start int32, skip int32) ([]int, int32, int32) {
	sz := int32(len(arr))

	for _, s := range input {
		for i := int32(0); i < s/int32(2); i++ {
			a := (start + i) % sz
			b := (start + (s - 1 - i)) % sz
			// fmt.Println("\t", s, i, a, b)
			arr[a], arr[b] = arr[b], arr[a]
		}

		start += s + skip
		skip++

		// fmt.Println(arr, start, skip)
	}

	return arr, start, skip
}

func GetStartingArray(sz int) []int {
	var r = make([]int, sz)
	for i := 0; i < sz; i++ {
		r[i] = i
	}
	return r
}

func Part2(input string) string {
	chars := []int32([]rune(input))
	additionalInput := []int32{17, 31, 73, 47, 23}
	chars = append(chars, additionalInput...)

	arr := GetStartingArray(256)
	var start int32
	var skip int32

	// Create sparse hash
	for i := 0; i < 64; i++ {
		arr, start, skip = Knot(arr, chars, start, skip)
	}

	// Convert to dense hash
	d := make([]int, len(arr)/16)
	hex := ""
	for i := 0; i < len(d); i++ {
		ai := i * 16
		for j := 0; j < 16; j++ {
			if j == 0 {
				d[i] = arr[ai+j]
			} else {
				d[i] = d[i] ^ arr[ai+j]
			}
		}

		hex += fmt.Sprintf("%02x", d[i])
	}

	// fmt.Println(d, hex)

	return hex
}
