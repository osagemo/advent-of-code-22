package main

import (
	"fmt"
	"testing"
)

const input1 = `A Y
B X
C Z`

func TestDay1Part1(t *testing.T) {
	result := Part1(input1)
	expected := 15

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay1Part2(t *testing.T) {
	result := Part2(input1)
	expected := 12

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
