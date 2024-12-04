package day3

import (
	"fmt"
	"strings"

	"github.com/dredly/aoc2024/internal/files"
)

const enablingInstruction = "do()"
const disablingInstruction = "don't()"

func Part2Answer() {
	input := files.MustRead("inputdata/day3/real.txt")
	fmt.Printf("Day 3 part 2 answer: %d\n", totalEnabledMultiplications(input))
}

func totalEnabledMultiplications(input string) int {
	disabledFrom := strings.Index(input, disablingInstruction)
	if disabledFrom == -1 {
		return totalMultiplications(input)
	}
	firstEnabledPortion := input[:disabledFrom]
	remaining := input[disabledFrom:]
	reEnabledFrom := strings.Index(remaining, enablingInstruction)
	if reEnabledFrom == -1 {
		return totalMultiplications(firstEnabledPortion)
	}
	afterReEnabled := remaining[reEnabledFrom+len(enablingInstruction):]
	return totalEnabledMultiplications(firstEnabledPortion) + totalEnabledMultiplications(afterReEnabled)
}