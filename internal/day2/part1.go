package day2

import (
	"fmt"
	"strings"

	"github.com/dredly/aoc2024/internal/files"
	"github.com/dredly/aoc2024/internal/maths"
	"github.com/dredly/aoc2024/internal/parsing"
)

func Part1Answer() {
	input := files.MustRead("inputdata/day2/real.txt")
	fmt.Printf("Day 1 part 1 answer: %d\n", numSafeReports(input, isSafeBasic))
}

func numSafeReports(input string, safetyFunc func(report string) bool) int {
	var total int
	for _, report := range strings.Split(input, "\n") {
		if safetyFunc(report) {
			total++
		}
	}
	return total
}

func isSafeBasic(report string) bool {
	var prev *int
	var prevDiff *int
	for _, lvlStr := range strings.Split(report, " ") {
		lvl := parsing.MustParseInt(lvlStr)
		if prev != nil {
			diff := lvl - *prev
			if maths.Abs(diff) < 1 || maths.Abs(diff) > 3 {
				return false
			}
			if prevDiff != nil && polaritiesDifferent(diff, *prevDiff) {
				return false
			}
			prevDiff = &diff
		}
		prev = &lvl
	}
	return true
}

func polaritiesDifferent(x, y int) bool {
	return (x < 0) != (y < 0)
}