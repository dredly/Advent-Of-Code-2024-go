package day4

import (
	"fmt"

	"github.com/dredly/aoc2024/internal/files"
	"github.com/dredly/aoc2024/internal/grid"
	"golang.org/x/exp/maps"
)

func Part2Answer() {
	input := files.MustRead("inputdata/day4/real.txt")
	fmt.Printf("Day 4 part 2 answer: %d\n", numCrossedMASAppearances(input))
}

func numCrossedMASAppearances(input string) int {
	var total int
	g := grid.NewRuneGrid(input)
	aCoords := g.FindAllCoords('A')
	for _, c := range(aCoords) {
		mCoords := g.SearchNeighbours(c, 'M', grid.DiagonalDirections)
		var foundMASes int
		for _, d := range maps.Keys(mCoords) {
			if g.IsInBounds(c.Neighbour(d.Opposite())) {
				opp := g.At(c.Neighbour(d.Opposite()))
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