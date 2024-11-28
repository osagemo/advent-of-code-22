package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

// A & X = Rock = 1
// B & Y= Paper = 2
// C = Scissors = 3

// Lose = 0
// Draw = 3
// Win = 6

func Part1(input string) int {

	return 0
}

func Part2(input string) int {

	return 0
}

func main() {
	input := strings.ReplaceAll(input, "\r\n", "\n")
	fmt.Println("Day 1")
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input))
	fmt.Println(time.Since(start))
}
