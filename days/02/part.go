package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
)

func part1(numbers [][]int) {
	sum := 0
	for _, row := range numbers {
		min, max := math.MinMax(row)
		sum += max - min
	}
	fmt.Println(sum)
}

func part2(numbers [][]int) {
	sum := 0
	for _, row := range numbers {
		found := false
		for i := 0; i < len(row) && !found; i++ {
			for j := i + 1; j < len(row) && !found; j++ {
				if row[i]%row[j] == 0 {
					sum += row[i] / row[j]
					found = true
				} else if row[j]%row[i] == 0 {
					sum += row[j] / row[i]
					found = true
				}
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	numbers := input.LoadDefaultIntTable()
	part1(numbers)
	part2(numbers)
}
