package day2

import (
	"fmt"
	"strings"

	"github.com/dredly/aoc2024/internal/files"
)

func Part2Answer() {
	input := files.MustRead("inputdata/day2/real.txt")
	fmt.Printf("Day 2 part 2 answer: %d\n", numSafeReports(input, isSafeWithDampener))
}

func isSafeWithDampener(report string) bool {
	if isSafeBasic(report) {
		return true
	}
	levelStrs := strings.Split(report, " ")
	for i := range levelStrs {
		without := make([]string, 0)
		without = append(without, levelStrs[:i]...)
		without = append(without, levelStrs[i+1:]...)
		if isSafeBasic(strings.Join(without, " ")) {
			return true
		}
	}
	return false
}