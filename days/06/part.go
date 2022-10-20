package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	star_strings "github.com/RaphaelPour/stellar/strings"
)

type Memory struct {
	Banks   []int
	history map[string]int
	cycles  int
}

func NewMemory(in string) Memory {
	banks := make([]int, 0)

	for _, bank := range strings.Split(in, "\t") {
		banks = append(banks, star_strings.ToInt(bank))
	}

	m := Memory{
		Banks: banks,
	}
	m.history = make(map[string]int)

	return m
}

func (m Memory) String() string {
	s := make([]string, len(m.Banks))

	for i := range m.Banks {
		s[i] = fmt.Sprintf("%d", m.Banks[i])
	}
	return strings.Join(s, " ")
}

func (m Memory) NextBank() int {
	i, max := -1, 0
	for j, blocks := range m.Banks {
		if blocks > max {
			i, max = j, blocks
		}
	}
	return i
}

func (m *Memory) Next() (int, bool) {
	m.history[m.String()] = m.cycles

	nextBankI := m.NextBank()
	offset := m.Banks[nextBankI]
	m.Banks[nextBankI] = 0

	for i := nextBankI + 1; i < nextBankI+1+offset; i++ {
		m.Banks[i%len(m.Banks)]++
	}

	m.cycles++
	return m.HasInfiniteLoop()
}

func (m Memory) HasInfiniteLoop() (int, bool) {
	cycles, ok := m.history[m.String()]
	if !ok {
		return -1, false
	}
	return m.cycles - cycles, true
}

func part1(numbers string) (int, int) {
	mem := NewMemory(numbers)
	cycles := 0
	for {
		cycles++
		if val, ok := mem.Next(); ok {
			return cycles, val
		}
	}
	return cycles, -1
}

func main() {
	numbers := input.LoadString("input")[0]
	cycles, size := part1(numbers)
	fmt.Printf("cycles: %d\nsize: %d\n", cycles, size)
}
