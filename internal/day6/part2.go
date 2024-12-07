package day6

import (
	"fmt"

	"github.com/dredly/aoc2024/internal/files"
	"github.com/dredly/aoc2024/internal/grid"
)

func Part2Answer() {
	input := files.MustRead("inputdata/day6/real.txt")
	fmt.Printf("Day 6 part 2 answer: %d\n", numPossibleObstructionPositions(input))
}

func numPossibleObstructionPositions(input string) int {
	found := NewSet[grid.Coord](nil)
	g := grid.NewRuneGrid(input)
	start, _ := g.FindCoord('^')
	visited := map[grid.Coord][]grid.Direction{
		start: {grid.DirectionUp},
	}
	current := start
	travelDirection := grid.DirectionUp
	OuterLoop:
		for {
			next := current.Neighbour(travelDirection)
			if !g.IsInBounds(next) {
				break OuterLoop
			}
			if g.At(next) == '#' {
				travelDirection = travelDirection.Rotate90DegreesClockwise()
				continue OuterLoop
			}
			current = next

			// fmt.Printf("Checking what would happen if we turn right here (%+v)\n", current)
			currentIfObstacle := next
			travelDirectionIfObstacle := travelDirection.Rotate90DegreesClockwise()
			visitedIfObstacle := make(map[grid.Coord][]grid.Direction)
			for k, v := range visited {
				copiedVal := make([]grid.Direction, len(v))
				copy(copiedVal, v)
				visitedIfObstacle[k] = copiedVal
			}
			InnerLoop:
				for {
					nextIfObstacle := currentIfObstacle.Neighbour(travelDirectionIfObstacle)
					potentialObstacle := currentIfObstacle.Neighbour(travelDirection)
					if !g.IsInBounds(nextIfObstacle) {
						break InnerLoop
					}
					if g.At(nextIfObstacle) == '#' {
						travelDirectionIfObstacle = travelDirectionIfObstacle.Rotate90DegreesClockwise()
						continue InnerLoop
					}
					currentIfObstacle = nextIfObstacle
					prevDirections, ok := visitedIfObstacle[currentIfObstacle]
					if !ok {
						visitedIfObstacle[currentIfObstacle] = []grid.Direction{travelDirectionIfObstacle}
						continue InnerLoop
					}
					visitedIfObstacle[currentIfObstacle] = append(prevDirections, travelDirectionIfObstacle)
					for _, d := range prevDirections {
						if travelDirectionIfObstacle == d {
							found.Add(potentialObstacle)
							break InnerLoop
						}
					}
				}

			prevDirections, ok := visited[current]
			if ok {
				visited[current] = append(prevDirections, travelDirection)
			} else {
				visited[current] = []grid.Direction{travelDirection}
			}
			
		}
	return found.Size()
}