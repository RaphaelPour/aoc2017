package main

import(
	"fmt"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/math"
)

func part1(number int) int {
	length := 1
	x,y := 0,0
	for {
		area := length*length
    if area >= number {
			lastArea := (length-2)*(length-2)
			ring := lastArea + 1
			fmt.Printf("n=%d, min=%d, max=%d\n", length, lastArea+1, area)

			if math.Within(number, ring, ring + length-2) {
				fmt.Println("right")
				// right
				x = length/2
				y = number - ring + length/2 -1 - lastArea
			} else if math.Within(number, ring + length -2, ring + length -2 + length) {
				fmt.Println("top")
				// top
				x = number - (ring + length - 2) - length/2 - lastArea
				y = length/2
			} else if math.Within(number, ring + length - 2 + length, ring + length - 2 + length + length -2) {
				fmt.Println("left")
				// left
				x = length/2
				y = number - (ring + length - 2 + length) - length/2 - lastArea
			} else {
				fmt.Println("bottom")
				// bottom
				//  22     - (
				// x = number - (ring + length - 2 + length + length -2 + 1) + length/2 - lastArea
				x = number - area + length/2
				y = length/2
			}
			fmt.Printf("dist(%d,%d) = %d\n", x,y, math.Abs(x) + math.Abs(y))
			return math.Abs(x) + math.Abs(y)
		}
		length +=2
	}
	return -1
}

func main() {
	number := input.LoadIntList("input")[0]
	fmt.Println("Solution: ", part1(number))
	fmt.Println(" too low: 303")
	fmt.Println("     bad: 337,372")
	fmt.Println("too high: 421")
}
