package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var input []rune
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			input = append(input, c)
		}
	}

	fmt.Println(GetScore(input))
}

func GetScore(in []rune) (int, int) {
	// fmt.Println("\n")
	var g = false
	var score = 0
	var garbage = 0
	var depth = 0

	for i := 0; i < len(in); i++ {
		c := in[i]

		if g {
			garbage++
		}

		switch c {
		case '<':
			if !g {
				g = true
			}
		case '>':
			if g {
				g = false
				garbage--
			}
		case '{':
			if !g {
				depth++
				score += depth
			}
		case '}':
			if !g {
				depth--
			}
		case '!':
			if g {
				i++
				garbage--
			}
		}
	}

	return score, garbage
}
