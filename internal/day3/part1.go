package day3

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/dredly/aoc2024/internal/files"
	"github.com/dredly/aoc2024/internal/parsing"
)

var mulRegex = regexp.MustCompile(`mul\(\d+,\d+\)`)

func Part1Answer() {
	input := files.MustRead("inputdata/day3/real.txt")
	fmt.Printf("Day 3 part 1 answer: %d\n", totalMultiplications(input))
}

func totalMultiplications(input string) int {
	multiplicationInstructions := mulRegex.FindAllString(input, -1)
	var total int
	for _, instr := range multiplicationInstructions {
		total += multiply(instr)
	}
	return total
}

func multiply(instruction string) int {
	trimmed := strings.TrimFunc(instruction, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	left, right, _ := strings.Cut(trimmed, ",")
	return parsing.MustParseInt(left) * parsing.MustParseInt(right)
}