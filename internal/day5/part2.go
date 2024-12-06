package day5

import (
	"fmt"
	"slices"

	"github.com/dredly/aoc2024/internal/files"
)

func Part2Answer() {
	input := files.MustRead("inputdata/day5/real.txt")
	fmt.Printf("Day 5 Part 2 Answer: %d\n", sumOfMiddlePageNumbersForReorderedUpdates(input))
}

func sumOfMiddlePageNumbersForReorderedUpdates(input string) int {
	rules, updates := parse(input)
	var sum int
	for _, u := range updates {
		if !updateSatisifiesAllRules(u, rules) {
			sum += middle(reorderUpdate(u, rules))
		}
	}
	return sum
}

func reorderUpdate(update []int, rules []PageOrderingRule) []int {
	for _, r := range rules {
		r.applyToUpdate(update)
	}
	if !updateSatisifiesAllRules(update, rules) {
		return reorderUpdate(update, rules)
	}
	return update
}

func (r PageOrderingRule) applyToUpdate(update []int) []int {
	if r.updateIsValid(update) {
		return update
	}
	beforeIdx := slices.Index(update, r.before)
	afterIdx := slices.Index(update, r.after)
	update[beforeIdx], update[afterIdx] = update[afterIdx], update[beforeIdx]
	return update
}