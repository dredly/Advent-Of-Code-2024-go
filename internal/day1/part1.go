package day1

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1Answer() {
	left, right := parse(mustReadFile("inputdata/day1/real.txt"))
	fmt.Printf("Day 1 part 1 answer: %d\n", totalDistance(left, right))
}

func mustReadFile(path string) string {
	bs, err := os.ReadFile(path)
	if err !=nil {
		panic("error reading file")
	}
	return string(bs)
}

func parse(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		l, r, _ := strings.Cut(line, "   ")
		left[i] = mustParseInt(l)
		right[i] = mustParseInt(r)
	}
	return left, right
}

func totalDistance(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	var total int
	for i, locIDLeft := range left {
		locIDRight := right[i]
		total += abs(locIDRight - locIDLeft)
	}
	return total
}

func mustParseInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic("Failed to convert string to int")
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}