#!/bin/bash

if [ -z "$1" ] || ! [[ $1 =~ ^[0-9]+$ ]] ; then
  echo "Usage: $0 <day-number>"
  exit 1
fi

day_without_zero=$1
day=$(printf "%02d" $1)

# Create folder
mkdir ${day}

# Create dayXX.go
cat <<EOL > ${day}/day${day}.go
package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

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
EOL

# Create dayXX_test.go
cat <<EOL > ${day}/day${day}_test.go
package main

import (
	"fmt"
	"testing"
)

const input1 = \`\`

func TestDay${day_without_zero}Part1(t *testing.T) {
	result := Part1(input1)
	expected := -1

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay${day_without_zero}Part2(t *testing.T) {
	result := Part2(input1)
	expected := -1

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
EOL

# Create an empty input.txt
touch ${day}/input.txt