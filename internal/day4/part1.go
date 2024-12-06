package day4

import (
	"fmt"

	"github.com/dredly/aoc2024/internal/files"
	"github.com/dredly/aoc2024/internal/grid"
	"golang.org/x/exp/maps"
)

func Part1Answer() {
	input := files.MustRead("inputdata/day4/real.txt")
	fmt.Printf("Day 4 part 1 answer: %d\n", numXmasAppearances(input))
}

func numXmasAppearances(input string) int {
	var total int
	g := grid.NewRuneGrid(input)
	xCoords := g.FindAllCoords('X')
	for _, c := range xCoords {
		mCoords := g.SearchNeighbours(c, 'M', grid.AllDirectionsClockwise)
		for _, d := range maps.Keys(mCoords) {
			if g.FindSeq(mCoords[d], d, []rune("AS")) {
				total++
			}
		}
	}
	return total
}

