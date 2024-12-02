package parsing

import "strconv"

func MustParseInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic("Failed to convert string to int")
	}
	return res
}