package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type instruction struct {
	addr   string
	op     string
	amount int

	condAddr   string
	condOp     string
	condAmount int
}

func (i *instruction) PerformOp(reg map[string]int) {
	a := i.amount
	if i.op == "dec" {
		a *= -1
	}
	reg[i.addr] = reg[i.addr] + a
}

func (i *instruction) CheckCond(reg map[string]int) bool {
	a := reg[i.condAddr]
	b := i.condAmount
	switch i.condOp {
	case "<":
		return a < b
	case ">":
		return a > b
	case "<=":
		return a <= b
	case ">=":
		return a >= b
	case "==":
		return a == b
	case "!=":
		return a != b
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var program []instruction
	for scanner.Scan() {
		line := scanner.Text()
		var i instruction
		fmt.Sscanf(line, "%s %s %d if %s %s %d",
			&i.addr,
			&i.op,
			&i.amount,
			&i.condAddr,
			&i.condOp,
			&i.condAmount,
		)

		program = append(program, i)
	}

	mem, part2 := Run(program)
	fmt.Println(FindMax(mem))
	fmt.Println(part2)
}

// Run runs the program and returns the resulting memory set and the maximum
// value that was ever set in any register
func Run(p []instruction) (map[string]int, int) {
	reg := make(map[string]int)
	p2max := math.MinInt32

	// Run program
	for _, i := range p {
		if i.CheckCond(reg) {
			i.PerformOp(reg)

			if reg[i.addr] > p2max {
				p2max = reg[i.addr]
			} else if reg[i.condAddr] > p2max {
				p2max = reg[i.condAddr]
			}
		}
	}

	return reg, p2max
}

// FindMax Returns the maximum value in a map
func FindMax(reg map[string]int) int {
	max := math.MinInt32
	for _, v := range reg {
		if v > max {
			max = v
		}
	}
	return max
}
