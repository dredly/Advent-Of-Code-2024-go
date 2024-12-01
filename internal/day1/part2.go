package day1

import "fmt"

func Part2Answer() {
	left, right := parse(mustReadFile("inputdata/day1/real.txt"))
	fmt.Printf("Day 1 part 2 answer: %d\n", similarity(left, right))
}

func similarity(left, right []int) int {
	frequencies := make(map[int]int)
	remainingRight := right
	var score int
	for _, locIDLeft := range left {
		freq, ok := frequencies[locIDLeft]
		if !ok {
			remainingRight, freq = removeAndCount(remainingRight, locIDLeft)
			frequencies[locIDLeft] = freq
		}
		score += freq * locIDLeft
	}
	return score
}

func removeAndCount(s []int, val int) ([]int, int) {
	var out []int
	var count int
	for _, v := range s {
		if v == val {
			count++
		} else {
			out = append(out, v)
		}
	}
	return out, count
}