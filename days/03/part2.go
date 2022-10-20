package main

import (
	"fmt"

	"github.com/RaphaelPour/stellar/math"
)

const (
	UP = iota
	LEFT
	DOWN
	RIGHT
)

type Spiral struct {
	fields    map[string]int
	length    int
	x, y      int
	direction int
}

func NewSpiral() Spiral {
	return Spiral{
		fields:    map[string]int{Key(0, 0): 1},
		direction: UP,
		length:    3,
	}
}

func Key(x, y int) string {
	return fmt.Sprintf("%d/%d", x, y)
}

func (s Spiral) CheckBoundary(x, y int) bool {
	return math.Abs(x) <= s.length/2 && math.Abs(y) <= s.length/2
}

func (s *Spiral) Next() int {
	switch s.direction {
	case UP:
		if s.CheckBoundary(s.x, s.y+1) {
			s.y++
		} else {
			s.direction = LEFT
			s.x--
		}
	case LEFT:
		if s.CheckBoundary(s.x-1, s.y) {
			s.x--
		} else {
			s.direction = DOWN
			s.y--
		}
	case DOWN:
		if s.CheckBoundary(s.x, s.y-1) {
			s.y--
		} else {
			s.direction = RIGHT
			s.x++
		}
	case RIGHT:
		if s.CheckBoundary(s.x+1, s.y) {
			s.x++
		} else {
			s.direction = UP
			s.length += 2
			s.y++
		}
	}

	return s.CalculateSum()
}

func (s *Spiral) CalculateSum() int {
	currentKey := Key(s.x, s.y)
	s.fields[currentKey] = 0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if y == 0 && x == 0 {
				continue
			}
			if val, ok := s.fields[Key(s.x+x, s.y+y)]; ok {
				s.fields[currentKey] += val
			}
		}
	}

	return s.fields[currentKey]
}

func main() {
	s := NewSpiral()
	for {
		if val := s.Next(); val > 368078 {
			fmt.Println(val)
			return
		}
	}
}
