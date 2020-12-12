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
	// attempt 1: snarf entire input as a string, slice on newlines
	input, err := ioutil.ReadFile("./input.txt")
	check(err)

	// slicey dicey!
	strnums := strings.Split(string(input), "\n")
	strnums = strnums[0 : len(strnums)-1]

	var nums []int

	for _, value := range strnums {
		num, err := strconv.Atoi(value)
		check(err)

		nums = append(nums, num)
	}

	// Brute force ftw. Create a sparse map to do it cleaner if time permits.
	// ref: https://blog.golang.org/maps m[2020-num1] == true?
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i >= j {
				continue
			}
			for k, num3 := range nums {
				if j >= k {
					continue
				}
				if num1+num2+num3 == 2020 {
					fmt.Printf("Part 2: Solution found: %d x %d x %d, %d\n", num1, num2, num3, (num1 * num2 * num3))
				}
			}
			if num1+num2 == 2020 {
				fmt.Printf("Part 1: Solution found: %d x %d, %d\n", num1, num2, (num1 * num2))
			}

		}
	}

}
