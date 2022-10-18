package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
)

func part1(numbers []int) {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == numbers[(i+1)%len(numbers)] {
			sum += numbers[i]
		}
	}
	fmt.Println(sum)
}

func part2(numbers []int) {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == numbers[(i+len(numbers)/2)%len(numbers)] {
			sum += numbers[i]
		}
	}
	fmt.Println(sum)
}

func main() {
	numbers := input.LoadDefaultDigits()
	part1(numbers)
	part2(numbers)
}
