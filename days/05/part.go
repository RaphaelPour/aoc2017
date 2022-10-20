package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

func part1(program []int) int {
	steps := 0
	for i := 0; i < len(program); steps++ {
		offset := program[i]
		program[i]++
		i += offset
	}
	return steps
}
func part2(program []int) int {
	steps := 0
	for i := 0; i < len(program); steps++ {
		offset := program[i]
		if offset >= 3 {
			program[i]--
		} else {
			program[i]++
		}
		i += offset
	}
	return steps
}

func main() {
	program := input.LoadInt("input")
	fmt.Println(part1(program))
	fmt.Println(part2(program))
}
