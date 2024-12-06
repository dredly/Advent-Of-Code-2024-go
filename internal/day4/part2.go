package day4

import (
	"fmt"
	"slices"

	"github.com/dredly/aoc2024/internal/files"
	"golang.org/x/exp/maps"
)

var diagonalDirections = []Direction{
	DirectionTopLeft, DirectionTopRight, DirectionBottomLeft, DirectionBottomRight,
}

func Part2Answer() {
	input := files.MustRead("inputdata/day4/real.txt")
	fmt.Printf("Day 4 part 2 answer: %d\n", numCrossedMASAppearances(input))
}

func numCrossedMASAppearances(input string) int {
	var total int
	g := NewRuneGrid(input)
	aCoords := g.findAllCoords('A')
	for _, c := range(aCoords) {
		mCoords := g.searchNeighbours(c, 'M', diagonalDirections)
		var foundMASes int
		for _, d := range maps.Keys(mCoords) {
			if g.isInBounds(c.Neighbour(d.Opposite())) {
				opp := g.at(c.Neighbour(d.Opposite()))
				if opp == 'S' {
					foundMASes++
				}
			}
		}
		if foundMASes == 2 {
			total++
		}
	}
	return total
}

func (d Direction) Opposite() Direction {
	idx := slices.Index(allDirectionsClockwise, d)
	oppIdx := (idx + 4) % 8
	return allDirectionsClockwise[oppIdx]
}