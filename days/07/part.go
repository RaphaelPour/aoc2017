package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	stellar_strings "github.com/RaphaelPour/stellar/strings"
	
)

var(
	//pattern = regexp.MustCompile(`^(\w+)\s+\((\d+)\)\s+->\s+(\w+|\w+(,\w+)+)$`)
	pattern = regexp.MustCompile(`^(\w+)\s+\((\d+)\)(\s+->\s+)?(.*)$`)
)

type Node struct {
	Name string
	Weight int
	Parent string
	Children []string
}

func part1(towers []string) string {
	nodes := make(map[string]*Node)
	for _, tower := range towers {
		match := pattern.FindStringSubmatch(tower)
		nodes[match[1]] = &Node{
			Name: match[1],
			Weight: stellar_strings.ToInt(match[2]),
			Children: strings.Split(match[4],", "),
		}
	}

	// Find children for any node
	for _, parent := range nodes {
		for _, other := range nodes {
			if parent.Name == other.Name {
				continue
			}

			for _, child:= range parent.Children {
				if child == other.Name {
					other.Parent = parent.Name
				}
			}
		} 
	}

	// Find Node without a parent
	for _, node := range nodes {
		if node.Parent == "" {
			fmt.Println(node.Name)
		}
	}
	return ""
}

func main() {
	towers := input.LoadString("input")
	fmt.Println(part1(towers))
}
