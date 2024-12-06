package day4

import (
	"fmt"
	"strings"

	"github.com/dredly/aoc2024/internal/files"
	"golang.org/x/exp/maps"
)

func Part1Answer() {
	input := files.MustRead("inputdata/day4/real.txt")
	fmt.Printf("Day 4 part 1 answer: %d\n", numXmasAppearances(input))
}

func numXmasAppearances(input string) int {
	var total int
	g := NewRuneGrid(input)
	xCoords := g.findAllCoords('X')
	for _, c := range xCoords {
		mCoords := g.searchNeighbours(c, 'M')
		for _, d := range maps.Keys(mCoords) {
			if g.findSeq(mCoords[d], d, []rune("AS")) {
				total++
			}
		}
	}
	return total
}

var allDirections = []Direction{
	DirectionUp, DirectionDown, DirectionLeft, DirectionRight, DirectionTopLeft, DirectionTopRight, DirectionBottomLeft, DirectionBottomRight,
}

type Grid[T comparable] struct {
	rows [][]T
}

func (g Grid[T]) findAllCoords(want T) []Coord {
	var found []Coord
	for i, row := range g.rows {
		for j, val := range row {
			if val == want {
				found = append(found, Coord{x: j, y: i})
			}
		}
	}
	return found
}

func (g Grid[T]) at(c Coord) T {
	return g.rows[c.y][c.x]
}

func (g Grid[T]) searchNeighbours(start Coord, want T) (map[Direction]Coord) {
	found := make(map[Direction]Coord)
	for _, d := range allDirections {
		neighbour := start.Neighbour(d)
		if g.isInBounds(neighbour) && g.at(neighbour) == want {
			found[d] = neighbour
		}
	}
	return found
}

func (g Grid[T]) findSeq(start Coord, d Direction, seq []T) bool {
	searchFrom := start
	for _, val := range seq {
		coord := searchFrom.Neighbour(d)
		if !g.isInBounds(coord) {
			return false
		}
		if g.at(coord) != val {
			return false
		}
		searchFrom = coord
	}
	return true
}

func (g Grid[T]) isInBounds(c Coord) bool {
	return c.x >= 0 && c.x < len(g.rows[0]) && c.y >= 0 && c.y < len(g.rows)
}

type Coord struct { x, y int }

func (c Coord) Neighbour(d Direction) Coord {
	switch d {
	case DirectionUp:
		return c.Above()
	case DirectionDown:
		return c.Below()
	case DirectionLeft:
		return c.Left()
	case DirectionRight:
		return c.Right()
	case DirectionTopLeft:
		return c.Left().Above()
	case DirectionTopRight:
		return c.Right().Above()
	case DirectionBottomLeft:
		return c.Left().Below()
	case DirectionBottomRight:
		return c.Right().Below()
	}
	panic("unrecognised direction")
}

func (c Coord) Above() Coord {
	return Coord{
		x: c.x, y: c.y - 1,
	}
}

func (c Coord) Below() Coord {
	return Coord{
		x: c.x, y: c.y + 1,
	}
}

func (c Coord) Left() Coord {
	return Coord{
		x: c.x - 1, y: c.y,
	}
}

func (c Coord) Right() Coord {
	return Coord{
		x: c.x + 1, y: c.y,
	}
}

type Direction string

const (
	DirectionUp Direction = "UP"
	DirectionDown Direction = "DOWN"
	DirectionLeft Direction = "LEFT"
	DirectionRight Direction = "RIGHT"
	DirectionTopLeft Direction = "TOPLEFT"
	DirectionTopRight Direction = "TOPRIGHT"
	DirectionBottomLeft Direction = "BOTTOMLEFT"
	DirectionBottomRight Direction = "BOTTOMRIGHT"
)

func NewRuneGrid(input string) Grid[rune] {
	var rows [][]rune
	for _, line := range strings.Split(input, "\n") {
		rows = append(rows, []rune(line))
	}
	return Grid[rune]{rows}
}