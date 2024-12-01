package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
	"unicode"
)

//go:embed input.txt
var input string

func Part1(input string) int {
	runeScores := map[rune]int{}
	for i := 1; i <= 26; i++ {
		runeScores[rune(i+96)] = i
		runeScores[unicode.ToUpper(rune(i+96))] = i + 26
	}

	totalScore := 0
	for _, line := range strings.Split(input, "\n") {
		currentLineLetters := map[rune]bool{}
		half := len(line) / 2

		for i := 0; i < len(line); i++ {
			if i < half {
				if _, ok := currentLineLetters[rune(line[i])]; !ok {
					currentLineLetters[rune(line[i])] = false
				}
				continue
			}

			if counted, ok := currentLineLetters[rune(line[i])]; ok {
				if counted {
					continue
				}
				totalScore += runeScores[rune(line[i])]
				currentLineLetters[rune(line[i])] = true
			}
		}

	}
	return totalScore
}

func Part2(input string) int {
	runeScores := map[rune]int{}
	for i := 1; i <= 26; i++ {
		runeScores[rune(i+96)] = i
		runeScores[unicode.ToUpper(rune(i+96))] = i + 26
	}

	totalScore := 0

	groupIndex := 0
	currentGroupLetters := map[rune]int{}
	for _, line := range strings.Split(input, "\n") {
		currentLineLetters := map[rune]bool{}
		for i := 0; i < len(line); i++ {
			c := rune(line[i])
			if _, ok := currentLineLetters[c]; !ok {
				currentLineLetters[c] = true
				if _, ok := currentGroupLetters[c]; ok {
					currentGroupLetters[c]++
					if currentGroupLetters[c] == 3 {
						totalScore += runeScores[c]
					}
				} else {
					currentGroupLetters[c] = 1
				}

			}
			continue
		}

		if groupIndex == 2 {
			groupIndex = 0
			currentGroupLetters = map[rune]int{}
		} else {
			groupIndex++
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
