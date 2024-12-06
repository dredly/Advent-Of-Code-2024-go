package day6

import (
	"fmt"

	"github.com/dredly/aoc2024/internal/files"
	"github.com/dredly/aoc2024/internal/grid"
)

func Part1Answer() {
	input := files.MustRead("inputdata/day6/real.txt")
	fmt.Printf("Day 6 part 1 answer: %d\n", distinctPositionsVisited(input))
}

func distinctPositionsVisited(input string) int {
	g := grid.NewRuneGrid(input)
	start, _ := g.FindCoord('^')
	visited := NewSet([]grid.Coord{start})
	current := start
	travelDirection := grid.DirectionUp
	for {
		next := current.Neighbour(travelDirection)
		if !g.IsInBounds(next) {
			break
		}
		if g.At(next) == '#' {
			travelDirection = travelDirection.Rotate90DegreesClockwise()
			continue
		}
		current = next
		visited.Add(current)
	}
	return visited.Size()
}

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](values []T) Set[T] {
	m := make(map[T]struct{})
	for _, v := range values {
		m[v] = struct{}{}
	}
	return Set[T]{m}
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Add(val T) {
	s.m[val] = struct{}{}
}