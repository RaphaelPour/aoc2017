package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

func part1(passphrases []string) int {
	sum := 0

	for _, pass := range passphrases {
		hist := make(map[string]struct{})
		valid := true
		for _, word := range strings.Split(pass, " ") {
			if _, found := hist[word]; found {
				valid = false
				break
			}
			hist[word] = struct{}{}
		}

		if valid {
			sum++
		}
	}
	return sum
}

func hasAnagram(words []string) bool {
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isAnagram(words[i], words[j]) {
				fmt.Println(words[i], words[j])
				return true
			}
		}
	}
	return false
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aList := strings.Split(a, "")
	sort.Strings(aList)

	bList := strings.Split(b, "")
	sort.Strings(bList)

	return strings.Join(aList, "") == strings.Join(bList, "")
}

func part2(passphrases []string) int {
	sum := 0
	for _, pass := range passphrases {
		if !hasAnagram(strings.Split(pass, " ")) {
			sum++
		}
	}
	return sum
}

func main() {
	passphrases := input.LoadDefaultString()
	fmt.Println(part1(passphrases))
	fmt.Println(part2(passphrases))
}
