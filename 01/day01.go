package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
// Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.
func Part1(input string) int {
	caloriesPerElf := getTotalCaloriesPerElf(input)
	largestCalories := caloriesPerElf[len(caloriesPerElf)-1]

	return largestCalories
}

func Part2(input string) int {
	caloriesPerElf := getTotalCaloriesPerElf(input)
	threeLargestCalories := caloriesPerElf[len(caloriesPerElf)-1] +
		caloriesPerElf[len(caloriesPerElf)-2] +
		caloriesPerElf[len(caloriesPerElf)-3]

	return threeLargestCalories
}

func getTotalCaloriesPerElf(input string) []int {
	caloriesPerElf := []int{}

	currentElfTotalCalories := 0
	for _, s := range strings.Split(input, "\n") {
		if len(s) <= 1 {
			caloriesPerElf = append(caloriesPerElf, currentElfTotalCalories)
			currentElfTotalCalories = 0
			continue
		}

		calories, error := strconv.Atoi(s)
		if error != nil {
			panic(error)
		}
		currentElfTotalCalories += calories
	}

	caloriesPerElf = append(caloriesPerElf, currentElfTotalCalories)
	sort.Ints(caloriesPerElf)
	return caloriesPerElf
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
