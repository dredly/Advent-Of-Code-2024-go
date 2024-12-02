package day1

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dredly/aoc2024/internal/files"
	"github.com/dredly/aoc2024/internal/maths"
	"github.com/dredly/aoc2024/internal/parsing"
)

func Part1Answer() {
	left, right := parse(files.MustRead("inputdata/day1/real.txt"))
	fmt.Printf("Day 1 part 1 answer: %d\n", totalDistance(left, right))
}

func parse(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		l, r, _ := strings.Cut(line, "   ")
		left[i] = parsing.MustParseInt(l)
		right[i] = parsing.MustParseInt(r)
	}
	return left, right
}

func totalDistance(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	var total int
	for i, locIDLeft := range left {
		locIDRight := right[i]
		total += maths.Abs(locIDRight - locIDLeft)
	}
	return total
}