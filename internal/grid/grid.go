package grid

import (
	"errors"
	"slices"
	"strings"
)

var AllDirectionsClockwise = []Direction{
	DirectionUp,  DirectionTopRight, DirectionRight, DirectionBottomRight, DirectionDown, DirectionBottomLeft, DirectionLeft, DirectionTopLeft, 
}

var DiagonalDirections = []Direction{
	DirectionTopLeft, DirectionTopRight, DirectionBottomLeft, DirectionBottomRight,
}

type Grid[T comparable] struct {
	rows [][]T
}

func (g Grid[T]) FindAllCoords(want T) []Coord {
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

func (g Grid[T]) FindCoord(want T) (Coord, error) {
	for i, row := range g.rows {
		for j, val := range row {
			if val == want {
				return Coord{x: j, y: i}, nil
			}
		}
	}
	return Coord{}, errors.New("could not find value in grid")
}

func (g Grid[T]) At(c Coord) T {
	return g.rows[c.y][c.x]
}

func (g Grid[T]) SearchNeighbours(start Coord, want T, directions []Direction) (map[Direction]Coord) {
	found := make(map[Direction]Coord)
	for _, d := range directions {
		neighbour := start.Neighbour(d)
		if g.IsInBounds(neighbour) && g.At(neighbour) == want {
			found[d] = neighbour
		}
	}
	return found
}

func (g Grid[T]) FindSeq(start Coord, d Direction, seq []T) bool {
	searchFrom := start
	for _, val := range seq {
		coord := searchFrom.Neighbour(d)
		if !g.IsInBounds(coord) {
			return false
		}
		if g.At(coord) != val {
			return false
		}
		searchFrom = coord
	}
	return true
}

func (g Grid[T]) IsInBounds(c Coord) bool {
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

func (d Direction) Opposite() Direction {
	idx := slices.Index(AllDirectionsClockwise, d)
	oppIdx := (idx + 4) % 8
	return AllDirectionsClockwise[oppIdx]
}

func (d Direction) Rotate90DegreesClockwise() Direction {
	idx := slices.Index(AllDirectionsClockwise, d)
	oppIdx := (idx + 2) % 8
	return AllDirectionsClockwise[oppIdx]
}

func NewRuneGrid(input string) Grid[rune] {
	var rows [][]rune
	for _, line := range strings.Split(input, "\n") {
		rows = append(rows, []rune(line))
	}
	return Grid[rune]{rows}
}