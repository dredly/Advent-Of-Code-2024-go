package day5

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dredly/aoc2024/internal/files"
	"github.com/dredly/aoc2024/internal/parsing"
)

func Part1Answer() {
	input := files.MustRead("inputdata/day5/real.txt")
	fmt.Printf("Day 5 Part 1 Answer: %d\n", sumOfMiddlePageNumbersForCorrectlyOrderedUpdates(input))
}

func sumOfMiddlePageNumbersForCorrectlyOrderedUpdates(input string) int {
	rules, updates := parse(input)
	var sum int
	for _, u := range updates {
		if updateSatisifiesAllRules(u, rules) {
			sum += middle(u)
		}
	}
	return sum
}

type PageOrderingRule struct {
	before, after int
}

func updateSatisifiesAllRules(update []int, rules []PageOrderingRule) bool {
	for _, r := range rules {
		if !r.updateIsValid(update) {
			return false
		}
	}
	return true
}

func (r PageOrderingRule) updateIsValid(update []int) bool {
	beforeIdx := slices.Index(update, r.before)
	if beforeIdx < 0 {
		return true
	}
	afterIdx := slices.Index(update, r.after)
	if afterIdx < 0 {
		return true
	}
	return beforeIdx < afterIdx
}

func middle[T comparable](s []T) T {
	if len(s) % 2 == 0 {
		panic("can only call middle on slice of odd length")
	}
	return s[(len(s) - 1) / 2]
}

func parse(input string) ([]PageOrderingRule, [][]int) {
	rulesPortion, updatesPortion, _ := strings.Cut(input, "\n\n")
	var rules []PageOrderingRule
	for _, line := range strings.Split(rulesPortion, "\n") {
		leftStr, rightStr, _ := strings.Cut(line, "|")
		rules = append(rules, PageOrderingRule{
			before: parsing.MustParseInt(leftStr),
			after: parsing.MustParseInt(rightStr),
		})
	}

	var updates [][]int
	for _, line := range strings.Split(updatesPortion, "\n") {
		updates = append(updates, parseIntList(line))
	}
	return rules, updates
}

func parseIntList(input string) []int {
	var res []int
	for _, val := range strings.Split(input, ",") {
		res = append(res, parsing.MustParseInt(val))
	}
	return res
}