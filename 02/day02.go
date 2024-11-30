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
// B & Y = Paper = 2
// C & Z = Scissors = 3

// Lose = 0
// Draw = 3
// Win = 6

func Part1(input string) int {
	totalScore := 0
	var aToScore = map[string]int{
		"X": 4,
		"Y": 8,
		"Z": 3,
	}
	var bToScore = map[string]int{
		"X": 1,
		"Y": 5,
		"Z": 9,
	}
	var cToScore = map[string]int{
		"X": 7,
		"Y": 2,
		"Z": 6,
	}

	for _, s := range strings.Split(input, "\n") {
		if len(s) == 0 {
			continue
		}
		letters := strings.Split(s, " ")
		if letters[0] == "A" {
			totalScore += aToScore[letters[1]]
		} else if letters[0] == "B" {
			totalScore += bToScore[letters[1]]
		} else if letters[0] == "C" {
			totalScore += cToScore[letters[1]]
		}
	}
	return totalScore
}

// A & X = Rock = 1
// B & Y = Paper = 2
// C & Z = Scissors = 3
func Part2(input string) int {
	totalScore := 0
	var aToScore = map[string]int{
		"X": 3,
		"Y": 4,
		"Z": 8,
	}
	var bToScore = map[string]int{
		"X": 1,
		"Y": 5,
		"Z": 9,
	}
	var cToScore = map[string]int{
		"X": 2,
		"Y": 6,
		"Z": 7,
	}

	for _, s := range strings.Split(input, "\n") {
		if len(s) == 0 {
			continue
		}
		letters := strings.Split(s, " ")
		if letters[0] == "A" {
			totalScore += aToScore[letters[1]]
		} else if letters[0] == "B" {
			totalScore += bToScore[letters[1]]
		} else if letters[0] == "C" {
			totalScore += cToScore[letters[1]]
		}
	}

	return totalScore
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
