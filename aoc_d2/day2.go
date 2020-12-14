package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	check(err)

	lines := strings.Split(string(input), "\n")
	lines = lines[0 : len(lines)-1]

	var valid, valid_v2 []string

	for _, line := range lines {
		items := strings.Split(line, " ")
		nums := strings.Split(items[0], "-")

		low, err := strconv.Atoi(nums[0])
		check(err)
		high, err := strconv.Atoi(nums[1])
		check(err)

		char := rune(items[1][0])
		password := items[2]

		occurences := strings.Count(password, string(char))

		// Part 1: Count occurences
		if low <= occurences && occurences <= high {
			valid = append(valid, items[2])
		}

		// Part 2: Only char in first slot, not in 2nd.
		if (rune(password[low-1]) == char) != (rune(password[high-1]) == char) {
			valid_v2 = append(valid_v2, password)
		}

	}
	fmt.Printf("\nSolution: %d valid passwords", len(valid))
	fmt.Printf("\nSolution PartB: %d valid passwords\n", len(valid_v2))
}
