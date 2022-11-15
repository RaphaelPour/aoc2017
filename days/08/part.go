package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/strings"
	
)

type Operation func(a,b int) bool

var(
	pattern = regexp.MustCompile(`^(\w+)\s(inc|dec)\s([-]?\d+)\sif\s(\w+)\s([><=!]+)\s([-]?\d+)$`)
	vars = make(map[string]int)

	operations = map[string]Operation{
		"==": func(a,b int) bool { return a == b	},
		"<=": func(a,b int) bool { return a <= b	},
		"<": func(a,b int) bool { return a < b	},
		">=": func(a,b int) bool { return a >= b	},
		">": func(a,b int) bool { return a > b	},
		"!=": func(a,b int) bool { return a != b	},
	}
)

func get(key string) int {
	if val, ok := vars[key]; ok {
		return val
	}
	return 0
}


func compare(a,b int, opp string) bool {
	switch opp {
		case "==":
			return a == b
		case "<=":
			return a <= b
		case "<":
			return a < b
		case ">=":
			return a >= b
		case ">":
			return a > b
		case "!=":
			return a != b
	}
	fmt.Printf("bad comparison: %d %s %d\n", a, opp, b)
	return false
}

func part1(instructions []string) (int,int) {
	maxPart1, maxPart2 := 0,0
	for _, instruction := range instructions {
		match := pattern.FindStringSubmatch(instruction)
		if len(match) == 0 {
			fmt.Printf("expected %s to match regex but got no match\n", instruction)
			return -1,-1
		}

		if !operations[match[5]](get(match[4]), strings.ToInt(match[6])) {
			continue
		}

		key := match[1]
		operator := match[2]
		value := strings.ToInt(match[3])
		
		if operator == "dec" {
			value *=-1
		}
		vars[key] = get(key) + value

		if get(key) > maxPart2 {
			maxPart2  = get(key)
		}
	}	

	for _, val := range vars {
		if val > maxPart1 {
			maxPart1 = val
		}
	}
	return maxPart1, maxPart2
}

func main() {
	max := input.LoadString("input")
	fmt.Println(part1(max))
}
