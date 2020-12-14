package main

import (
	"fmt"
	"io/ioutil"
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

	/*
		Right 1, down 1.
		Right 3, down 1. (original check)
		Right 5, down 1.
		Right 7, down 1.
		Right 1, down 2. */

	var r1d1, r3d1, r5d1, r7d1, r1d2 int

	for i, line := range lines {

		if string(line[(i*3)%len(line)]) == "#" { // R3,D1
			r3d1 += 1
		}
		if string(line[(i*1)%len(line)]) == "#" { // R1, D1
			r1d1 += 1
		}
		if string(line[(i*5)%len(line)]) == "#" { // R5, D1
			r5d1 += 1
		}
		if string(line[(i*7)%len(line)]) == "#" { // R7, D1
			r7d1 += 1
		}
		if (i % 2) == 1 {
			continue
		} else if string(line[((i/2)*1)%len(line)]) == "#" { // ** R1, D2
			r1d2 += 1
		}
	}
	total := r3d1 + r1d1 + r5d1 + r7d1 + r1d2
	mult := r3d1 * r1d1 * r5d1 * r7d1 * r1d2

	fmt.Printf("r3d1: %d r1d1: %d r5d1: %d r7d1: %d r1d2: %d\n", r3d1, r1d1, r5d1, r7d1, r1d2)
	fmt.Printf("Total trees smashies: %d\n", total)
	fmt.Printf("Multiples of smashies: %d\n", mult)
}
